package tables

type DocumentModelField struct {
	Id      int64       `json:"id" xorm:"id pk autoincr"`
	Name    string      `json:"name" xorm:"comment('组件名称') varchar(30)"`
	Type    string      `json:"type" xorm:"comment('数据库字段类型') varchar(20)"`
	Desc    string      `json:"desc" xorm:"comment('组件使用场景描述') text"`
	VueComp string      `json:"vue_comp" xorm:"comment('对应vue组件') varchar(30)"`
	Props   interface{} `json:"props" xorm:"comment('属性配置') json"`
	Html    string      `json:"html" xorm:"comment('自定义html') text"`
}
