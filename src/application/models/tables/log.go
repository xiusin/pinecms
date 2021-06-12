package tables

import "time"

type Log struct {
	Id       int64     `json:"id"`
	Params   string    `json:"params"`
	Uri      string    `json:"uri"`
	Userid   int64     `json:"userid"`
	Username string    `json:"username"`
	Ip       string    `json:"ip"`
	Time     time.Time `json:"time"`
	Method   string    `json:"method"`
}
