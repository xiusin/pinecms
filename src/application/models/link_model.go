package models

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type LinkModel struct {
	orm *xorm.Engine
}

func NewLinkModel() *LinkModel {
	return &LinkModel{orm: helper.GetORM()}
}

func (l *LinkModel) GetList(page, limit int64) ([]tables.Link, int64) {
	offset := (page - 1) * limit
	var list = []tables.Link{}
	var total int64
	var err error
	if total, err = l.orm.Desc("listorder").Limit(int(limit), int(offset)).FindAndCount(&list); err != nil {
		pine.Logger().Error(helper.GetCallerFuncName(), err)
	}
	return list, total
}

func (l *LinkModel) Add(data *tables.Link) int64 {
	id, err := l.orm.InsertOne(data)
	if err != nil {
		pine.Logger().Error(helper.GetCallerFuncName(), err)
	}
	return id
}

func (l *LinkModel) Delete(id int64) bool {
	res, err := l.orm.ID(id).Delete(&tables.Link{})
	if err != nil {
		pine.Logger().Error(helper.GetCallerFuncName(), err)
	}
	return res > 0
}

func (l *LinkModel) Get(id int64) *tables.Link {
	var link = &tables.Link{}
	ok, _ := l.orm.ID(id).Get(link)
	if !ok {
		return nil
	}
	return link
}

func (l *LinkModel) Update(data *tables.Link) bool {
	id, err := l.orm.ID(data.Id).Update(data)
	if err != nil {
		pine.Logger().Error(helper.GetCallerFuncName(), err)
	}
	return id > 0
}
