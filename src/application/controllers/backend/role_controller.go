package backend

import (
	"errors"
	"fmt"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type AdminRoleController struct {
	BaseController
}

func (c *AdminRoleController) Construct() {
	c.Group = "角色管理"
	c.ApiEntityName = "角色"
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "rolename", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Orm = helper.GetORM()
	c.Table = &tables.AdminRole{}
	c.Entries = &[]tables.AdminRole{}
	c.BaseController.Construct()
	c.OpBefore = c.before
	c.OpAfter = c.after
}

func (c *AdminRoleController) before(opType int, param interface{}) error {
	if opType == OpDel {
		p := param.(*idParams)
		for _, id := range p.Ids {
			if 1 == id {
				return errors.New("不可删除超级管理员角色")
			}
		}
	} else if opType == OpEdit {
		if c.Ctx().Value("roleid").(int64) != 1 {
			return errors.New("您的角色无法修改超级管理员信息")
		}
	}
	return nil
}

func (c *AdminRoleController) PostAdd() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	if _, err := c.Orm.Transaction(func(session *xorm.Session) (interface{}, error) {
		if err := c.add(); err != nil  {
			return nil, err
		}
		t := c.Table.(*tables.AdminRole)
		id := t.Id
		if err := c.perm(id, t.MenuIdList); err != nil {
			return nil, err
		}
		return nil, nil
	}); err == nil {
		helper.Ajax("新增数据成功", 0, c.Ctx())
	} else {
		helper.Ajax(err, 1, c.Ctx())
	}
}

func (c *AdminRoleController) PostEdit() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	c.edit()
	t := c.Table.(*tables.AdminRole)
	id := t.Id
	c.perm(id, t.MenuIdList)
	helper.Ajax("修改数据成功", 0, c.Ctx())
}

func (c *AdminRoleController) after(opType int, param interface{}) error {
	switch opType {
	case OpDel:
		p := param.(*idParams)
		_, err := c.Orm.In("role_id", p.Ids).Delete(&tables.AdminRolePriv{})
		return err
	case OpInfo:
		t := c.Table.(*tables.AdminRole)
		var priv []tables.AdminRolePriv
		c.Orm.Where("role_id = ?", t.Id).Find(&priv)
		var menuIds = make([]int64, len(priv))
		for _, i2 := range priv {
			menuIds = append(menuIds, i2.MenuId)
		}
		t.MenuIdList = menuIds
		c.Table = t
	}
	return nil
}

func (c *AdminRoleController) perm(roleid int64, menuIds []int64) error {
	_, err := c.Orm.Where("role_id=?", roleid).Delete(&tables.AdminRolePriv{})
	if err != nil {
		return err
	}
	if len(menuIds) == 0 {
		return nil
	}
	var inserts []tables.AdminRolePriv
	for _, v := range menuIds {
		if v == 0 {
			continue
		}
		menu := tables.Menu{Id: v}
		exist, err := di.MustGet(&xorm.Engine{}).(*xorm.Engine).Get(&menu)
		if err != nil || !exist {
			continue
		}

		inserts = append(inserts, tables.AdminRolePriv{RoleId: roleid, MenuId: v})
	}
	res, err := c.Orm.Insert(inserts)
	if int(res) != len(menuIds) {
		return fmt.Errorf("更新权限错误: 更新%d,成功%d", len(menuIds), res)
	}
	return nil
}
