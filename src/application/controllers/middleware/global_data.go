package middleware

import (
	"encoding/json"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
)

func FrontendGlobalViewData(xorm *xorm.Engine, cache *cache.Cache) func(ctx context.Context) {
	return func(ctx context.Context) {
		if ctx.Path() == "/" {
			settingData, err := getSetting(xorm, cache)
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

func SetGlobalConfigData(xorm *xorm.Engine, cache *cache.Cache) func(ctx context.Context) {
	//读取配置项
	return func(ctx context.Context) {
		settingData, err := getSetting(xorm, cache)
		if err != nil {
			ctx.Application().Logger().Error("无法读取到配置内容:" + err.Error())
			ctx.StopExecution()
			return
		}
		ctx.Values().Set(string(controllers.CacheSetting), settingData) //todo 这里有问题吗?
		ctx.Next()
	}

}

func getSetting(xorm *xorm.Engine, cache *cache.Cache) (map[string]string, error) {
	var settingData = map[string]string{}
	res := cache.Get(controllers.CacheSetting)
	if len(res) == 0 {
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
		setDataStr, err := json.Marshal(&settingData)
		if err == nil {
			if err := cache.Set(controllers.CacheSetting, setDataStr); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		err := json.Unmarshal([]byte(res), &settingData)
		if err != nil {
			return nil, err
		}
	}
	return settingData, nil
}
