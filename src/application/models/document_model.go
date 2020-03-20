package models

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

const (
	CUSTOM_TYPE = iota
	SYSTEM_TYPE
)

type DocumentModel struct {
	orm *xorm.Engine
}

func NewDocumentModel() *DocumentModel {
	return &DocumentModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (d *DocumentModel) GetList(page, limit int64) (list []tables.DocumentModel, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = d.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err != nil {
		pine.Logger().Error(err)
	}
	return list, total
}

func (d *DocumentModel) GetTableName(id int64) string {
	icache := di.MustGet(controllers.ServiceICache).(cache.ICache)
	key := fmt.Sprintf(controllers.CacheDocumentModelPrefix, id)
	tableName, _ := icache.Get(key)
	if len(tableName) == 0 {
		var detail = &tables.DocumentModel{}
		exists, _ := d.orm.ID(id).Get(detail)
		if !exists {
			return ""
		}
		tableName = []byte(detail.Table)
		icache.Set(key, tableName)
	}
	return string(tableName)
}

func (d *DocumentModel) GetByID(id int64) *tables.DocumentModel {
	var detail = &tables.DocumentModel{}
	if _, err := d.orm.ID(id).Get(detail); err != nil {
		pine.Logger().Error("document.model", err)
	}
	return detail
}

func (d *DocumentModel) DeleteByID(id int64) (bool, error) {
	total, err := d.orm.Where("model_id = ?", id).Count(&tables.Category{})
	if err != nil || total > 0 {
		return false, errors.New("模型已经被使用, 请删除使用分类后再执行删除操作")
	}
	if _, err := d.orm.Transaction(func(session *xorm.Session) (i interface{}, err error) {
		i, err = d.orm.ID(id).Delete(&tables.DocumentModel{})
		if err != nil {
			pine.Logger().Error(err)
			return nil, err
		}
		if i == 0 {
			return nil, errors.New("删除了0条模型记录,错误表现") // 删除了0条记录, 返回失败
		}
		if !NewDocumentFieldDslModel().DeleteByMID(id) {
			pine.Logger().Errorf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id)
			return nil, errors.New(fmt.Sprintf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id))
		}
		icache := di.MustGet(controllers.ServiceICache).(cache.ICache)
		key := fmt.Sprintf(controllers.CacheDocumentModelPrefix, id)
		icache.Delete(key)
		return true, nil
	}); err != nil {
		return false, err
	}
	return true, nil
}
