package wechat

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
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
	c.OpBefore = c.before
}

func (c WechatRuleController) before(act int, params interface{}) error {
	if act == backend.OpEdit || act == backend.OpAdd {
		//sess := params.(*xorm.Session).Clone()
		sess := params.(*xorm.Session)
		data := c.Table.(*tables.WechatMsgReplyRule)
		sess.Where("Match_Value = ?", data.MatchValue).Where("appid = ?", data.AppId)
		if act == backend.OpEdit {
			sess.Where("id = ?", data.Id)
		}
		if exist, _ := sess.Exist(&tables.WechatMsgReplyRule{}); exist {
			return errors.New("规则匹配值已经存在")
		}
	}
	return nil
}
