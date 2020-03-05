package models

import (
	"errors"
	"github.com/xiusin/pine/di"
	"log"

	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type AdminModel struct {
	orm *xorm.Engine
}


func NewAdminModel() *AdminModel {
	return &AdminModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

//登录用户
func (a *AdminModel) Login(username, password, ip string) (tables.IriscmsAdmin, error) {
	admin := tables.IriscmsAdmin{Username: username}
	has, _ := a.orm.Get(&admin)
	if !has {
		return admin, errors.New("管理员不存在")
	}
	password = helper.Password(password, admin.Encrypt)
	if password != admin.Password {
		return admin, errors.New("密码错误，请再次尝试！")
	}
	admin.Lastloginip = ip
	a.orm.Id(admin.Userid).Update(admin)
	return admin, nil
}

//获取用户信息
func (a *AdminModel) GetUserInfo(userid int64) (tables.IriscmsAdmin, error) {
	admin := tables.IriscmsAdmin{Userid: userid}
	has, _ := a.orm.Get(&admin)
	if has {

		return admin, nil
	}
	return admin, errors.New("没有找到用户")
}

//编辑密码
func (a *AdminModel) EditPassword(userid int64, password string) bool {
	admin := tables.IriscmsAdmin{Userid: userid}
	res, _ := a.orm.Get(&admin)
	encrypt := string(helper.Krand(8, 3))
	if res {
		admin.Password = helper.Password(password, encrypt)
		admin.Encrypt = encrypt
		result, _ := a.orm.Id(admin.Userid).Update(&admin)
		if result == 0 {
			return false
		} else {
			return true
		}
	}
	return false
}

//获取管理员列表
func (a *AdminModel) GetList(where string, page, rows int, order string, sortType string) []tables.IriscmsAdmin {
	start := (page - 1) * rows
	admins := []tables.IriscmsAdmin{}
	if sortType == "asc" {
		a.orm.Where(where).Asc(order).Limit(rows, start).Find(&admins)
	} else {
		a.orm.Where(where).Desc(order).Limit(rows, start).Find(&admins)
	}
	return admins
}

func (a *AdminModel) GetRoleList(where string, page, rows int) []tables.IriscmsAdminRole {
	start := (page - 1) * rows
	myroles := []tables.IriscmsAdminRole{}
	err := a.orm.Where(where).Limit(rows, start).Find(&myroles)
	if err != nil {
		log.Println(err.Error())
	}
	return myroles
}

func (a *AdminModel) CheckRoleName(rolename string) bool {
	role, err := a.orm.Get(&tables.IriscmsAdminRole{Rolename: rolename})
	if err != nil {
		return false
	}
	return role
}

func (a *AdminModel) AddRole(rolename, description string, disabled, listorder int64) bool {
	newRole := tables.IriscmsAdminRole{
		Rolename:    rolename,
		Description: description,
		Disabled:    disabled,
		Listorder:   listorder,
	}
	insertId, err := a.orm.Insert(&newRole)
	if err != nil || insertId == 0 {
		log.Println(err, insertId)
		return false
	}
	return true
}

func (a *AdminModel) GetRoleById(id int64) (tables.IriscmsAdminRole, error) {
	role := tables.IriscmsAdminRole{Roleid: id}
	_, err := a.orm.Get(&role)
	return role, err
}

func (a *AdminModel) UpdateRole(role tables.IriscmsAdminRole) bool {
	res, err := a.orm.Where("roleid=?", role.Roleid).Update(&role)
	if err != nil || res == 0 {
		log.Println("AdminModel::UpdateRole", err, res)
		return false
	}
	return true
}

func (a *AdminModel) DeleteRole(role tables.IriscmsAdminRole) bool {
	res, err := a.orm.Delete(&role)
	if err != nil || res == 0 {
		log.Println("AdminModel::DeleteRole", err, res)
		return false
	}
	return true
}

func (a *AdminModel) HasAdminByRoleId(roleid int64) bool {
	res, err := a.orm.Count(&tables.IriscmsAdmin{Roleid: roleid})
	if err != nil || res == 0 {
		log.Println("AdminModel::DeleteRole", err, res)
		return false
	}
	return true
}
