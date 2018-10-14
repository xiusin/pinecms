package models

import (
	"iriscms/application/models/tables"

	"github.com/go-xorm/xorm"
)

type WechatMemberModel struct {
	Orm *xorm.Engine
}

func NewWechatMemberModel(orm *xorm.Engine) *WechatMemberModel {
	return &WechatMemberModel{Orm: orm}
}

func (this *WechatMemberModel) GetList(page, limit int64) (list []tables.IriscmsWechatMember, total int64) {
	offset := (page - 1) * limit
	total, _ = this.Orm.Limit(int(limit),int(offset)).FindAndCount(&list)
	return list, total
}

func (this *WechatMemberModel) GetInfo(id int64)  tables.IriscmsWechatMember {
	var member tables.IriscmsWechatMember
	this.Orm.Where("id = ?", id).Get(&member)
	return member
}

