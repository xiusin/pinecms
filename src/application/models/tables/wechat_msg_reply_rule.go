package tables

type WechatMsgReplyRule struct {
	Id              int64     `json:"id"`
	AppId           string    `json:"appid" xorm:"char(20) appid"`
	RuleName        string    `json:"ruleName"`
	MatchValue      string    `json:"matchValue"`
	ExactMatch      bool      `json:"exactMatch" xorm:"comment('是否精确匹配')"`
	ReplyType       string    `json:"replyType"`
	ReplyContent    string    `json:"replyContent"`
	Status          bool      `json:"status"`
	Desc            string    `json:"desc"`
	EffectTimeStart string    `json:"effectTimeStart" xorm:"time"`
	EffectTimeEnd   string    `json:"effectTimeEnd" xorm:"time"`
	Priority        uint      `json:"priority"`
	CreatedAt       LocalTime `json:"created_at" xorm:"created"`
	UpdatedAt       LocalTime `json:"updated_at" xorm:"updated"`
}
