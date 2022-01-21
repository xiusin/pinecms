package models

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type LogModel struct {
	orm *xorm.Engine
}

func NewLogModel() *LogModel {
	return &LogModel{orm: helper.GetORM()}
}

func (l *LogModel) GetList(page, limit int64) ([]tables.Log, int64) {
	offset := (page - 1) * limit
	var list = []tables.Log{}
	var total int64
	total, _ = l.orm.Count(&tables.Log{})
	if err := l.orm.Desc("logid").Limit(int(limit), int(offset)).Find(&list); err != nil {
		pine.Logger().Error(err)
	}
	return list, total
}

func (l *LogModel) DeleteAll() bool {
	res, err := l.orm.Where("1=1").Delete(&tables.Log{})
	if err != nil {
		pine.Logger().Error(err)
		return false
	}
	if res > 0 {
		return true
	}
	return false
}

func (l *LogModel) DeleteBeforeByDate(date string) bool {
	res, err := l.orm.Where("`time` <= ? ", date).Delete(&tables.Log{})
	if err != nil {
		pine.Logger().Error(err)
		return false
	}
	if res > 0 {
		return true
	}
	return false
}
