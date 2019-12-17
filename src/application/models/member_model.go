package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
)

type MemberModel struct {
	Orm *xorm.Engine
}

func NewMemberModel(orm *xorm.Engine) *MemberModel {
	return &MemberModel{Orm: orm}
}

func (this *MemberModel) GetList(page, limit int64) (list []tables.IriscmsMember, total int64) {
	offset := (page - 1) * limit
	total, _ = this.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}

func (this *MemberModel) GetInfo(id int64) tables.IriscmsMember {
	var member tables.IriscmsMember
	this.Orm.Where("id = ?", id).Get(&member)
	return member
}
