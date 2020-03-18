package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type LogModel struct {
	orm *xorm.Engine
}

func NewLogModel() *LogModel {
	return &LogModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (l *LogModel) GetList(page, limit int64) ([]tables.Log, int64) {
	offset := (page - 1) * limit
	var list []tables.Log
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
