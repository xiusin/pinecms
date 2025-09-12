package models

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xiusin/pine/contracts"

	"github.com/xiusin/pinecms/src/common/helper"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"

	"xorm.io/xorm"
)

type CategoryModel struct {
	orm   *xorm.Engine
	cache contracts.Cache
}

func init() {
	model := &CategoryModel{}
	di.Set(model, func(_ di.AbstractBuilder) (i any, err error) {
		return &CategoryModel{
			orm:   helper.GetORM(),
			cache: helper.Cache(),
		}, nil
	}, true)

	di.Bind(controllers.ServiceCatUrlPrefixFunc, func(_ di.AbstractBuilder) (any, error) { // (id int64) string
		return model.GetUrlPrefix, nil
	})
}

func NewCategoryModel() *CategoryModel {
	return di.MustGet(&CategoryModel{}).(*CategoryModel)
}

func (c *CategoryModel) GetPosArr(id int64) ([]tables.Category, error) {
	category := tables.Category{Catid: id}
	exists, err := helper.GetORM().Get(&category)
	if err != nil {
		pine.Logger().Error("getting category %d failed: %s", id, err.Error())
		return nil, ErrInternal
	}
	if !exists {
		return nil, ErrCategoryNotFound
	}
	var links []tables.Category
	for category.Parentid != 0 {
		links = append(links, category)
		parentid := category.Parentid
		category = tables.Category{Catid: parentid}
		helper.GetORM().Get(&category)
	}
	links = append(links, category)
	var reverse = func(s []tables.Category) []tables.Category {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}
	return reverse(links), nil
}
// GetTree generates a tree of categories.
func (c *CategoryModel) GetTree(categories []tables.Category, parentid int64) []map[string]any {
	var res = []map[string]any{}
	if len(categories) != 0 {
		models, _ := NewDocumentModel().GetList(1, 1000)
		var m = map[int64]string{}
		var modelMap = map[int64]string{}
		m[0] = "单页面"
		for _, model := range models {
			m[model.Id] = model.Name
			modelMap[model.Id] = model.Table
		}
		for _, category := range categories {
			if category.Parentid == parentid {
				modelName := m[category.ModelId]
				var total int64 = 0
				if category.Type == 2 {
					modelName = ""
				} else if category.Type == 0 {
					tableName := modelMap[category.ModelId]
					total, _ = c.orm.Table(controllers.GetTableName(tableName)).Where("catid = ?", category.Catid).Where("deleted_time IS NULL").Count()
				}

				var url = category.Url
				if category.Type != 2 {
					url = fmt.Sprintf("/%s/", c.GetUrlPrefix(category.Catid))
				}
				son := map[string]any{
					"parentid":    category.Parentid,
					"catid":       category.Catid,
					"catname":     category.Catname,
					"model_id":    modelName,
					"dir":         category.Dir,
					"url":         url,
					"type":        category.Type,
					"description": category.Description,
					"ismenu":      category.Ismenu,
					"listorder":   category.Listorder,
					"operateid":   category.Catid,
					"total":       total,
				}
				son["children"] = c.GetTree(categories, category.Catid)
				res = append(res, son)
			}
		}
	}
	return res
}

func (c *CategoryModel) GetWithDirForBE(dir string) *tables.Category {
	categories := c.GetAll(true)
	for _, v := range categories {
		if v.Dir == dir {
			return &v
		}
	}
	return nil
}

func (c *CategoryModel) GetAll(cache bool) ([]tables.Category, error) {
	var categories []tables.Category
	if !cache {
		if err := helper.Cache().Delete(controllers.CacheCategories); err != nil {
			pine.Logger().Error("deleting categories cache failed: %s", err.Error())
		}
	}
	err := helper.Cache().Remember(controllers.CacheCategories, &categories, func() (any, error) {
		err := c.orm.Asc("listorder").Desc("id").Find(&categories)
		if err != nil {
			pine.Logger().Error("getting all categories failed: %s", err.Error())
			return nil, ErrInternal
		}
		return &categories, nil
	})

	if err != nil {
		pine.Logger().Error("remembering categories cache failed: %s", err.Error())
		return nil, ErrInternal
	}
	return categories, nil
}

func (c *CategoryModel) GetCategoryMap(cache bool) (map[int64]tables.Category, error) {
	categories, err := c.GetAll(cache)
	if err != nil {
		return nil, err
	}
	m := map[int64]tables.Category{}
	for _, v := range categories {
		m[v.Catid] = v
	}
	return m, nil
}

func (c *CategoryModel) GetNextCategory(parentid int64) []tables.Category {
	var categories []tables.Category
	c.orm.Where("parentid=?", parentid).Asc("listorder").Desc("id").Find(&categories)
	if len(categories) != 0 {
		for _, v := range categories {
			categories = append(categories, c.GetNextCategory(v.Catid)...)
		}
	}
	return categories
}

func (c *CategoryModel) GetNextCategoryOnlyCatids(parentid int64, withSelf bool) []int64 {
	categories := c.GetNextCategory(parentid)
	var ids []int64
	if withSelf {
		ids = append(ids, parentid)
	}
	for _, v := range categories {
		ids = append(ids, v.Catid)
	}
	return ids
}

func (c *CategoryModel) GetSelectTree(parentid int64) ([]map[string]any, error) {
	categories := new([]tables.Category)
	err := c.orm.Where("parentid = ?", parentid).OrderBy("`listorder` ASC,`id` DESC").Find(categories)
	if err != nil {
		pine.Logger().Error("getting category select tree for parent %d failed: %s", parentid, err.Error())
		return nil, ErrInternal
	}
	maps := []map[string]any{}
	if len(*categories) > 0 {
		for _, v := range *categories {
			children, err := c.GetSelectTree(v.Catid)
			if err != nil {
				return nil, err
			}
			maps = append(maps, map[string]any{
				"value":    v.Catid,
				"label":    v.Catname,
				"children": children,
			})
		}
	}
	return maps, nil
}

// 取得内容管理右部分类tree结构
func (c *CategoryModel) GetContentRightCategoryTree(categories []tables.Category, parentid int64) []map[string]any {
	maps := []map[string]any{}
	if len(categories) > 0 {
		for _, v := range categories {
			if v.Parentid == parentid {
				maps = append(maps, map[string]any{
					"label":    v.Catname,
					"value":    v.Catid,
					"children": c.GetContentRightCategoryTree(categories, v.Catid),
				})
			}
		}
	}
	return maps
}

func (c *CategoryModel) DeleteById(id int64) error {
	res, err := c.orm.Delete(tables.Category{Catid: id})
	if err != nil {
		pine.Logger().Error("deleting category %d failed: %s", id, err.Error())
		return ErrInternal
	}
	if res == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

func (c *CategoryModel) GetCategory(id int64) *tables.Category {
	category := tables.Category{Catid: id}
	res, _ := c.orm.Get(&category)
	if res {
		return &category
	}
	return nil
}

func (c *CategoryModel) GetCategoryByModelID(id int64) ([]tables.Category, error) {
	category := []tables.Category{}
	_ = c.orm.Where("model_id = ?", id).Find(&category)
	return category, nil
}

// 读取单个分类的信息
func (c *CategoryModel) GetCategoryFByIdForBE(id int64) (category *tables.Category, err error) {
	cacheKey := fmt.Sprintf(controllers.CacheCategoryInfoPrefix, id)
	category = &tables.Category{}
	var exists bool
	err = c.cache.GetWithUnmarshal(cacheKey, category)
	if err != nil {
		exists, err = c.orm.ID(id).Get(category)
		if err != nil {
			pine.Logger().Error("getting category %d failed: %s", id, err.Error())
			return nil, ErrInternal
		}
		if !exists {
			return nil, ErrCategoryNotFound
		}
		category.Page = NewPageModel().GetPage(id)
		prefix, err := c.GetUrlPrefix(id)
		if err != nil {
			return nil, err
		}
		category.UrlPrefix = prefix
		c.cache.SetWithMarshal(cacheKey, category)
	}
	if category.Type == 0 {
		category.Model = NewDocumentModel().GetByIDForBE(category.ModelId)
	}
	if category.Page != nil {
		category.Content = category.Page.Content // 内容关联读取到缓存
	}
	return category, nil
}
// GetUrlPrefixWithCategoryArr generates the URL prefix for a category from an array of its ancestors.
func (c *CategoryModel) GetUrlPrefixWithCategoryArr(cats []tables.Category) string {
	var urlPrefix string
	cur := cats[len(cats)-1]
	prev := "" // 记录同名排除. 有时会父级和子级使用相同的dir,用来指向子级内容
	for _, v := range cats {
		v.Dir = strings.Trim(v.Dir, " /")
		if v.Dir == "" || prev == v.Dir {
			continue
		}
		prev = v.Dir
		urlPrefix = filepath.Join(urlPrefix, v.Dir)
	}
	if cur.Dir == "" {
		if cur.Type == 0 {
			model := NewDocumentModel().GetByIDForBE(cur.ModelId)
			urlPrefix = filepath.Join(urlPrefix, fmt.Sprintf("%s_%d", model.Table, cur.Catid))
		} else {
			urlPrefix = filepath.Join(urlPrefix, fmt.Sprintf("page_%d", cur.Catid))
		}
	}
	if runtime.GOOS == "windows" {
		urlPrefix = strings.ReplaceAll(urlPrefix, "\\", "/")
	}
	return urlPrefix
}

func (c *CategoryModel) GetUrlPrefix(id int64) (string, error) {
	arr, err := c.GetPosArr(id)
	if err != nil {
		return "", err
	}
	return c.GetUrlPrefixWithCategoryArr(arr), nil
}
func (c *CategoryModel) AddCategory(category tables.Category) error {
	_, err := c.orm.Insert(&category)
	if err != nil {
		pine.Logger().Error("adding category failed: %s", err.Error())
		return ErrInsertFailed
	}
	return nil
}

func (c *CategoryModel) UpdateCategory(category *tables.Category) error {
	res, err := c.orm.Where("id=?", category.Catid).Update(category)
	if err != nil {
		pine.Logger().Error("updating category %d failed: %s", category.Catid, err.Error())
		return ErrUpdateFailed
	}
	if res == 0 {
		return ErrCategoryNotFound
	}
	return nil
}

// 判断是否是子分类
func (c *CategoryModel) IsSonCategory(id, parentid int64) bool {
	cat := []tables.Category{}
	c.orm.Where("parentid=?", id).Find(&cat)
	if len(cat) == 0 {
		return false
	}
	flag := false
	for _, son := range cat {
		if son.Catid == parentid {
			return true
		}
		if c.IsSonCategory(son.Catid, parentid) {
			flag = true
		}
	}
	return flag
}
