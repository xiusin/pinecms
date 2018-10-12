package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type WechatController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (c *WechatController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/wechat/verify", "Verify")
}

func (c *WechatController) Verify() {
}
