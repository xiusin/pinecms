package models

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

const (
	IMG_TYPE  = "img"
	FILE_TYPE = "file"
)

type AttachmentsModel struct {
	orm *xorm.Engine
}

func NewAttachmentsModel() *AttachmentsModel {
	return &AttachmentsModel{orm: helper.GetORM()}
}

func (a *AttachmentsModel) GetList(keywords string, page, limit int64) (list []tables.Attachments, total int64) {
	offset := (page - 1) * limit
	var err error
	sess := a.orm.Limit(int(limit), int(offset)).Desc("id")
	if len(keywords) != 0 {
		likePrtten := "%" + keywords + "%"
		sess.Where("origin_name like ?", likePrtten).Or("name like ?", likePrtten)
	}
	total, err = sess.FindAndCount(&list)
	if err != nil {
		pine.Logger().Error(err)
	}
	if list == nil {
		list = []tables.Attachments{}
	}
	return list, total
}

func (a *AttachmentsModel) Delete(id int64) bool {
	res, _ := a.orm.ID(id).Delete(&tables.Attachments{})
	return res > 0
}
