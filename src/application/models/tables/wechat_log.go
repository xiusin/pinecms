package tables

import "github.com/silenceper/wechat/v2/officialaccount/message"

type WechatLog struct {
	Id        int64               `json:"id" xorm:"pk autoincr"`
	AppId     string              `json:"appid" xorm:"char(20)"`
	OpenId    string              `json:"openid" xorm:"varchar(32)"`
	Inout     uint                `json:"in_out" xorm:"tinyint(1)" xorm:"comment('1=来自公众号的回复,0=来自粉丝的消息')"`
	MsgType   string              `json:"msg_type" xorm:"varchar(50)"`
	Detail    *message.MixMessage `json:"detail" xorm:"json"`
	CreatedAt LocalTime           `json:"created_at" xorm:"created"`
	FansInfo  *WechatMember       `json:"fans_info" xrom:"-"`
}
