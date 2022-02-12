package Apiform

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/common"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type Login struct {
	Phone int64  `form:"phone" binding:"required"`
	Code  string `form:"code" binding:"required"`
}

type Send struct {
	Phone string `form:"phone" binding:"required"`
}

type Slist struct {
	Page  int `form:"page" binding:"required"`
	Limit int `form:"limit" binding:"required"`
}

type List_resp struct {
	List  []tables.SSHServer
	Count uint
}

type GetTerm struct {
	ID       int64  `form:"id" binding:"required"`
	Password string `form:"password"`
	Setpass  string `from:"setpass"`
}

type WsAuth struct {
	Sid string `uri:"sid" binding:"required,uuid"`
}

type Edit struct {
	ID       int64  `form:"id" binding:"required"`
	Nickname string `form:"nickname"`
	Ip       string `form:"ip"`
	Port     int    `form:"port"`
	Username string `form:"username"`
}

type EditPass struct {
	ID       int64  `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Addser struct {
	Nickname string `form:"nickname"`
	Ip       string `form:"ip" binding:"required"`
	Port     int    `form:"port" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Desc     string `form:"desc"`
}

type SerInfo struct {
	ID       int64
	Ip       string
	Port     int
	Username string
	Password string
	BindUser int64
}

func (t *GetTerm) Decode(server tables.SSHServer) (sid string, err error) {
	sid = uuid.Must(uuid.NewV4(), nil).String()
	sPass, err := common.AesDecryptCBC(server.Password, []byte(t.Setpass))
	if err != nil {
		return "", err
	}
	if sPass == "" {
		return "", errors.New("秘钥验证失败")
	} else {
		var serinfo = SerInfo{server.Id, server.Ip, server.Port, server.Username, sPass, server.BindUser}
		sInfo, _ := json.Marshal(serinfo)
		helper.AbstractCache().Set(sid, sInfo, 10)
	}

	return sid, nil
}
