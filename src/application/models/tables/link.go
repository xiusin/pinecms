package tables

import "time"

type Link struct {
	Linkid    int64     `xorm:"pk autoincr" json:"linkid"`
	Linktype  int64     `json:"linktype"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	Url       string    `json:"url"`
	Introduce string    `json:"introduce"`
	Listorder int64     `json:"listorder"`
	Passed    int8      `json:"passed"`
	Addtime   time.Time `json:"addtime"`
}
