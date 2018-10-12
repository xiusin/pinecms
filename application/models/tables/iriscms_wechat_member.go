package tables

import "time"

type IriscmsWechatMember struct {
	Id          int64 `xorm:"pk"`
	Openid         string
	Mpid           string
	Nickname       string
	Sex            int
	HeadImgUrl     string
	SubScribeScene string
	Time           time.Time
}
