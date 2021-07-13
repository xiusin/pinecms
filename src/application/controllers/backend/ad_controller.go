package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type AdController struct {
	BaseController
}

func (c *AdController) Construct() {
	c.Table = &tables.Advert{}
	c.Entries = &[]tables.Advert{}
	c.AppId = "admin"
	c.Group = "广告管理"
	c.SubGroup = "广告管理"
	c.ApiEntityName = "广告"
	c.BaseController.Construct()
}
