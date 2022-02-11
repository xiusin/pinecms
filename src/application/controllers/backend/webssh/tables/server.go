package tables

import (
	"time"
)

type SSHServer struct {
	Id         int64
	Nickname   string
	Ip         string
	Port       int
	Username   string
	Password   string `json:"-"`
	BindUser   uint   `json:"-"`
	BeforeTime time.Time
}
