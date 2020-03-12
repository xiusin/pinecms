package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pine/di"
)

type MemberModel struct {
	orm *xorm.Engine
}

func NewMemberModel() *MemberModel {
	return &MemberModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (m *MemberModel) GetList(page, limit int64) (list []tables.Member, total int64) {
	offset := (page - 1) * limit
	total, _ = m.orm.Limit(int(limit), int(offset)).FindAndCount(&list)
	return list, total
}

func (m *MemberModel) GetInfo(id int64) *tables.Member {
	var member tables.Member
	m.orm.ID(id).Get(&member)
	return &member
}

func (m *MemberModel) Edit(id int64, members *tables.Member) bool {
	res,err := m.orm.ID(id).MustCols("enabled").Update(members)
	if err != nil {
		m.orm.Logger().Error(helper.GetCallerFuncName(), err)
		return false
	}
	return res > 0
}

