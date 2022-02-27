package models

import (
	"errors"
	"fmt"
	"github.com/xiusin/pinecms/src/common/helper"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"

	"xorm.io/xorm"
)

type CategoryModel struct {
	orm   *xorm.Engine
	cache cache.AbstractCache
}

var ErrCategoryNotExists = errors.New("category not exists")

func init() {
	model := &CategoryModel{}
	di.Set(model, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return &CategoryModel{
			orm:   builder.MustGet(controllers.ServiceXorm).(*xorm.Engine),
			cache: builder.MustGet("cache.AbstractCache").(cache.AbstractCache),
		}, nil
	}, true)

	di.Bind(controllers.ServiceCatUrlPrefixFunc, func(builder di.AbstractBuilder) (interface{}, error) { // (id int64) string
		return model.GetUrlPrefix, nil
	})
}

func NewCategoryModel() *CategoryModel {
	return di.MustGet(&CategoryModel{}).(*CategoryModel)
}


func (c *CategoryModel) GetPosArr(id int64) []tables.Category {
	category := tables.Category{Catid: id}
	exists, err := helper.GetORM().Get(&category)
	if !exists {
		panic(fmt.Sprintf("分类:%d不存在: %s", id, err))
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
	return reverse(links)
}
func (c *CategoryModel) GetTree(categorys []tables.Category, parentid int64) []map[string]interface{} {
	var res = []map[string]interface{}{}
	if len(categorys) != 0 {
		models, _ := NewDocumentModel().GetList(1, 1000)
		var m = map[int64]string{}
		var modelMap = map[int64]string{}
		m[0] = "单页面"
		for _, model := range models {
			m[model.Id] = model.Name
			modelMap[model.Id] = model.Table
		}
		for _, category := range categorys {
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
				son := map[string]interface{}{
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
				son["children"] = c.GetTree(categorys, category.Catid)
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

func (c *CategoryModel) GetAll(withCache bool) []tables.Category {
	var categories []tables.Category
	if !withCache {
		_ = helper.AbstractCache().Delete(controllers.CacheCategories)
	}
	_ = helper.AbstractCache().Remember(controllers.CacheCategories, &categories, func() (interface{}, error) {
		if err := c.orm.Asc("listorder").Desc("id").Find(&categories); err != nil {
			pine.Logger().Error("查询语句错误", err)
		}
		return &categories, nil
	})
	return categories
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

func (c *CategoryModel) GetSelectTree(parentid int64) []map[string]interface{} {
	categorys := new([]tables.Category)
	err := c.orm.Where("parentid = ?", parentid).OrderBy("`listorder` ASC,`id` DESC").Find(categorys)
	if err != nil {
		log.Println(err.Error())
	}
	maps := []map[string]interface{}{}
	if len(*categorys) > 0 {
		for _, v := range *categorys {
			maps = append(maps, map[string]interface{}{
				"value":    v.Catid,
				"label":    v.Catname,
				"children": c.GetSelectTree(v.Catid),
			})
		}
	}
	return maps
}

//取得内容管理右部分类tree结构
func (c *CategoryModel) GetContentRightCategoryTree(categorys []tables.Category, parentid int64) []map[string]interface{} {
	maps := []map[string]interface{}{}
	if len(categorys) > 0 {
		for _, v := range categorys {
			if v.Parentid == parentid {
				maps = append(maps, map[string]interface{}{
					"label":    v.Catname,
					"value":    v.Catid,
					"children": c.GetContentRightCategoryTree(categorys, v.Catid),
				})
			}
		}
	}
	return maps
}

func (c *CategoryModel) DeleteById(id int64) bool {
	res, err := c.orm.Delete(tables.Category{Catid: id})
	if err != nil || res == 0 {
		return false
	}
	return true
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
	caheKey := fmt.Sprintf(controllers.CacheCategoryInfoPrefix, id)
	category = &tables.Category{}
	var exists bool
	err = c.cache.GetWithUnmarshal(caheKey, category)
	if err != nil {
		exists, err = c.orm.ID(id).Get(category)
		if err != nil || !exists {
			if err == nil {
				err = ErrCategoryNotExists
			}
			category = nil
			return
		}
		category.Page = NewPageModel().GetPage(id)
		category.UrlPrefix = c.GetUrlPrefix(id)
		c.cache.SetWithMarshal(caheKey, category)
	}
	if category.Type == 0 {
		category.Model = NewDocumentModel().GetByIDForBE(category.ModelId)
	}
	if category.Page != nil {
		category.Content = category.Page.Content // 内容关联读取到缓存
	}
	return category, nil
}
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

func (c *CategoryModel) GetUrlPrefix(id int64) string {
	return c.GetUrlPrefixWithCategoryArr(c.GetPosArr(id))
}
func (c *CategoryModel) AddCategory(category tables.Category) bool {
	_, err := c.orm.Insert(&category)
	if err != nil {
		pine.Logger().Error("AddCategoryError", err)
		return false
	}

	return true
}

func (c *CategoryModel) UpdateCategory(category *tables.Category) bool {
	res, err := c.orm.Where("id=?", category.Catid).Update(category)
	if err != nil || res == 0 {
		pine.Logger().Error("CategoryModel::UpdateCategory", err, res)
		return false
	}
	return true
}

//判断是否是子分类
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
