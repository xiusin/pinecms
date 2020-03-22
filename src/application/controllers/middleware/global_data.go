package middleware

import (
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"strings"

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
		lower := map[string]string{}
		for k, v := range settingData {
			lower[strings.ToLower(k)] = v
		}
		ctx.Render().ViewData("global", lower)
		host := ctx.Request().Host
		ctx.Render().ViewData("url", func(path string, params ...map[string]interface{}) string {
			var query = "?"
			if len(params) > 0 {
			for k, v := range params[0] {
				query += k + "=" + fmt.Sprintf("%s", v)
			}
			}
			if len(query) == 1 {
				query = ""
			}
			return "//" + host + "/" + path + query
		})
		ctx.Render().ViewData("detail_url", func(aid string, tid ...string) string {
			if len(tid) == 0 {
				tid = []string{ctx.GetString("tid")}
			}
			return fmt.Sprintf("//%s/detail?aid=%s&tid=%s", host, aid, tid[0])
		})
		ctx.Next()
	}

}
