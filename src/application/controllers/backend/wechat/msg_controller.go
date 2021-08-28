package wechat

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

type WechatMagController struct {
	backend.BaseController
}

func (c *WechatMagController) Construct() {
	c.Table = &tables.WechatLog{}
	c.Entries = &[]tables.WechatLog{}
	c.BaseController.Construct()

	c.OpAfter = c.after
}

func (c *WechatMagController) PostReply() {
	var inputs = struct {
		AppId        string `json:"appid"`
		OpenId       string `json:"openid"`
		ReplyContent string `json:"replyContent"`
		ReplyType    string `json:"replyType"`
	}{}

	if err := c.Ctx().BindJSON(&inputs); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	} else if len(inputs.AppId) == 0 || len(inputs.OpenId) == 0 {
		helper.Ajax("消息参数错误", 1, c.Ctx())
		return
	}

	account, _ := GetOfficialAccount(inputs.AppId)

	msg := &message.CustomerMessage{
		ToUser:  inputs.OpenId,
		Msgtype: message.MsgType(inputs.ReplyType),
		Text:    &message.MediaText{Content: inputs.ReplyContent},
	}

	if err := account.GetCustomerMessageManager().Send(msg); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	//回复消息：演示回复用户发送的消息
	c.Orm.InsertOne(&tables.WechatLog{
		AppId:     inputs.AppId,
		OpenId:    inputs.OpenId,
		Inout:     1,
		MsgType:   inputs.ReplyType,
		Detail:    &message.MixMessage{Content: inputs.ReplyContent},
		CreatedAt: tables.LocalTime(time.Now()),
	})

	helper.Ajax("回复成功", 0, c.Ctx())
}

func (c *WechatMagController) after(act int, params interface{}) error {
	if act == backend.OpList {
		lists := c.Entries.(*[]tables.WechatLog)
		openIds := backend.ArrayCol(*lists, "OpenId")
		var fans []tables.WechatMember
		c.Orm.In("openid", openIds).Cols("openid", "nickname", "headimgurl").Find(&fans)

		fansMap := backend.ArrayColMap(fans, "Openid")
		for i, log := range *lists {
			member, ok := fansMap[log.OpenId].(tables.WechatMember)
			if ok {
				log.FansInfo = &member
			}
			(*lists)[i] = log
		}
	}
	return nil
}
