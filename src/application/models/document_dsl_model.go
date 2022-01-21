package models

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type DocumentModelDslModel struct {
	orm *xorm.Engine
}

func NewDocumentFieldDslModel() *DocumentModelDslModel {
	return &DocumentModelDslModel{orm: helper.GetORM()}
}

func (w *DocumentModelDslModel) GetList(mid int64) []tables.DocumentModelDsl {
	var list []tables.DocumentModelDsl
	err := w.orm.Where("mid = ?", mid).Asc("listorder").Asc("id").Find(&list)
	if err != nil {
		pine.Logger().Error("NewDocumentFieldDslModel: ", err)
	}
	return list
}

func (w *DocumentModelDslModel) DeleteByMID(mid int64) bool {
	_, err := w.orm.Where("mid=?", mid).Delete(&tables.DocumentModelDsl{})
	if err != nil {
		pine.Logger().Error("DocumentModelDslModel.DeleteByMID", err)
		return false
	}
	return true
}
