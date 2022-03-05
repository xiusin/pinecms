package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type LevelController struct {
	BaseController
}

func (c *LevelController) Construct() {
	c.Table = &tables.Level{}
	c.Entries = &[]tables.Level{}
	c.Group = "系统管理"
	c.SubGroup = "职级管理"
	c.ApiEntityName = "职级"
	c.BaseController.Construct()
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "name", Op: "="},
	}
}
