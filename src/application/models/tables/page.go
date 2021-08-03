package tables

type Page struct {
	Id          int64     `xorm:"pk comment('栏目ID')" json:"id"`
	Title       string    `json:"title" xorm:"comment('页面标题') varchar(100)"`
	Keywords    string    `json:"keywords" xorm:"comment('页面关键字') varchar(255)"`
	Description string    `json:"description" xorm:"comment('页面描述') varchar(255)"`
	Content     string    `json:"content" xorm:"comment('页面内容') text"`
	CreatedAt   LocalTime `json:"created_at" xorm:"created"`
	UpdatedAt   LocalTime `json:"updated_at" xorm:"updated"`
}
