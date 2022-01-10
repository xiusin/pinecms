package models

import (
	"errors"
	"fmt"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
)

type DocumentModel struct {
	orm   *xorm.Engine
	cache cache.AbstractCache
}

func init() {
	di.Set(&DocumentModel{}, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return &DocumentModel{
			orm:   builder.MustGet("*xorm.Engine").(*xorm.Engine),
			cache: builder.MustGet("cache.AbstractCache").(cache.AbstractCache),
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
		pine.Logger().Error(err)
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
			return nil, fmt.Errorf("删除数据模型ID: %d 成功, 删除关联字段失败, 回滚数据", id)
		}
		icache := di.MustGet(controllers.ServiceICache).(cache.AbstractCache)
		key := fmt.Sprintf(controllers.CacheDocumentModelPrefix, id)
		icache.Delete(key)
		return true, nil
	}); err != nil {
		return false, err
	}
	return true, nil
}
