package wechat

import (
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type WechatAccountController struct {
	backend.BaseController
}

func (c *WechatAccountController) Construct() {
	c.Table = &tables.WechatAccount{}
	c.Entries = &[]tables.WechatAccount{}
	c.BaseController.Construct()
}
