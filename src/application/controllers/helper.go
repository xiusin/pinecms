package controllers

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/xiusin/pine/di"
)

type FieldShowInPageList struct {
	Show      bool   `json:"show"`
	Search    int    `json:"search"`
	FeSearch  bool   `json:"fe_search"`
	FormShow  bool   `json:"form_show"`
	Formatter string `json:"formatter"`
}

type LoginAdminPayload struct {
	jwt.Payload
	AdminId int64 `json:"admin_id"`
	RoleID  int64 `json:"role_id"`
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
