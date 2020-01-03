package tables

import "time"

type IriscmsDocumentModelDsl struct {
	Id           int64
	Mid          int64
	FieldType    int64 // 字段ID  命名搞错了 先这样写程序吧
	FormName     string
	TableField   string
	Html         string
	Required     int
	Datasource   string
	RequiredTips string
	Validator    string
	DeletedAt    time.Time `xorm:"deleted_at" json:"deleted_at"`
}
