package tables

type DocumentModelField struct {
	Id       int64       `json:"id" xorm:"id pk autoincr"`
	Name     string      `json:"name" xorm:"comment('组件名称') varchar(30)"`
	Type     string      `json:"type" xorm:"comment('数据库字段类型') varchar(20)"`
	Desc     string      `json:"desc" xorm:"comment('组件使用场景描述') text"`
	ListComp string      `json:"list_comp" xorm:"comment('列表渲染组件') text"`
	FormComp string      `json:"form_comp" xorm:"comment('对应vue组件') text"`
	Props    interface{} `json:"props" xorm:"comment('属性配置') json"`
}
