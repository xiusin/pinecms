package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type MemberModel struct {
	Orm *xorm.Engine
}

func NewMemberModel(orm *xorm.Engine) *MemberModel {
	return &MemberModel{Orm: orm}
}

func (m *MemberModel) GetList(page, limit int64) (list []tables.IriscmsMember, total int64) {
	offset := (page - 1) * limit
	total, _ = m.Orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}

func (m *MemberModel) GetInfo(id int64) *tables.IriscmsMember {
	var member tables.IriscmsMember
	m.Orm.ID(id).Get(&member)
	return &member
}

func (m *MemberModel) Edit(id int64, members *tables.IriscmsMember) bool {
	res,err := m.Orm.ID(id).MustCols("enabled").Update(members)
	if err != nil {
		m.Orm.Logger().Error(helper.GetCallerFuncName(), err)
		return false
	}
	return res > 0
}

