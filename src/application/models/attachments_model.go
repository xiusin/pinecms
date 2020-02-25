package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
)


const (
	IMG_TYPE = "img"
	FILE_TYPE = "file"
)

type AttachmentsModel struct {
	Orm *xorm.Engine
}

func NewAttachmentsModel(orm *xorm.Engine) *AttachmentsModel {
	return &AttachmentsModel{Orm: orm}
}

func (a *AttachmentsModel) GetList(page, limit int64) (list []tables.IriscmsAttachments, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = a.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err != nil {
		golog.Error(err)
	}
	return list, total
}
