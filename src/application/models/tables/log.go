package tables

import "time"

type Log struct {
	Logid    int64 `xorm:"pk"`
	Type     int
	Message  string
	Userid   int64 `json:"userid"`
	Username string
	Ip       string
	Time     time.Time
}
