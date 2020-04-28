package tables

import "time"

type WechatMessageLog struct {
	Logid   int64 `xorm:"pk"`
	Content string
	Time    time.Time
}
