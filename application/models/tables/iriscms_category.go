package tables

type IriscmsCategory struct {
	Catid       int64 `xorm:"pk" json:"Catid"`
	Parentid    int64 `json:"Parentid"`
	Catname     string `json:"Catname"`
	Type        int64 `json:"Type"`
	Description string `json:"Description"`
	Thumb       string `json:"Thumb"`
	Url         string `json:"Url"`
	Listorder   int64 `json:"Listorder"`
	Ismenu      int64 `json:"Ismenu"`
	TplPrefix   string `json:"TplPrefix"`
	HomeTpl     string `json:"HomeTpl"`
	ContentTpl  string `json:"ContentTpl"`
}
