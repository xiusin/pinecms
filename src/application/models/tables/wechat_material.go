package tables

type WechatMaterial struct {
	Id         int64     `json:"id" xorm:"pk autoincr"`
	Appid      string    `json:"appid" xorm:"char(20) not null comment('appid')"`
	Type       string    `json:"type" xorm:"varchar(5) comment('媒体素材类型')"`
	MediaId    string    `json:"media_id" xorm:"varchar(40) comment('媒体ID')"`
	Url        string    `json:"url"`
	UpdateTime LocalTime `json:"update_time"`
}
