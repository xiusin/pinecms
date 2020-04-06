package models

import (
	"errors"
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"log"
	"path/filepath"
	"strings"

	"github.com/xiusin/pinecms/src/application/models/tables"

	"github.com/go-xorm/xorm"
)

type CategoryModel struct {
	orm *xorm.Engine
}

var ErrCategoryNotExists = errors.New("category not exists")

func NewCategoryModel() *CategoryModel {
	return &CategoryModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (c CategoryModel) GetPosArr(id int64) []tables.Category {
	category := tables.Category{Catid: id}
	c.orm.Get(&category)
	var links []tables.Category
	for category.Parentid != 0 {
		links = append(links, category)
		parentid := category.Parentid
		category = tables.Category{Catid: parentid}
		c.orm.Get(&category)
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

func (c CategoryModel) GetTree(categorys []tables.Category, parentid int64) []map[string]interface{} {
	var res = []map[string]interface{}{}
	if len(categorys) != 0 {
		// 筛选
		models, _ := NewDocumentModel().GetList(1, 1000)
		var m = map[int64]string{}
		m[0] = "单页面"
		for _, model := range models {
			m[model.Id] = model.Name
		}
		for _, category := range categorys {
			if category.Parentid == parentid {
				modelName := m[category.ModelId]
				if category.Type > 1 {
					modelName = ""
				}
				son := map[string]interface{}{
					"catid":       category.Catid,
					"catname":     category.Catname,
					"model_id":    modelName,
					"dir":         category.Dir,
					"type":        category.Type,
					"description": category.Description,
					"ismenu":      category.Ismenu,
					"listorder":   category.Listorder,
					"operateid":   category.Catid,
					"router":      category.ManagerContentRouter,
				}
				son["children"] = c.GetTree(categorys, category.Catid)
				res = append(res, son)
			}
		}
	}
	return res
}

func (c CategoryModel) GetAll() []tables.Category {
	categorys := []tables.Category{}
	c.orm.Asc("listorder").Desc("catid").Find(&categorys)
	return categorys
}

func (c CategoryModel) GetNextCategory(parentid int64) []tables.Category {
	var categories []tables.Category
	c.orm.Where("parentid=?", parentid).Asc("listorder").Desc("catid").Find(&categories)
	return categories
}

func (c CategoryModel) GetSelectTree(parentid int64) []map[string]interface{} {
	categorys := new([]tables.Category)
	err := c.orm.Where("parentid = ?", parentid).OrderBy("`listorder` ASC,`catid` DESC").Find(categorys)
	if err != nil {
		log.Println(err.Error())
	}
	maps := []map[string]interface{}{}
	if len(*categorys) > 0 {
		for _, v := range *categorys {
			maps = append(maps, map[string]interface{}{
				"id":       v.Catid,
				"text":     v.Catname,
				"children": c.GetSelectTree(v.Catid),
			})
		}
	}
	return maps
}

//取得内容管理右部分类tree结构
func (c CategoryModel) GetContentRightCategoryTree(categorys []tables.Category, parentid int64) []map[string]interface{} {
	maps := []map[string]interface{}{}
	if len(categorys) > 0 {
		for _, v := range categorys {
			if v.Parentid == parentid {
				maps = append(maps, map[string]interface{}{
					"id":       v.Catid,
					"text":     v.Catname,
					"type":     v.Type,
					"url":      v.Url,
					"children": c.GetContentRightCategoryTree(categorys, v.Catid),
				})
			}
		}
	}
	return maps
}

func (c CategoryModel) DeleteById(id int64) bool {
	res, err := c.orm.Delete(tables.Category{Catid: id})
	if err != nil || res == 0 {
		return false
	}
	return true
}

func (c CategoryModel) GetCategory(id int64) (tables.Category, error) {
	category := tables.Category{Catid: id}
	res, err := c.orm.Get(&category)
	if err != nil || !res {
		return category, err
	}
	return category, nil
}

func (c CategoryModel) GetCategoryByModelID(id int64) ([]tables.Category, error) {
	category := []tables.Category{}
	_ = c.orm.Where("model_id = ?", id).Find(&category)
	return category, nil
}

func (c CategoryModel) GetCategoryFullWithCache(id int64) (category *tables.Category, err error) {
	caheKey := fmt.Sprintf(controllers.CacheCategoryInfoPrefix, id)
	icache := di.MustGet(controllers.ServiceICache).(cache.AbstractCache)
	category = &tables.Category{}
	var exists bool
	err = icache.GetWithUnmarshal(caheKey, category)
	if err != nil {
		exists, err = c.orm.ID(id).Get(category)
		if err != nil || !exists {
			if err == nil {
				err = ErrCategoryNotExists
			}
			category = nil
			return
		}
		category.UrlPrefix = c.GetUrlPrefix(id)
		icache.SetWithMarshal(caheKey, category) // 忽略失败判断
	}
	category.Model = NewDocumentModel().GetByIDWithCache(category.ModelId)
	if category.Model == nil {
		return nil, ErrCategoryNotExists
	}

	return category, nil
}

func (c CategoryModel) GetUrlPrefixWithCategoryArr(cats []tables.Category) string {
	var urlPrefix string
	cur := cats[len(cats)-1]
	for _, v := range cats {
		v.Dir = strings.Trim(v.Dir, " /")
		if v.Dir == "" {
			continue
		}
		urlPrefix = filepath.Join(urlPrefix, v.Dir)
	}
	if cur.Dir == "" {
		if cur.Type == 0 {
			model := NewDocumentModel().GetByID(cur.ModelId)
			urlPrefix = filepath.Join(urlPrefix, fmt.Sprintf("%s_%d", model.Table, cur.Catid))
		} else {
			urlPrefix = filepath.Join(urlPrefix, fmt.Sprintf("page_%d", cur.Catid))
		}
	}
	return urlPrefix
}

func (c CategoryModel) GetUrlPrefix(id int64) string {
	return c.GetUrlPrefixWithCategoryArr(c.GetPosArr(id))
}
func (c CategoryModel) AddCategory(category tables.Category) bool {
	_, err := c.orm.Insert(&category)
	if err != nil {
		pine.Logger().Error("AddCategoryError", err)
		return false
	}

	return true
}

func (c CategoryModel) UpdateCategory(category tables.Category) bool {
	res, err := c.orm.Where("catid=?", category.Catid).Update(&category)
	if err != nil || res == 0 {
		pine.Logger().Error("CategoryModel::UpdateCategory", err, res)
		return false
	}
	return true
}


//判断是否是子分类
func (c CategoryModel) IsSonCategory(id, parentid int64) bool {
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
