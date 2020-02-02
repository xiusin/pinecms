package models

import (
	"log"

	tables "github.com/xiusin/iriscms/src/application/models/tables"

	"github.com/go-xorm/xorm"
)

type CategoryModel struct {
	orm *xorm.Engine
}

func NewCategoryModel(orm *xorm.Engine) *CategoryModel {
	return &CategoryModel{orm: orm}
}

func (this CategoryModel) GetTree(categorys []tables.IriscmsCategory, parentid int64) []map[string]interface{} {
	var res []map[string]interface{}
	if len(categorys) != 0 {
		// 筛选
		models ,_ := NewDocumentModel(this.orm).GetList(1, 1000)
		var m = map[int64]string {}
		for _, model := range models {
			m[model.Id] = model.Name
		}

		for _, category := range categorys {
			if category.Parentid == parentid {
				son := map[string]interface{}{
					"catid":       category.Catid,
					"catname":     category.Catname,
					"model_id":    m[category.ModelId],
					"type":        category.Type,
					"description": category.Description,
					"ismenu":      category.Ismenu,
					"listorder":   category.Listorder,
					"operateid":   category.Catid,
				}
				son["children"] = this.GetTree(categorys, category.Catid)
				res = append(res, son)
			}
		}
	}
	return res
}

func (this CategoryModel) GetAll() []tables.IriscmsCategory {
	categorys := new([]tables.IriscmsCategory)
	this.orm.Where("`type`<>? and `ismenu`=?", 2, 1).Asc("listorder").Desc("catid").Find(categorys)
	return *categorys
}

func (this CategoryModel) GetNextCategory(parentid int64) []tables.IriscmsCategory {
	categorys := new([]tables.IriscmsCategory)
	this.orm.Where("parentid=?", parentid).Asc("listorder").Desc("catid").Find(categorys)
	return *categorys
}

func (this CategoryModel) GetSelectTree(parentid int64) []map[string]interface{} {
	categorys := new([]tables.IriscmsCategory)
	err := this.orm.Where("parentid = ?", parentid).OrderBy("`listorder` ASC,`catid` DESC").Find(categorys)
	if err != nil {
		log.Println(err.Error())
	}
	maps := []map[string]interface{}{}
	if len(*categorys) > 0 {
		for _, v := range *categorys {
			maps = append(maps, map[string]interface{}{
				"id":       v.Catid,
				"text":     v.Catname,
				"children": this.GetSelectTree(v.Catid),
			})
		}
	}
	return maps
}

//取得内容管理右部分类tree结构
func (this CategoryModel) GetContentRightCategoryTree(categorys []tables.IriscmsCategory, parentid int64) []map[string]interface{} {
	maps := []map[string]interface{}{}
	if len(categorys) > 0 {
		for _, v := range categorys {
			if v.Parentid == parentid {
				maps = append(maps, map[string]interface{}{
					"id":       v.Catid,
					"text":     v.Catname,
					"type":     v.Type,
					"url":      v.Url,
					"children": this.GetContentRightCategoryTree(categorys, v.Catid),
				})
			}
		}
	}
	return maps
}

func (this CategoryModel) DeleteById(id int64) bool {
	res, err := this.orm.Delete(tables.IriscmsCategory{Catid: id})
	if err != nil || res == 0 {
		return false
	}
	return true
}

func (this CategoryModel) GetCategory(id int64) (tables.IriscmsCategory, error) {
	category := tables.IriscmsCategory{Catid: id}
	res, err := this.orm.Get(&category)
	if err != nil || !res {
		return category, err
	}
	return category, nil
}

func (this CategoryModel) AddCategory(category tables.IriscmsCategory) bool {
	res, err := this.orm.Insert(&category)
	if err != nil || res == 0 {
		return false
	}
	return true
}

func (this CategoryModel) UpdateCategory(category tables.IriscmsCategory) bool {
	res, err := this.orm.Where("catid=?", category.Catid).Update(&category)
	if err != nil || res == 0 {
		log.Println("CategoryModel::UpdateCategory", err, res)
		return false
	}
	return true
}

//判断是否是子分类
func (this CategoryModel) IsSonCategory(id, parentid int64) bool {
	cat := []tables.IriscmsCategory{}
	err := this.orm.Where("parentid=?", id).Find(&cat)
	if err != nil {
		log.Println("CategoryModel::IsSonCategory", err.Error())
		return false
	}
	if len(cat) == 0 {
		return false
	}
	flag := false
	for _, son := range cat {
		if son.Catid == parentid {
			return true
		}
		if this.IsSonCategory(son.Catid, parentid) {
			flag = true
		}
	}
	return flag
}
