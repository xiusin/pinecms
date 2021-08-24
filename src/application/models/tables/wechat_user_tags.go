package tables

// WechatMember 微信粉丝
type WechatMember struct {
	Id             int64   `xorm:"pk autoincr"`
	Appid          string  `json:"appid"`
	Openid         string  `json:"openid"`
	Phone          string  `json:"phone"`
	Nickname       string  `json:"nickname"`
	Sex            int     `json:"sex"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Headimgurl     string  `json:"headimgurl"`
	SubscribeTime  string  `json:"subscribe_time"`
	Subscribe      bool    `json:"subscribe"`
	Unionid        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	TagidList      []int64 `json:"tagid_list" xorm:"json"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`
	Poster         string  `json:"poster"`
}
