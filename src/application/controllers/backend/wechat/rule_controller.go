package wechat

import (
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type WechatRuleController struct {
	backend.BaseController
}

func (c *WechatRuleController) Construct() {
	c.Table = &tables.WechatMsgReplyRule{}
	c.Entries = &[]tables.WechatMsgReplyRule{}
	c.SearchFields = []backend.SearchFieldDsl{
		{Field: "appid"},
	}
	c.BaseController.Construct()
}
