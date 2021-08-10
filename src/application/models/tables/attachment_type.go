package tables

type Attachments struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" xorm:"comment('附件名称')"`
	Url        string    `json:"url" xorm:"comment('完整的链接地址')"`
	OriginName string    `json:"original" xorm:"comment('原始名称')"`
	Size       int64     `json:"size" xorm:"comment('附件大小')"`
	CreatedAt  LocalTime `json:"upload_time" xorm:"created"`
	Type       string    `json:"type" xorm:"comment('类型') varchar(30)"`
	ClassifyId int64     `json:"classifyId" xorm:"comment('归属分类ID') int(5)"`
	Md5        string    `json:"md5" xorm:"comment('附件的md5值') unique varchar(32)"`
}
