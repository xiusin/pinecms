package controllers

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/xiusin/pine/di"
)

//LoginAdminPayload 登录JWT载体
type LoginAdminPayload struct {
	jwt.Payload
	Id        int64   `json:"id"`
	AdminId   int64   `json:"admin_id"`
	AdminName string  `json:"admin_name"`
	RoleID    []int64 `json:"role_id"`
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
