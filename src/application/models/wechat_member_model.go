package models

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type WechatMemberModel struct {
	orm *xorm.Engine
}

func NewWechatMemberModel() *WechatMemberModel {
	return &WechatMemberModel{orm: helper.GetORM()}
}

func (w *WechatMemberModel) GetList(page, limit int64) (list []tables.WechatMember, total int64) {
	offset := (page - 1) * limit
	total, _ = w.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if list == nil {
		list = []tables.WechatMember{}
	}
	return list, total
}

func (w *WechatMemberModel) GetInfo(id int64) tables.WechatMember {
	var member tables.WechatMember
	w.orm.Where("id = ?", id).Get(&member)
	return member
}
