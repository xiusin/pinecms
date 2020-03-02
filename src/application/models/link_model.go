package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type LinkModel struct {
	orm *xorm.Engine
}

func NewLinkModel(orm *xorm.Engine) *LinkModel {
	return &LinkModel{orm: orm}
}

func (l *LinkModel) GetList(page, limit int64) ([]tables.IriscmsLink, int64) {
	offset := (page - 1) * limit
	var list = []tables.IriscmsLink{}
	var total int64
	var err error
	if total, err = l.orm.Desc("listorder").Limit(int(limit), int(offset)).FindAndCount(&list); err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
	}
	return list, total
}

func (l *LinkModel) Add(data *tables.IriscmsLink) int64 {
	id, err := l.orm.InsertOne(data)
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
	}
	return id
}

func (l *LinkModel) Delete(id int64) bool {
	res, err := l.orm.ID(id).Delete(&tables.IriscmsLink{})
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
	}
	return res > 0
}

func (l *LinkModel) Get(id int64) *tables.IriscmsLink {
	var link = &tables.IriscmsLink{}
	ok, _ := l.orm.ID(id).Get(link)
	if !ok {
		return nil
	}
	return link
}

func (l *LinkModel) Update(data *tables.IriscmsLink) bool {
	id, err := l.orm.ID(data.Linkid).Update(data)
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
	}
	return id > 0
}
