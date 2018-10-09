package tables

type IriscmsCategory struct {
	Catid       int64 `xorm:"pk"`
	Parentid    int64
	Catname     string
	Type        int64
	Description string
	Thumb       string
	Url         string
	Listorder   int64
	Ismenu      int64
	TplPrefix   string
	HomeTpl     string
	ContentTpl  string
}
