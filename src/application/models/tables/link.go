package tables

type Link struct {
	Id        int64     `xorm:"pk autoincr id" json:"id"`
	Linktype  uint8     `json:"linktype" xorm:"comment('链接类型') tinyint(3)"`
	Name      string    `json:"name" xorm:"comment('链接名称') varchar(50)"`
	Logo      string    `json:"logo" xorm:"comment('链接LOGO') varchar(100)"`
	Url       string    `json:"url" xorm:"comment('链接地址') varchar(255)"`
	Introduce string    `json:"introduce" xorm:"comment('描述')"`
	Listorder int64     `json:"listorder" xorm:"comment('排序') int(10)"`
	Status    bool      `json:"status" xorm:"comment('状态:0=禁用 1=正常')"`
	CreatedAt LocalTime `json:"created_at" xorm:"comment('创建时间') created"`
}
