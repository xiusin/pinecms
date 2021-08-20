package wechat

import (
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type WechatUserController struct {
	backend.BaseController
}

func (c *WechatUserController) Construct() {
	c.Table = &tables.WechatMember{}
	c.Entries = &[]tables.WechatMember{}
	c.BaseController.Construct()
}
