package models

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/xiusin/iriscms/src/application/models/tables"
)


const (
	CUSTOM_TYPE = iota
	SYSTEM_TYPE
)

type DocumentModel struct {
	Orm *xorm.Engine
}

func NewDocumentModel(orm *xorm.Engine) *DocumentModel {
	return &DocumentModel{Orm: orm}
}

func (d *DocumentModel) GetList(page, limit int64) (list []tables.IriscmsDocumentModel, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = d.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err != nil {
		golog.Error(err)
	}
	return list, total
}
func (d *DocumentModel) GetByID(id int64) *tables.IriscmsDocumentModel {
	var detail = &tables.IriscmsDocumentModel{}
	if _, err := d.Orm.ID(id).Get(detail); err != nil {
		golog.Error("document.model", err)
	}
	return detail
}

func (d *DocumentModel) DeleteByID(id int64) (bool, error) {
	// 先查找是否在用
	total, err := d.Orm.Where("model_id = ?", id).Count(&tables.IriscmsCategory{})
	if err != nil || total > 0 {
		return false, errors.New("模型已经被使用, 请删除使用分类后再执行删除操作")
	}
	if _, err := d.Orm.Transaction(func(session *xorm.Session) (i interface{}, err error) {
		i, err = d.Orm.ID(id).Delete(&tables.IriscmsDocumentModel{})
		if err != nil {
			golog.Error(err)
			return nil, err
		}
		if i == 0 {
			return nil, errors.New("删除了0条模型记录,错误表现") // 删除了0条记录, 返回失败
		}
		if !NewDocumentFieldDslModel(d.Orm).DeleteByMID(id) {
			golog.Errorf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id)
			return nil, errors.New(fmt.Sprintf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id))
		}
		return true, nil
	}); err != nil {
		return false, err
	}
	return true, nil
}
