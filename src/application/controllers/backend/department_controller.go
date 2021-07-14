package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type DepartmentController struct {
	BaseController
}

func (c *DepartmentController) Construct() {
	c.Table = &tables.Department{}
	c.Entries = &[]tables.Department{}
	c.Group = "系统管理"
	c.SubGroup = "部门管理"
	c.ApiEntityName = "部门"
	c.BaseController.Construct()
}


func (c *DepartmentController) GetSelect() {
	c.Orm.Find(c.Entries)
	var kvs = []tables.KV{
		{Label: "顶级部门", Value: 0},
	}
	for _, department := range *c.Entries.(*[]tables.Department) {
		kvs = append(kvs, tables.KV{
			Label: department.Name,
			Value: department.Id,
		})
	}
	helper.Ajax(kvs, 0, c.Ctx())
}
