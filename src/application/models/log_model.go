package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
	"github.com/xiusin/pine/di"
)

type LogModel struct {
	orm *xorm.Engine
}

func NewLogModel() *LogModel {
	return &LogModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (l *LogModel) GetList(page, limit int64) ([]tables.IriscmsLog, int64) {
	offset := (page - 1) * limit
	var list []tables.IriscmsLog
	var total int64
	total, _ = l.orm.Count(&tables.IriscmsLog{})
	if err := l.orm.Desc("logid").Limit(int(limit), int(offset)).Find(&list); err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
	}
	return list, total
}


func (l *LogModel) DeleteAll() bool {
	res, err := l.orm.Where("1=1").Delete(&tables.IriscmsLog{})
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
		return false
	}
	if res > 0 {
		return true
	}
	return false
}

func (l *LogModel) DeleteBeforeByDate(date string) bool {
	res, err := l.orm.Where("`time` <= ? ", date).Delete(&tables.IriscmsLog{})
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
		return false
	}
	if res > 0 {
		return true
	}
	return false
}
