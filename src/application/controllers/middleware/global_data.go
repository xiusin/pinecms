package middleware

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
	controllers "github.com/xiusin/iriscms/src/application/controllers"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
)

func FrontendGlobalViewData(xorm *xorm.Engine) func(ctx context.Context) {
	return func(ctx context.Context) {
		if ctx.Path() == "/" {
			settingData, err := getSetting(xorm)
			if err != nil {
				ctx.Application().Logger().Error("无法读取到配置内容:" + err.Error())
				ctx.StopExecution()
				return
			}
			if settingData["SITE_OPEN"] == "关闭" {
				ctx.Redirect("/site/close", iris.StatusFound)
				return
			}
			ctx.ViewData("setting", settingData)
		}
		ctx.Next()
	}
}

func SetGlobalConfigData(xorm *xorm.Engine, cache *boltdb.Database) func(ctx context.Context) {
	//读取配置项
	return func(ctx context.Context) {
		var settingData = map[string]string{}
		settingData, err := getSetting(xorm)
		if err != nil {
			ctx.Application().Logger().Error("无法读取到配置内容:" + err.Error())
			ctx.StopExecution()
			return
		}
		ctx.Values().Set(controllers.CacheSetting, settingData) //todo 这里有问题吗?
		ctx.Next()
	}

}

func getSetting(xorm *xorm.Engine) (map[string]string, error) {
	var settingData = map[string]string{}
	var settings []tables.IriscmsSetting
	err := xorm.Find(&settings)
	if err != nil {
		return nil, err
	}
	if len(settings) != 0 {
		for _, v := range settings {
			settingData[v.Key] = v.Value
		}
	}
	return settingData, nil
}
