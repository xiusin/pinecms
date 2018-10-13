package models

import (
	"iriscms/application/models/tables"

	"github.com/go-xorm/xorm"
)

type MemberModel struct {
	Orm *xorm.Engine
}

func NewMemberModel(orm *xorm.Engine) *MemberModel {
	return &MemberModel{Orm: orm}
}

func (this *MemberModel) GetList(page, limit int64) (list []tables.IriscmsMember, total int64) {
	offset := (page - 1) * limit
	total, _ = this.Orm.Limit(int(limit),int(offset)).FindAndCount(&list)
	return list, total
}

