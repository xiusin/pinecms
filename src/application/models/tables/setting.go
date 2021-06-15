package tables

type Setting struct {
	Id        uint   `json:"id" xorm:"pk autoincr"`
	Key       string `xorm:"unique" json:"key" schema:"key"`
	FormName  string `json:"form_name" schema:"form_name"`
	Value     string `json:"value" schema:"value"`
	Group     string `json:"group" schema:"group"`
	Default   string `json:"default" schema:"default"`
	Listorder uint   `json:"listorder"`
	Editor    string `json:"editor" schema:"editor"`
	Extra     string `json:"extra" xorm:"-" schema:"extra"`
	Options   []KV   `json:"options" xorm:"json"`
}

type KV struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}
