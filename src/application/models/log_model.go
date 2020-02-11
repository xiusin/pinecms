package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type LogModel struct {
	Orm *xorm.Engine
}

func NewLogModel(orm *xorm.Engine) *LogModel {
	return &LogModel{Orm: orm}
}

func (m *LogModel) GetList(page, limit int64) ([]tables.IriscmsLog, int64) {
	offset := (page - 1) * limit
	var list []tables.IriscmsLog
	var total int64
	total, _ = m.Orm.Count(&tables.IriscmsLog{})
	if err := m.Orm.Desc("logid").Limit(int(limit), int(offset)).Find(&list); err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
	}
	return list, total
}


func (m *LogModel) DeleteAll() bool {
	res, err := m.Orm.Where("1=1").Delete(&tables.IriscmsLog{})
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
		return false
	}
	if res > 0 {
		return true
	}
	return false
}

func (m *LogModel) DeleteBeforeByDate(date string) bool {
	res, err := m.Orm.Where("`time` <= ? ", date).Delete(&tables.IriscmsLog{})
	if err != nil {
		golog.Error(helper.GetCallerFuncName(), err)
		return false
	}
	if res > 0 {
		return true
	}
	return false
}
