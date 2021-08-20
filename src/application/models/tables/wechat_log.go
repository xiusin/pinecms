package tables

type WechatLog struct {
	Id        int64     `json:"id"`
	AppId     string    `json:"appid" xorm:"char(20)"`
	OpenId    string    `json:"openid" xorm:"varchar(32)"`
	Inout     uint      `json:"in_out" xorm:"tinyint(1)"`
	MsgType   string    `json:"msg_type" xorm:"varchar(50)"`
	Detail    string    `json:"detail" xorm:"json"`
	CreatedAt LocalTime `json:"created_at" xorm:"created"`
}
