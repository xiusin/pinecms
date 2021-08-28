package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/builder"
)

type MemberController struct {
	BaseController
}

func (c *MemberController) Construct() {
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Table = &tables.Member{}
	c.Entries = &[]tables.Member{}
	c.ApiEntityName = "会员"
	c.Group = "会员管理"
	c.BaseController.Construct()
	c.OpBefore = c.before
}

func (c *MemberController) before(act int, params interface{}) error {
	switch act {
	case OpAdd:
		data := c.Table.(*tables.Member)
		if exist, _ := c.Orm.Table(c.Table).Where("account = ?", data.Account).Or("email = ?", data.Email).Exist(); exist {
			return errors.New("账号或邮箱已存在")
		}

	case OpEdit:
		data := c.Table.(*tables.Member)
		if exist, _ := c.Orm.Table(c.Table).Where("id <> ?", data.Id).
			Where(builder.Eq{"account": data.Account}.Or(builder.Eq{"email": data.Email})).Exist(); exist {
			return errors.New("账号或邮箱已存在")
		}
	}
	return nil
}
