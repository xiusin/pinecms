package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

type DocumentModelDslModel struct {
	Orm *xorm.Engine
}

func NewDocumentFieldDslModel(orm *xorm.Engine) *DocumentModelDslModel {
	return &DocumentModelDslModel{Orm: orm}
}

func (w *DocumentModelDslModel) GetList(id, page, limit int64) (list []*tables.IriscmsDocumentModelDsl, total int64) {
	offset := (page - 1) * limit
	total, _ = w.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}
