package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/pine/di"
)

type WechatMemberModel struct {
	orm *xorm.Engine
}

func NewWechatMemberModel() *WechatMemberModel {
	return &WechatMemberModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (w *WechatMemberModel) GetList(page, limit int64) (list []tables.IriscmsWechatMember, total int64) {
	offset := (page - 1) * limit
	total, _ = w.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	if list == nil {
		list = []tables.IriscmsWechatMember{}
	}
	return list, total
}

func (w *WechatMemberModel) GetInfo(id int64) tables.IriscmsWechatMember {
	var member tables.IriscmsWechatMember
	w.orm.Where("id = ?", id).Get(&member)
	return member
}
