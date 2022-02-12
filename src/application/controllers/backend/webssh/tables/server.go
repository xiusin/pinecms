package tables

import (
	"time"
)

type SSHServer struct {
	Id         int64 `xorm:"pk autoincr"`
	Nickname   string
	Ip         string
	Port       int
	Username   string
	Password   string `json:"-"`
	BindUser   int64   `json:"-"`
	BeforeTime time.Time
}
