package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type MemberGroupController struct {
	BaseController
}

func (c *MemberGroupController) Construct() {
	c.Table = &tables.MemberGroup{}
	c.Entries = &[]tables.MemberGroup{}
	c.ApiEntityName = "会员分组"
	c.Group = "会员分组管理"
	c.BaseController.Construct()
	c.OpBefore = c.before
}

func (c *MemberGroupController) before(act int, params interface{}) error {
	switch act {
	case OpAdd:
		data := c.Table.(*tables.MemberGroup)
		if exist, _ := c.Orm.Table(c.Table).Where("name = ?", data.Name).Exist(); exist {
			return errors.New("分组已存在")
		}

	case OpEdit:
		data := c.Table.(*tables.MemberGroup)
		if exist, _ := c.Orm.Table(c.Table).Where("id <> ?", data.Id).Where("name = ?", data.Name).Exist(); exist {
			return errors.New("分组已存在")
		}
	}
	return nil
}

func (c *MemberGroupController) GetSelect()  {
	c.Orm.Find(c.Entries)
	helper.Ajax(c.Entries, 0, c.Ctx())
}
