package tables

type DocumentModelField struct {
	Id   int64  `json:"id" xorm:"id pk autoincr"`
	Name string `json:"name"`
	Opt  string `xorm:"-"` //字段默认配置设置方式是: attr:type:value 后端拆分为类型区分并加入到attr属性内
	Type string `json:"type"`
	Desc string `json:"desc"`
	Html string `json:"html"`
}
