package controllers

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"path/filepath"
)

type FieldShowInPageList struct {
	Show      bool   `json:"show"`
	Search    int    `json:"search"`
	FeSearch  bool   `json:"fe_search"`
	Formatter string `json:"formatter"`
}

func GetStaticFile(filename string) string {
	return filepath.Join("./resources/pages/", filename)
}

func GetSetting(xorm *xorm.Engine, cache cache.AbstractCache) (map[string]string, error) {
	var settingData = map[string]string{}
	err := cache.GetWithUnmarshal(CacheSetting, &settingData)
	if err != nil {
		var settings []tables.Setting
		err := xorm.Find(&settings)
		if err != nil {
			return nil, err
		}
		if len(settings) != 0 {
			for _, v := range settings {
				settingData[v.Key] = v.Value
			}
		}
		if err = cache.SetWithMarshal(CacheSetting, &settingData); err != nil {
			return nil, err
		}
	}

	return settingData, nil
}

func GetInMap(data map[string]FieldShowInPageList, key string) FieldShowInPageList {
	s, o := data[key]
	if o {
		return s
	} else {
		return FieldShowInPageList{}
	}
}

func InStringArr(data []string, key string) bool {
	for _,v := range data {
		if v == key {
			return true
		}
	}
	return false
}

func GetTableName(name string) string {
	tablePrefix := di.MustGet(ServiceTablePrefix).(string)
	return tablePrefix + name
}
