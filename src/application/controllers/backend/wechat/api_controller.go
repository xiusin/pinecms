package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"net/http"
	"time"
)

func msgHandler(ctx *pine.Context) {
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error(err)
			ctx.WriteJSON(pine.H{"code": 400500, "message": err})
		}
	}()

	appid := ctx.Params().Get("appid")
	if len(appid) == 0 {
		ctx.Abort(404)
		return
	}

	account, _ := GetOfficialAccount(appid)

	req := &http.Request{}
	if err := fasthttpadaptor.ConvertRequest(ctx.RequestCtx, req, true); err != nil {
		ctx.Abort(500, "无法转换请求: "+err.Error())
		return
	}

	srv := account.GetServer(req, &wechatResponseWrapper{ctx.RequestCtx})
	if !srv.Validate() {
		ctx.Abort(403, "消息无法验证")
		return
	}

	orm := ctx.Value("orm").(*xorm.Engine)
	srv.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		var rules []*tables.WechatMsgReplyRule
		var msgData interface{}
		var replyMsg interface{}
		var err error

		return nil

		//goqu.From(controllers.GetTableName("wechat_msg_reply_rule")).Where(
		//	goqu.Or(goqu.Ex{
		//		"match_value": msg.Content,
		//		"exact_match": 1,
		//	}, goqu.Ex{
		//		"match_value": msg.Content,
		//		"exact_match": 0,
		//	}))

		var baseSql = "SELECT * FROM %s WHERE ((match_value = ? AND exact_match = 1) OR " +
			"(INSTR(?, match_value) > 0 AND  exact_match = 0)) AND appid = '" + appid +
			"' AND status = 1 ORDER BY exact_match DESC, id DESC LIMIT 1"

		orm.SQL(fmt.Sprintf(baseSql, controllers.GetTableName("wechat_msg_reply_rule")), msg.Content, msg.Content).Find(&rules)

		if len(rules) == 0 || len(rules[0].ReplyContent) == 0 {
			return nil
		}

		rule := rules[0]

		if message.MsgTypeMiniprogrampage == rule.ReplyType || message.MsgTypeMusic == rule.ReplyType || message.MsgTypeVideo == rule.ReplyType {
			replyMsg = &WechatMsg{}
			json.Unmarshal([]byte(rule.ReplyContent), replyMsg)
		} else if message.MsgTypeNews == rule.ReplyType {
			replyMsg = []*message.Article{}
			err = json.Unmarshal([]byte(rule.ReplyContent), &replyMsg)
		}
		if err != nil {
			pine.Logger().Error("自动文章信息失败", err)
			return nil
		}

		switch rule.ReplyType {
		case string(message.MsgTypeText):
			msgData = message.NewText(rule.ReplyContent)
		case message.MsgTypeImage:
			msgData = message.NewImage(rule.ReplyContent)
		case message.MsgTypeMiniprogrampage:
			rm := replyMsg.(*WechatMsg)
			msgData = message.NewCustomerMiniprogrampageMessage(msg.OpenID, rm.Title, appid, rm.PagePath, rm.ThumbMediaID)
		case message.MsgTypeNews:
			msgData = message.NewNews(replyMsg.([]*message.Article))
		case message.MsgTypeMusic:
			rm := replyMsg.(*WechatMsg)
			msgData = message.NewMusic(rm.Title, rm.Description, rm.MusicURL, rm.HQMusicURL, rm.ThumbMediaID)
		case message.MsgTypeVideo:
			rm := replyMsg.(*WechatMsg)
			msgData = message.NewVideo(rm.MediaID, rm.Title, rm.Description)
		case message.MsgTypeVoice:
			msgData = message.NewVoice(rule.ReplyContent)
		case message.MsgTypeTransfer:
			msgData = message.NewTransferCustomer(rule.ReplyContent)
		}
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: msgData}
	})

	//处理消息接收以及回复
	if err := srv.Serve(); err != nil {
		pine.Logger().Warning("处理消息异常", err)
		return
	}

	if srv.RequestMsg != nil {
		orm.InsertOne(&tables.WechatLog{
			AppId:     appid,
			OpenId:    string(srv.RequestMsg.FromUserName),
			MsgType:   string(srv.RequestMsg.MsgType),
			Detail:    srv.RequestMsg,
			CreatedAt: tables.LocalTime(time.Unix(srv.RequestMsg.CreateTime, 0)),
		})
	}

	if err := srv.Send(); err != nil {
		pine.Logger().Error("编码回复消息失败", err)
	}
}
