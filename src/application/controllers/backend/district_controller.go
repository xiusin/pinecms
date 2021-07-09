package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DistrictController struct {
	BaseController
}

func (c *DistrictController) Construct() {
	c.Table = &tables.District{}
	c.Entries = &[]*tables.District{}
	c.AppId = "admin"
	c.Group = "字典管理"
	c.SubGroup = "地区管理"
	c.ApiEntityName = "地区"
	c.BaseController.Construct()


	// todo 增加自动下载导入数据库功能
}
