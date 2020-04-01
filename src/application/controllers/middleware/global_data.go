package middleware

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
	"strings"
)

func SetGlobalConfigData() pine.Handler {
	return func(ctx *pine.Context) {
		settingData, err := config.SiteConfig()
		if err != nil {
			pine.Logger().Error("无法读取到配置内容:" + err.Error())
			return
		}
		ctx.Set(controllers.CacheSetting, settingData)

		lower := map[string]string{}
		for k, v := range settingData {
			lower[strings.ToLower(k)] = v
		}

		ctx.Render().ViewData("global", lower)
		ctx.Next()
	}
}
