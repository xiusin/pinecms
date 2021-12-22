package models

import (
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
)

type DocumentModelFieldModel struct {
	orm *xorm.Engine
}

func NewDocumentModelFieldModel() *DocumentModelFieldModel {
	return &DocumentModelFieldModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (w *DocumentModelFieldModel) GetList(page, limit int64) (list []*tables.DocumentModelField, total int64) {
	offset := (page - 1) * limit
	total, _ = w.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}

func (w *DocumentModelFieldModel) GetMap() map[int64]*tables.DocumentModelField {
	var list []*tables.DocumentModelField
	var mapList = map[int64]*tables.DocumentModelField{}
	_ = w.orm.Find(&list)
	for _, v := range list {
		mapList[v.Id] = v
	}
	return mapList
}
