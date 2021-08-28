package backend

import "github.com/xiusin/pinecms/src/application/models/tables"

type TagsController struct {
	BaseController
}

func (c *TagsController) Construct() {
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Table = &tables.Tags{}
	c.Entries = &[]*tables.Tags{}
	c.ApiEntityName = "标签"
	c.Group = "标签管理"
	c.BaseController.Construct()
}
