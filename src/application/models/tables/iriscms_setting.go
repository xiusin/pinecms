package tables

type IriscmsSetting struct {
	Key   string `xorm:"unique" json:"key"`
	Value string `json:"value"`
}
