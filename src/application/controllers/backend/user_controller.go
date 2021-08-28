package backend

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"strings"
)

type UserController struct {
	BaseController
}

func (c *UserController) Construct() {
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "username", Op: "LIKE", DataExp: "%$?%"},
		{Field: "email", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.Admin{}
	c.Entries = &[]*tables.Admin{}

	c.TableKey = "id"
	c.TableStructKey = "Userid"

	c.OpBefore = c.before
	c.OpAfter = c.after
}

func (c *UserController) before(opType int, param interface{}) error {
	if opType == OpAdd || opType == OpEdit {
		p, exist := param.(*tables.Admin), false
		if p.Userid > 0 {
			exist, _ = c.Orm.Table(p).Where("id <> ? and (username = ? or email = ?)", p.Userid, p.Username, p.Email).Exist()
		} else {
			exist, _ = c.Orm.Table(p).Where("username = ? or email = ?", p.Username, p.Email).Exist()
		}
		if exist {
			return errors.New("用户名或邮箱已存在")
		}
	}
	return nil
}

func (c *UserController) after(opType int, param interface{}) error {
	if opType == OpList {
		admins := c.Entries.(*[]*tables.Admin)
		roles := models.NewAdminRoleModel().All()
		for _, admin := range *admins {
			for _, roleId := range admin.RoleIdList {
				if r, ok := roles[roleId]; ok {
					admin.RoleName += r.Rolename + ","
				}
			}
			admin.RoleName = strings.TrimRight(admin.RoleName, ",")
		}
	}
	return nil
}
