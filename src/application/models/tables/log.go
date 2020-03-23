package tables

import "time"

type Log struct {
	Id       int64 `xorm:"pk"`
	Type     int
	Message  string
	UserID   int64 `json:"userid"`
	Username string
	Ip       string
	Time     time.Time
}
