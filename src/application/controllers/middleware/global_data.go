package middleware

import (
	"encoding/json"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

func SetGlobalConfigData(xorm *xorm.Engine, iCache cache.ICache) pine.Handler {
	return func(ctx *pine.Context) {
		settingData, err := getSetting(xorm, iCache)
		if err != nil {
			pine.Logger().Error("无法读取到配置内容:" + err.Error())
			return
		}
		ctx.Set(controllers.CacheSetting, settingData)
		ctx.Next()
	}

}

func getSetting(xorm *xorm.Engine, cache cache.ICache) (map[string]string, error) {
	var settingData = map[string]string{}
	res, err := cache.Get(controllers.CacheSetting)
	if err != nil {
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
		setDataStr, _ := json.Marshal(&settingData)
		if err := cache.Set(controllers.CacheSetting, setDataStr); err != nil {
			return nil, err
		}

	} else {
		err := json.Unmarshal(res, &settingData)
		if err != nil {
			return nil, err
		}
	}
	return settingData, nil
}
