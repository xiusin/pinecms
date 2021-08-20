package tables

type WechatMsgReplyRule struct {
	Id         int64     `json:"id"`
	AppId      string    `json:"appid" xorm:"char(20)"`
	RuleName   string    `json:"rule_name"`
	MatchValue string    `json:"match_value"`
	ExactMatch bool      `json:"exact_match" xorm:"comment('是否精确匹配')"`
	// todo 未结束
	CreatedAt  LocalTime `json:"created_at" xorm:"created"`
}
