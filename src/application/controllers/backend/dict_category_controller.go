package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DictCategoryController struct {
	BaseController
}

func (c *DictCategoryController) Construct() {
	c.AppId = "admin"
	c.Group = "字典分类管理"
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
		"list":  {Title: "日志列表", Desc: "查询系统接口请求日志列表"},
		"clear": {Title: "清空日志", Desc: "一键清理系统所有日志"},
	}
	c.BaseController.Construct()
}
