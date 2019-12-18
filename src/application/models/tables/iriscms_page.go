package tables

type IriscmsPage struct {
	Catid      int64 `xorm:"pk"`
	Title      string
	Keywords   string
	Content    string
	Updatetime int64
}
