package tables

type DocumentModelField struct {
	Id       int64
	Name     string
	Opt      string //字段默认配置设置方式是: attr:type:value 后端拆分为类型区分并加入到attr属性内
	AmisType string // Amis框架的类型
	Type     string
	Desc     string
	Html     string
}
