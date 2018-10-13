package middleware

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iriscms/application/models/tables"
)

func FrontendGlobalViewData(app *iris.Application) func(ctx context.Context) {
	return func(ctx context.Context) {
		ctx.ViewData("setting", nil)
		ctx.Next()
	}
}

func SetGlobalConfigData(xorm *xorm.Engine, redisClient *redis.Pool) func(ctx context.Context) {
	//读取配置项
	return func(ctx context.Context) {
		var settingData = map[string]string{}
		client := redisClient.Get()
		defer client.Close()
		res, err := redis.StringMap(client.Do("GET", "setting"))
		if err != nil || len(res) == 0 {
			var settings []tables.IriscmsSetting
			err := xorm.Find(&settings)
			if err != nil {
				ctx.Application().Logger().Error("无法读取到配置内容:" + err.Error())
				ctx.StopExecution()
				return
			}
			if len(settings) != 0 {
				for _, v := range settings {
					settingData[v.Key] = v.Value
				}
			}
			setDataStr, err := json.Marshal(settingData)
			if err == nil {
				_, err = client.Do("SET", "setting", string(setDataStr))
				if err != nil {
					ctx.Application().Logger().Error("保存配置到redis失败",err.Error())
				}
			} else {
				ctx.Application().Logger().Error("保存配置到redis失败",err.Error())
			}
		} else {
			res = settingData
		}
		ctx.Values().Set("setting", settingData)
		ctx.Next()
	}

}
