package tables

import "time"

type IriscmsWechatMessageLog struct {
	Logid       int64 `xorm:"pk"`
	Content 	string
	Time        time.Time
}