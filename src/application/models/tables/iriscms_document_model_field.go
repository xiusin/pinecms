package tables

type IriscmsDocumentModelField struct {
	Id   int64 `xorm:"pk"`
	Name string
	Type string
	Desc string
	Html string
}
