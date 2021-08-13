package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type LogController struct {
	BaseController
}

func (c *LogController) Construct() {
	c.Group = "操作日志"

	c.KeywordsSearch = []KeywordWhere{
		{Field: "username", Op: "LIKE", DataExp: "%$?%"},
		{Field: "ip", Op: "LIKE", DataExp: "%$?%"},
		{Field: "uri", Op: "LIKE", DataExp: "%$?%"},
		{Field: "params", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"method": {Op: "="},
	}
	c.Table = &tables.RequestLog{}
	c.Entries = &[]tables.RequestLog{}
	c.apiEntities = map[string]apidoc.Entity{
		"list":  {Title: "日志列表", Desc: "查询系统接口请求日志列表"},
		"clear": {Title: "清空日志", Desc: "一键清理系统所有日志"},
	}
	c.BaseController.Construct()
}

func (c *LogController) PostClear() {
	_, _ = c.Orm.Where("id > 0").Delete(c.Table)
	helper.Ajax("清理成功", 0, c.Ctx())
}
