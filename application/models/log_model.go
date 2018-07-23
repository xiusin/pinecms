package models

import (
	"fmt"
	"iriscms/application/models/tables"

	"github.com/go-ffmt/ffmt"
	"github.com/go-xorm/xorm"
)

type LogModel struct {
	Orm *xorm.Engine
}

func NewLogModel(orm *xorm.Engine) *LogModel {
	return &LogModel{Orm: orm}
}

func (this *LogModel) GetList(page, limit int64) ([]tables.IriscmsLog, int64) {
	offset := (page - 1) * limit
	list := []tables.IriscmsLog{}
	fmt.Println(offset, limit)
	var total int64
	total, _ = this.Orm.Count(&tables.IriscmsLog{})
	if err := this.Orm.Desc("logid").Limit(int(limit), int(offset)).Find(&list); err != nil {
		fmt.Println(err.Error())
	}
	return list, total
}

func (this *LogModel) DeleteBeforeByDate(date string) bool {
	res, err := this.Orm.Where("`time` <= ? ", date).Delete(&tables.IriscmsLog{})
	if err != nil {
		ffmt.Println(err.Error())
		return false
	}
	if res > 0 {
		return true
	}
	return false
}
