package models

import (
	"errors"
	"fmt"
	"github.com/xiusin/pine/contracts"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
)

type DocumentModel struct {
	orm   *xorm.Engine
	cache contracts.Cache
}

func init() {
	di.Set(&DocumentModel{}, func(builder di.AbstractBuilder) (i any, err error) {
		return &DocumentModel{
			orm:   builder.MustGet(controllers.ServiceXorm).(*xorm.Engine),
			cache: builder.MustGet(controllers.ServiceICache).(contracts.Cache),
		}, nil
	}, true)
}

func NewDocumentModel() *DocumentModel {
	return di.MustGet(&DocumentModel{}).(*DocumentModel)
}

func (d *DocumentModel) GetList(page, limit int64) (list []tables.DocumentModel, total int64) {
	offset := (page - 1) * limit
	var err error
	total, err = d.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if err != nil {
		pine.Logger().Error(err.Error())
	}
	return list, total
}

func (d *DocumentModel) GetAllForBE() []tables.DocumentModel {
	var list []tables.DocumentModel
	err := d.cache.GetWithUnmarshal(controllers.CacheModels, &list)
	if err != nil || len(list) == 0 {
		d.orm.Find(&list)
		d.cache.SetWithMarshal(controllers.CacheModels, &list)
	}
	return list
}

func (d *DocumentModel) GetByID(id int64) *tables.DocumentModel {
	detail := &tables.DocumentModel{}
	if exists, _ := d.orm.ID(id).Get(detail); !exists {
		return nil
	}
	return detail
}

func (d *DocumentModel) GetByIDForBE(id int64) *tables.DocumentModel {
	models := d.GetAllForBE()
	for _, model := range models {
		if model.Id == id {
			return &model
		}
	}
	return nil
}

func (d *DocumentModel) GetWithTableNameForBE(name string) *tables.DocumentModel {
	models := d.GetAllForBE()
	for _, model := range models {
		if model.Table == name {
			return &model
		}
	}
	return nil
}

func (d *DocumentModel) DeleteByID(id int64) (bool, error) {
	total, err := d.orm.Where("model_id = ?", id).Count(&tables.Category{})
	if err != nil {
		pine.Logger().Error("checking for categories using model %d failed: %s", id, err.Error())
		return false, ErrInternal
	}
	if total > 0 {
		return false, ErrModelInUse
	}
	if _, err := d.orm.Transaction(func(session *xorm.Session) (i any, err error) {
		i, err = d.orm.ID(id).Delete(&tables.DocumentModel{})
		if err != nil {
			pine.Logger().Error("deleting document model %d failed: %s", id, err.Error())
			return nil, ErrInternal
		}
		if i == 0 {
			return nil, ErrModelNotFound
		}
		if !NewDocumentFieldDslModel().DeleteByMID(id) {
			pine.Logger().Error(fmt.Sprintf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id))
			return nil, ErrInternal
		}
		icache := di.MustGet(controllers.ServiceICache).(contracts.Cache)
		key := fmt.Sprintf(controllers.CacheDocumentModelPrefix, id)
		icache.Delete(key)
		return true, nil
	}); err != nil {
		return false, err
	}
	return true, nil
}
