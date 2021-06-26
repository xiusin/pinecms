package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DictCategoryController struct {
	BaseController
}

func (c *DictCategoryController) Construct() {
	c.AppId = "admin"
	c.Group = "字典管理"
	c.SubGroup = "字典分类管理"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "key", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"status": {Op: "="},
	}
	c.Table = &tables.DictCategory{}
	c.Entries = &[]*tables.DictCategory{}
	c.apiEntities = map[string]apidoc.Entity{
		"list":   {Title: "分类列表", Desc: "查询系统内未删除的分类列表"},
		"add":    {Title: "新增分类", Desc: "新增一个分类，标识不可重复"},
		"edit":   {Title: "编辑分类", Desc: "编辑一个分类，标识不可重复"},
		"delete": {Title: "删除分类", Desc: "删除指定分类"},
		"info":   {Title: "分类详情", Desc: "获取指定分类的详情信息"},
	}
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
		exist, err := c.Orm.Table(&tables.Dict{}).Where("cid = ?", param.(*idParams).Id).Exist()
		if err != nil {
			return err
		}
		if exist {
			return errors.New("分类下存在字典数据，无法删除")
		}
	}
	return nil
}
