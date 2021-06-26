package tables

import "time"

type DictCategory struct {
	Id        uint      `xorm:"pk autoincr 'id'" json:"id"`
	Key       string    `json:"key" xorm:"key unique notnull"`
	Name      string    `json:"name notnull"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type Dict struct {
	Id        uint      `xorm:"pk autoincr 'id'" json:"id"`
	Cid       uint      `json:"cid" xorm:"notnull"`
	Name      string    `json:"name" xorm:"unique notnull"`
	Value     string    `json:"value"`
	Sort      uint      `json:"sort"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
