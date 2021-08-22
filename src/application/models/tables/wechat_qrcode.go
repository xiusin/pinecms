package tables

type WechatQrcode struct {
	Id         int64     `json:"id" xorm:"pk autoincr"`
	AppId      string    `json:"appid" xorm:"char(20)"`
	IsTemp     bool      `json:"is_temp" xorm:"comment('是否为临时二维码')"`
	SceneStr   string    `json:"scene_str"`
	Ticket     string    `json:"ticket"`
	Url        string    `json:"url"`
	ExpireTime LocalTime `json:"expire_time"`
	CreatedAt  LocalTime `json:"created_at" xorm:"created"`
}
