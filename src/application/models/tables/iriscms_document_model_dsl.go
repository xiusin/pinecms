package tables

import "time"

type IriscmsDocumentModelDsl struct {
	Id           int64
	Mid          int64
	FormName     string
	TableField   string
	Html         string
	Required     int
	Datasource   string
	RequiredTips string
	Validator    string
	DeletedAt    time.Time `xorm:"deleted_at" json:"deleted_at"`
}
