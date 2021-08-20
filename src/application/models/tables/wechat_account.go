package tables

type WechatAccount struct {
	Id       int64  `json:"id" xorm:"pk autoincr"`
	AppId    string `json:"appid" xorm:"char(20) not null comment('appid')"`
	Name     string `json:"name" xorm:"varchar(50) not null comment('公众号名称')"`
	Type     uint   `json:"type" xorm:"tinyint(1) comment('账号类型')"`
	Verified bool   `json:"verified" xorm:"tinyint(1) comment('认证状态')"`
	Secret   string `json:"secret" xorm:"char(32)"`
	Token    string `json:"token" xorm:"varchar(32)"`
	AesKey   string `json:"aesKey" xorm:"varchar(43)"`
}
