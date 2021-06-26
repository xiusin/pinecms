package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Construct() {
	c.KeywordsSearch = []KeywordWhere{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Table = &tables.Menu{}
	c.Entries = &[]tables.Menu{}

	c.AppId = "admin"
	c.Group = "菜单管理"
	c.SubGroup = "菜单管理"

	c.apiEntities = map[string]apidoc.Entity{
		"list":   {Title: "菜单列表", Desc: "获悉系统内配置的菜单列表"},
		"add":    {Title: "新增菜单", Desc: "新增一个菜单，可基于任何菜单添加子菜单"},
		"edit":   {Title: "编辑菜单", Desc: "修改给定菜单信息"},
		"delete": {Title: "删除菜单", Desc: "删除指定菜单"},
		"info":   {Title: "菜单详情", Desc: "获取指定菜单的详情信息"},
	}

	c.BaseController.Construct()
}

func (c *MenuController) PostList() {
	c.Orm.Find(c.Entries)
	helper.Ajax(c.Entries, 0, c.Ctx())
}
