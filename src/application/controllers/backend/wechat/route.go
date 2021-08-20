package wechat

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

func InitRouter(router *pine.Router) {

	go func() {
		err := di.MustGet(&xorm.Engine{}).(*xorm.Engine).Sync2(
			&tables.WechatAccount{},
			&tables.WechatMember{},
		)
		if err != nil {
			pine.Logger().Error("同步表结构失败", err)
		}
	}()

	router.Handle(new(WechatAccountController), "/wechat/account")
	router.Handle(new(WechatUserController), "/wechat/user")
}
