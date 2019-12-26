package tables

type IriscmsDocumentModel struct {
	Id          int64 `xorm:"pk"`
	Name        int64
	Enabled     int
	IsSystem    int
	ModelType   int
	FeTplIndex  string
	FeTplList   string
	FeTplDetail string
}
