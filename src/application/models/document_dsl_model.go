package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/pine/di"
)

type DocumentModelDslModel struct {
	orm *xorm.Engine
}

func NewDocumentFieldDslModel() *DocumentModelDslModel {
	return &DocumentModelDslModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (w *DocumentModelDslModel) GetList(mid int64)  []tables.IriscmsDocumentModelDsl {
	//todo need cache handler
	var list []tables.IriscmsDocumentModelDsl
	err := w.orm.Where("mid = ?", mid).Find(&list)
	if err != nil {
		golog.Error("NewDocumentFieldDslModel: ", err)
	}
	return list
}

func (w *DocumentModelDslModel) DeleteByMID(mid int64) bool {
	_, err := w.orm.Where("mid=?", mid).Delete(&tables.IriscmsDocumentModelDsl{})
	if err != nil {
		golog.Error("DocumentModelDslModel.DeleteByMID", err)
		return false
	}
	return true
}