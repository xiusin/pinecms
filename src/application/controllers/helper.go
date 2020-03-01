package controllers

import (
	"encoding/json"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/pine/cache"
)

type FieldShowInPageList struct {
	Show      bool   `json:"show"`
	Formatter string `json:"formatter"`
}
func GetSetting(xorm *xorm.Engine, cache cache.ICache) (map[string]string, error) {
	var settingData = map[string]string{}
	res, err := cache.Get(CacheSetting)
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
		if err := cache.Set(CacheSetting, setDataStr); err != nil {
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

func GetInMap(data map[string]FieldShowInPageList,key string) FieldShowInPageList {
	s, o := data[key]
	if o {
		return s
	} else {
		return FieldShowInPageList{}
	}
}

