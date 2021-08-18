package models

import (
	"errors"
	"github.com/xiusin/pine/di"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type AdminModel struct {
	orm *xorm.Engine
}

func NewAdminModel() *AdminModel {
	return &AdminModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

// Login 登录用户
func (a *AdminModel) Login(username, password, ip string) (tables.Admin, error) {
	admin := tables.Admin{Username: username}
	exist, _ := a.orm.Get(&admin)
	if !exist {
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
func (a *AdminModel) GetUserInfo(userid int64) (tables.Admin, error) {
	admin := tables.Admin{Userid: userid}
	has, _ := a.orm.Get(&admin)
	if has {

		return admin, nil
	}
	return admin, errors.New("没有找到用户")
}

//编辑密码
func (a *AdminModel) EditPassword(userid int64, password string) bool {
	admin := tables.Admin{Userid: userid}
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
func (a *AdminModel) GetList(where string, page, rows int, order string, sortType string) []tables.Admin {
	start := (page - 1) * rows
	admins := []tables.Admin{}
	if sortType == "asc" {
		a.orm.Where(where).Asc(order).Limit(rows, start).Find(&admins)
	} else {
		a.orm.Where(where).Desc(order).Limit(rows, start).Find(&admins)
	}
	return admins
}
