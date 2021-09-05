package wechat

import (
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
		pine.Logger().Error("转换请求失败", err)
		ctx.Abort(500, "转换请求失败")
		return
	}
	srv := account.GetServer(req, &wechatResponseWriter{ctx.RequestCtx})
	if !srv.Validate() {
		ctx.Abort(403, "消息来源异常")
		return
	}

	orm := ctx.Value("orm").(*xorm.Engine)

	//设置接收消息的处理方法
	srv.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		// 处理自动回复消息
		var rules []*tables.WechatMsgReplyRule

		var msgData interface{}

		var baseSql = "SELECT * FROM %s WHERE " +
			"((match_value = ? AND exact_match = 1) OR " +
			"(INSTR(?, match_value) > 0 AND  exact_match = 0)) AND appid = '" + appid +
			"' AND status = 1 ORDER BY exact_match DESC, id DESC LIMIT 1"

		sess := orm.SQL(fmt.Sprintf(baseSql, controllers.GetTableName("wechat_msg_reply_rule")), msg.Content, msg.Content)

		sess.Find(&rules)

		if len(rules) == 0 {
			return nil
		}

		rule := rules[0]

		switch rule.ReplyType {
		case string(message.MsgTypeText):
			msgData = message.NewText(rule.ReplyContent)
		case message.MsgTypeImage:
			msgData = message.NewImage(rule.ReplyContent)
		case message.MsgTypeMiniprogrampage:
			msgData = message.NewCustomerMiniprogrampageMessage(msg.OpenID, "", "", "", "")
		case message.MsgTypeNews:
			msgData = message.NewNews(nil)
		case message.MsgTypeMusic:
			msgData = message.NewMusic("", "", "", "", "")
		case message.MsgTypeVideo:
			msgData = message.NewVideo("", "", "")
		case message.MsgTypeVoice:
			msgData = message.NewVoice(rule.ReplyContent)
		case message.MsgTypeTransfer:
			msgData = message.NewTransferCustomer(rule.ReplyContent)
		}
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: msgData}
	})

	//处理消息接收以及回复
	if err := srv.Serve(); err != nil {
		pine.Logger().Error("处理消息异常", err)
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

	_ = srv.Send()
}

// Plugin TODO 回复插件
type Plugin struct {
}

type WechatMsg struct {
	Title string `json:"title"`
	// 小程序
	AppID        string `json:"appid"`
	PagePath     string `json:"pagePath"`
	ThumbMediaID string `json:"thumb_media_id"`

	// 音乐 视频
	Description string `json:"description"`
	MusicURL    string `json:"music_url"`
	HQMusicURL  string `json:"hq_music_url"`

	MediaID string `json:"media_id"`

	KfAccount string `json:"kf_account"`

	Articles []*message.Article
}
