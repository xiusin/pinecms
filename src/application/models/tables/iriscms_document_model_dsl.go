package tables

type IriscmsDocumentModelDsl struct {
	Id           int64 `xorm:"pk"`
	Mid          int64
	FormName     string
	Html         string
	Required     int
	Datasource   string
	RequiredTips string
	Validator    string
}
