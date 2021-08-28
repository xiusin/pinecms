package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type PositionController struct {
	BaseController
}

func (c *PositionController) Construct() {
	c.Table = &tables.Position{}
	c.Entries = &[]tables.Position{}
	c.Group = "系统管理"
	c.SubGroup = "岗位管理"
	c.ApiEntityName = "岗位"
	c.BaseController.Construct()
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "name", Op: "="},
		{Field: "code", Op: "="},
	}
}
