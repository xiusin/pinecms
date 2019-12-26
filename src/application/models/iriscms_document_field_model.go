package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

type DocumentModelFieldModel struct {
	Orm *xorm.Engine
}

func NewDocumentModelFieldModel(orm *xorm.Engine) *DocumentModelFieldModel {
	return &DocumentModelFieldModel{Orm: orm}
}

func (w *DocumentModelFieldModel) GetList(page, limit int64) (list []*tables.IriscmsDocumentModelField, total int64) {
	offset := (page - 1) * limit
	total, _ = w.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}
