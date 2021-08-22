package wechat

import (
	"github.com/go-xorm/xorm"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/xiusin/pine"
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
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		ctx.Abort(404)
		return
	}
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
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	if err = srv.Serve(); err != nil {
		pine.Logger().Error("处理消息异常", err)
		return
	}

	//回复消息：演示回复用户发送的消息
	orm.InsertOne(&tables.WechatLog{
		AppId:     appid,
		OpenId:    string(srv.RequestMsg.FromUserName),
		MsgType:   string(srv.RequestMsg.MsgType),
		Detail:    srv.RequestMsg,
		CreatedAt: tables.LocalTime(time.Unix(srv.RequestMsg.CreateTime, 0)),
	})

	_ = srv.Send()
}
