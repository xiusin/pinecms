package models

import (
	"errors"
	"log"

	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type AdminModel struct {
	Orm *xorm.Engine
}

func NewAdminModel(orm *xorm.Engine) *AdminModel {
	return &AdminModel{Orm: orm}
}

//登录用户
func (this AdminModel) Login(username, password, ip string) (tables.IriscmsAdmin, error) {
	admin := tables.IriscmsAdmin{Username: username}
	has, _ := this.Orm.Get(&admin)
	if !has {
		return admin, errors.New("管理员不存在")
	}
	password = helper.Password(password, admin.Encrypt)
	if password != admin.Password {
		return admin, errors.New("密码错误，请再次尝试！")
	}
	admin.Lastloginip = ip
	this.Orm.Id(admin.Userid).Update(admin)
	return admin, nil
}

//获取用户信息
func (this AdminModel) GetUserInfo(userid int64) (tables.IriscmsAdmin, error) {
	admin := tables.IriscmsAdmin{Userid: userid}
	has, _ := this.Orm.Get(&admin)
	if has {

		return admin, nil
	}
	return admin, errors.New("没有找到用户")
}

//编辑密码
func (this AdminModel) EditPassword(userid int64, password string) bool {
	admin := tables.IriscmsAdmin{Userid: userid}
	res, _ := this.Orm.Get(&admin)
	encrypt := string(helper.Krand(8, 3))
	if res {
		admin.Password = helper.Password(password, encrypt)
		admin.Encrypt = encrypt
		result, _ := this.Orm.Id(admin.Userid).Update(&admin)
		if result == 0 {
			return false
		} else {
			return true
		}
	}
	return false
}

//获取管理员列表
func (this AdminModel) GetList(where string, page, rows int, order string, sortType string) []tables.IriscmsAdmin {
	start := (page - 1) * rows
	admins := []tables.IriscmsAdmin{}
	if sortType == "asc" {
		this.Orm.Where(where).Asc(order).Limit(rows, start).Find(&admins)
	} else {
		this.Orm.Where(where).Desc(order).Limit(rows, start).Find(&admins)
	}
	return admins
}

func (this AdminModel) GetRoleList(where string, page, rows int) []tables.IriscmsAdminRole {
	start := (page - 1) * rows
	myroles := []tables.IriscmsAdminRole{}
	err := this.Orm.Where(where).Limit(rows, start).Find(&myroles)
	if err != nil {
		log.Println(err.Error())
	}
	return myroles
}

func (this AdminModel) CheckRoleName(rolename string) bool {
	role, err := this.Orm.Get(&tables.IriscmsAdminRole{Rolename: rolename})
	if err != nil {
		return false
	}
	return role
}

func (this AdminModel) AddRole(rolename, description string, disabled, listorder int64) bool {
	new_role := tables.IriscmsAdminRole{
		Rolename:    rolename,
		Description: description,
		Disabled:    disabled,
		Listorder:   listorder,
	}
	insertId, err := this.Orm.Insert(&new_role)
	if err != nil || insertId == 0 {
		log.Println(err, insertId)
		return false
	}
	return true
}

func (this AdminModel) GetRoleById(id int64) (tables.IriscmsAdminRole, error) {
	role := tables.IriscmsAdminRole{Roleid: id}
	_, err := this.Orm.Get(&role)
	return role, err
}

func (this AdminModel) UpdateRole(role tables.IriscmsAdminRole) bool {
	res, err := this.Orm.Where("roleid=?", role.Roleid).Update(&role)
	if err != nil || res == 0 {
		log.Println("AdminModel::UpdateRole", err, res)
		return false
	}
	return true
}

func (this AdminModel) DeleteRole(role tables.IriscmsAdminRole) bool {
	res, err := this.Orm.Delete(&role)
	if err != nil || res == 0 {
		log.Println("AdminModel::DeleteRole", err, res)
		return false
	}
	return true
}

func (this AdminModel) HasAdminByRoleId(roleid int64) bool {
	res, err := this.Orm.Count(&tables.IriscmsAdmin{Roleid: roleid})
	if err != nil || res == 0 {
		log.Println("AdminModel::DeleteRole", err, res)
		return false
	}
	return true
}
