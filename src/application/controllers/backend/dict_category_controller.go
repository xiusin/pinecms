package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type DictCategoryController struct {
	BaseController
}

func (c *DictCategoryController) Construct() {
	c.AppId = "admin"
	c.Group = "字典管理"
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "key", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = []SearchFieldDsl{
		{Field: "status"},
	}
	c.Table = &tables.DictCategory{}
	c.Entries = &[]*tables.DictCategory{}
	c.ApiEntityName = "字典分类"
	c.BaseController.Construct()
	c.OpBefore = c.before
}

func (c *DictCategoryController) before(act int, param interface{}) error {
	switch act {
	case OpAdd, OpEdit:
		key := param.(*tables.DictCategory).Key
		id := param.(*tables.DictCategory).Id
		sess := c.Orm.Table(c.Table).Where("`key` = ?", key)
		if OpEdit == act {
			sess = sess.Where("id <> ?", id)
		}
		exist, err := sess.Exist()
		if err != nil {
			return err
		}
		if exist {
			return errors.New("该字典分类标识已经存在")
		}
	case OpDel:
		exist, err := c.Orm.Table(&tables.Dict{}).In("cid", param.(*idParams).Ids).Exist()
		if err != nil {
			return err
		}
		if exist {
			return errors.New("分类下存在字典数据，无法删除")
		}
	}
	return nil
}

// GetSelect 下拉列表
func (c *DictCategoryController) GetSelect() {
	_ = c.Orm.Where("status = 1").Find(c.Entries)
	m := c.Entries.(*[]*tables.DictCategory)
	var kv []tables.KV
	for _, model := range *m {
		kv = append(kv, tables.KV{Label: model.Name, Value: model.Key})
	}
	helper.Ajax(kv, 0, c.Ctx())
}
