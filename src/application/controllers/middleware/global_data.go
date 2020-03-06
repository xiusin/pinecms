package middleware

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers"
)

func SetGlobalConfigData(xorm *xorm.Engine, iCache cache.ICache) pine.Handler {
	return func(ctx *pine.Context) {
		settingData, err := controllers.GetSetting(xorm, iCache)
		if err != nil {
			pine.Logger().Error("无法读取到配置内容:" + err.Error())
			return
		}
		ctx.Set(controllers.CacheSetting, settingData)
		ctx.Next()
	}

}

