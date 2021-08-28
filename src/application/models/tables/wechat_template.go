package tables

type WechatMsgTemplate struct {
	Id              int64                    `json:"id" xorm:"pk autoincr"`
	Appid           string                   `json:"appid" xorm:"char(20) not null comment('appid')"`
	TemplateId      string                   `json:"template_id" xorm:"varchar(50) not null comment('模板ID')"`
	Title           string                   `json:"title" xorm:"varchar(20) comment('模板标题')"`
	Name            string                   `json:"name" xorm:"varchar(20) comment('模板名称')"`
	PrimaryIndustry string                   `json:"primary_industry"`
	DeputyIndustry  string                   `json:"deputy_industry"`
	Content         string                   `json:"content" xorm:"text"`
	Data            []map[string]interface{} `json:"data" xorm:"json"`
	Url             string                   `json:"url" xorm:"varchar(255)"`
	MiniProgram     map[string]interface{}   `json:"miniprogram" xorm:"json"`
	Status          bool                     `json:"status" xorm:"comment('是否有效0=无效,1=有效')"`
	Example         string                   `json:"example" xorm:"comment('模板示例')"`
	CreatedAt       LocalTime                `json:"created_at" xorm:"created"`
	UpdatedAt       LocalTime                `json:"updated_at" xom:"updated"`
}
