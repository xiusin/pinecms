package tables

type Setting struct {
	Id        uint                   `json:"id" xorm:"pk autoincr"`
	Key       string                 `xorm:"unique comment('配置KEY') varchar(50)" json:"key" schema:"key"`
	FormName  string                 `json:"form_name" xorm:"comment('名称') varchar(100)" schema:"form_name"`
	Value     string                 `json:"value" schema:"value" xorm:"comment('配置值') text"`
	Group     string                 `json:"group" schema:"group" xorm:"comment('所属分组')"`
	Default   string                 `json:"default" schema:"default" xorm:"comment('默认值')"`
	Listorder uint                   `json:"listorder" xorm:"comment('列表排序')"`
	Remark    string                 `json:"remark" xorm:"comment('配置描述')"`
	Editor    string                 `json:"editor" schema:"editor"`
	Extra     string                 `json:"extra" xorm:"-" schema:"extra"`
	Options   map[string]interface{} `json:"options" xorm:"json"`
}

type KV struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}
