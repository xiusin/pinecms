package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

type DocumentModel struct {
	Orm *xorm.Engine
}

func NewDocumentModel(orm *xorm.Engine) *DocumentModel {
	return &DocumentModel{Orm: orm}
}

func (w *DocumentModel) GetList(page, limit int64) (list []tables.IriscmsDocumentModel, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = w.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err !=nil {
		golog.Error(err)
	}
	return list, total
}
