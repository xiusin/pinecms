package models

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
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

func (m *MemberModel) Add(members *tables.Member) int64 {
	res, err := m.orm.InsertOne(members)
	if err != nil {
		pine.Logger().Error(err)
		return 0
	}
	return res
}

func (m *MemberModel) Exist(account string) bool {
	count, _ := m.orm.Where("account = ?", account).Count(&tables.Member{})
	return count > 0
}

func (m *MemberModel) Edit(id int64, members *tables.Member) bool {
	res, err := m.orm.ID(id).MustCols("enabled", "integral").Update(members)
	if err != nil {
		return false
	}
	return res > 0
}
