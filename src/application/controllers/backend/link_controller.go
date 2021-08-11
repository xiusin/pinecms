package backend

import "github.com/xiusin/pinecms/src/application/models/tables"

type LinkController struct {
	BaseController
}

func (c *LinkController) Construct() {
	c.KeywordsSearch = []KeywordWhere{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Table = &tables.Link{}
	c.Entries = &[]tables.Link{}
	c.ApiEntityName = "友链"
	c.Group = "友情链接管理"
	c.BaseController.Construct()
}
