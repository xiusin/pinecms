package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type AdSpaceModel struct {
	orm *xorm.Engine
}

func NewAdSpaceModel() *AdSpaceModel {
	return &AdSpaceModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (l *AdSpaceModel) All() []tables.AdvertSpace {
	var list = []tables.AdvertSpace{}
	if err := l.orm.Desc("id").Find(&list); err != nil {
		pine.Logger().Error(err)
	}
	return list
}

func (l *AdSpaceModel) GetList(page, limit int64) ([]tables.AdvertSpace, int64) {
	offset := (page - 1) * limit
	var list = []tables.AdvertSpace{}
	var total int64
	var err error
	if total, err = l.orm.Desc("id").Limit(int(limit), int(offset)).FindAndCount(&list); err != nil {
		pine.Logger().Error(err)
	}
	return list, total
}

func (l *AdSpaceModel) Add(data *tables.AdvertSpace) int64 {
	id, err := l.orm.Insert(data)
	if err != nil {
		pine.Logger().Error(err)
	}
	return id
}

func (l *AdSpaceModel) Delete(id int64) bool {
	res, err := l.orm.ID(id).Delete(&tables.AdvertSpace{})
	if err != nil {
		pine.Logger().Error(err)
	}
	return res > 0
}

func (l *AdSpaceModel) Get(id int64) *tables.AdvertSpace {
	var link = &tables.AdvertSpace{}
	ok, _ := l.orm.ID(id).Get(link)
	if !ok {
		return nil
	}
	return link
}

func (l *AdSpaceModel) Update(data *tables.AdvertSpace) bool {
	id, err := l.orm.ID(data.Id).Update(data)
	if err != nil {
		pine.Logger().Error(err)
	}

	return id > 0
}
