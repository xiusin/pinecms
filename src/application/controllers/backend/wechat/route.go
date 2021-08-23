package wechat

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	app.ANY("/api/wechat/msg/:appid", msgHandler)
	router.Handle(new(WechatAccountController), "/wechat/account")
	router.Handle(new(WechatUserController), "/wechat/user")
	router.Handle(new(WechatMagController), "/wechat/msg")
	router.Handle(new(WechatQrcodeController), "/wechat/qrcode")
	router.Handle(new(WechatRuleController), "/wechat/rule")
	router.Handle(new(WechatMaterialController), "/wechat/material")
}
