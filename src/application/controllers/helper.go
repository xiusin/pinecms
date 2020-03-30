package controllers

import (
	"github.com/xiusin/pine/di"
)

type FieldShowInPageList struct {
	Show      bool   `json:"show"`
	Search    int    `json:"search"`
	FeSearch  bool   `json:"fe_search"`
	FormShow  bool   `json:"form_show"`
	Formatter string `json:"formatter"`
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
	for _, v := range data {
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
