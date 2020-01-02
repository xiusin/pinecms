package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

type DocumentModelDslModel struct {
	Orm *xorm.Engine
}

func NewDocumentFieldDslModel(orm *xorm.Engine) *DocumentModelDslModel {
	return &DocumentModelDslModel{Orm: orm}
}

func (w *DocumentModelDslModel) GetList(mid int64)  []tables.IriscmsDocumentModelDsl {
	var list []tables.IriscmsDocumentModelDsl
	err := w.Orm.Where("mid = ?", mid).Find(&list)
	if err != nil {
		golog.Error("NewDocumentFieldDslModel: ", err)
	}
	return list
}

func (w *DocumentModelDslModel) DeleteByMID(mid int64) bool {
	count, err := w.Orm.Where("mid=?", mid).Delete(&tables.IriscmsDocumentModelDsl{})
	if err != nil {
		golog.Error(err)
	}
	return count > 0
}