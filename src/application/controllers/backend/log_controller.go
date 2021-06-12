package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type LogController struct {
	BaseController
}

func (c *LogController) Construct() {
	c.KeywordsSearch = []KeywordWhere{
		{Field: "username", Op: "LIKE", DataExp: "%$?%"},
		{Field: "ip", Op: "LIKE", DataExp: "%$?%"},
		{Field: "uri", Op: "LIKE", DataExp: "%$?%"},
		{Field: "params", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Table = &tables.Log{}
	c.Entries = &[]*tables.Log{}

	c.BaseController.Construct()
}

func (c *LogController) PostClear() {
	c.Orm.Where("id > 0").Delete(c.Table)
	helper.Ajax("清理成功", 0, c.Ctx())
}
