package tables

import "time"

type IriscmsAttachments struct {
	Id         int64 `xorm:"pk" json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	OriginName string `json:"origin_name"`
	Size       int64 `json:"size"`
	UploadTime time.Time `json:"upload_time"`
	Type       string `json:"type"`
}
