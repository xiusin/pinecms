package tables

type AdvertSpace struct {
	Id        int64  `xorm:"pk" xorm:"autoincr" json:"id"`
	Name      string `json:"name"`
}
