package tables

type Setting struct {
	Key      string `xorm:"unique" json:"key" schema:"key"`
	FormName string `json:"form" schema:"form"`
	Value    string `json:"value" schema:"value"`
	Group    string `json:"group" schema:"group"`
	Default  string `json:"default" schema:"default"`
	Editor   string `json:"editor" schema:"editor"`
	Extra    string `json:"extra" xorm:"-" schema:"extra"`
}
