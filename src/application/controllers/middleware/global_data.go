package middleware

import (
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"strconv"
	"strings"

	"github.com/xiusin/pinecms/src/application/controllers"
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
		host := ctx.Request().Host
		ctx.Render().ViewData("detail_url", func(aid string, tid ...string) string {
			if len(tid) == 0 {
				tid = []string{ctx.Params().Get("tid")}
			}
			iaid,_ := strconv.Atoi(aid)
			itid,_ := strconv.Atoi(tid[0])
			return fmt.Sprintf("//%s%s", host, helper.DetailUrl(itid, iaid))
		})
		ctx.Next()
	}

}
