package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Construct() {
	c.Table = &tables.Category{}
	c.Entries = &[]*tables.Category{}
	c.AppId = "admin"
	c.Group = "内容管理"
	c.SubGroup = "分类管理"
	c.ApiEntityName = "分类"
	c.BaseController.Construct()
}
