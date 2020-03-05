package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/pine/di"
)

const (
	IMG_TYPE  = "img"
	FILE_TYPE = "file"
)

type AttachmentsModel struct {
	orm *xorm.Engine
}

func NewAttachmentsModel() *AttachmentsModel {
	return &AttachmentsModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (a *AttachmentsModel) GetList(page, limit int64) (list []tables.IriscmsAttachments, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = a.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err != nil {
		golog.Error(err)
	}
	return list, total
}

func (a *AttachmentsModel) Delete(id int64) bool {
	res, _ := a.orm.ID(id).Delete(&tables.IriscmsAttachments{})
	return res > 0
}
