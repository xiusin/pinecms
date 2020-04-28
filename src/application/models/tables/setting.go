package tables

type Setting struct {
	Key       string      `xorm:"unique" json:"key"`
	Value     string      `json:"value"`
	Group     string      `json:"group"`
	Default   string      `json:"default"`
	FormName  string      `json:"form_name"`
	Editor    string      `json:"editor"`
	EditorOpt interface{} `json:"-" xorm:"-"`
}
