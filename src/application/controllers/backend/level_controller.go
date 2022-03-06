package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
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
	c.SelectOp = func(session *xorm.Session) {
		session.Where("status = ?", 1)
	}

	c.UniqCheckOp = func(i int, session *xorm.Session) {
		entity := c.Table.(*tables.Level)
		if i == OpAdd {
			session.Where("name = ?", entity.Name)
		} else {
			session.Where("name = ?", entity.Name).Where("id <> ?", entity.Id)
		}
	}
}
