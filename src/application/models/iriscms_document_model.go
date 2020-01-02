package models

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
)

type DocumentModel struct {
	Orm *xorm.Engine
}

func NewDocumentModel(orm *xorm.Engine) *DocumentModel {
	return &DocumentModel{Orm: orm}
}

func (w *DocumentModel) GetList(page, limit int64) (list []tables.IriscmsDocumentModel, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = w.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err != nil {
		golog.Error(err)
	}
	return list, total
}
func (w *DocumentModel) GetByID(id int64) *tables.IriscmsDocumentModel {
	var detail = &tables.IriscmsDocumentModel{}
	if err := w.Orm.ID(id).Find(detail); err != nil {
		golog.Error(err)
	}
	return detail
}

func (w *DocumentModel) DeleteByID(id int64) bool {
	if _, err := w.Orm.Transaction(func(session *xorm.Session) (i interface{}, err error) {
		i, err = w.Orm.ID(id).Delete(&tables.IriscmsDocumentModel{})
		if err != nil {
			golog.Error(err)
			return nil, err
		}
		if i == 0 {
			return nil, errors.New("删除了0条模型记录,错误表现") // 删除了0条记录, 返回失败
		}
		if !NewDocumentFieldDslModel(w.Orm).DeleteByMID(id) {
			golog.Errorf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id)
			return nil, errors.New(fmt.Sprintf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id))
		}
		return true, nil
	}); err != nil {
		return false
	}
	return true
}
