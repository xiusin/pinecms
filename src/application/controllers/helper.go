package controllers

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/xiusin/pine/di"
)

type LoginAdminPayload struct {
	jwt.Payload
	AdminId int64 `json:"admin_id"`
	RoleID  int64 `json:"role_id"`
}

// InStringArr 判断是否在字符串数组内容
func InStringArr(data []string, key string) bool {
	for _, v := range data {
		if v == key {
			return true
		}
	}
	return false
}

// GetTableName 获取表名
func GetTableName(name string) string {
	tablePrefix := di.MustGet(ServiceTablePrefix).(string)
	return tablePrefix + name
}
