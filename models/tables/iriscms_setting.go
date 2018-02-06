package tables

type IriscmsSetting struct {
	Key   string `xorm:"unique"`
	Value string
}