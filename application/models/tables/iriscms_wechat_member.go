package tables

import "time"

type IriscmsWechatMember struct {
	Id             int64 `xorm:"pk"`
	Openid         string
	Mpid           string
	Nickname       string
	Sex            int
	Headimgurl     string
	SubScribeScene string `xorm:"subscribe_scene"`
	Time           time.Time
}
