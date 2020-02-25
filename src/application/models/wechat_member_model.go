package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
)

type WechatMemberModel struct {
	Orm *xorm.Engine
}

func NewWechatMemberModel(orm *xorm.Engine) *WechatMemberModel {
	return &WechatMemberModel{Orm: orm}
}

func (w *WechatMemberModel) GetList(page, limit int64) (list []tables.IriscmsWechatMember, total int64) {
	offset := (page - 1) * limit
	total, _ = w.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if list == nil {
		list = []tables.IriscmsWechatMember{}
	}
	return list, total
}

func (w *WechatMemberModel) GetInfo(id int64) tables.IriscmsWechatMember {
	var member tables.IriscmsWechatMember
	w.Orm.Where("id = ?", id).Get(&member)
	return member
}
