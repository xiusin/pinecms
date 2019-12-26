package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

type DocumentModel struct {
	Orm *xorm.Engine
}

func NewDocumentModel(orm *xorm.Engine) *DocumentModel {
	return &DocumentModel{Orm: orm}
}

func (w *DocumentModel) GetList(id, page, limit int64) (list []*tables.IriscmsDocumentModel, total int64) {
	offset := (page - 1) * limit
	total, _ = w.Orm.Id(id).Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}
