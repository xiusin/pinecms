package models

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
)

type MenuModel struct {
	Orm *xorm.Engine
}

func NewMenuModel(orm *xorm.Engine) *MenuModel {
	return &MenuModel{Orm: orm}
}

//根据父级ID获取菜单列表 不递归
func (this *MenuModel) GetMenu(parentid, roleid int64) []tables.IriscmsMenu {
	menus := new([]tables.IriscmsMenu)
	this.Orm.Where("parentid = ? and display= ?", parentid, 1).Asc("listorder").Find(menus)
	if roleid == 1 {
		return *menus
	}
	retmenus := []tables.IriscmsMenu{}
	//结合角色权限进行菜单返回
	for _, menu := range *menus {
		total, _ := this.Orm.Where("c=? and a=? and roleid=?", menu.C, menu.A, roleid).Count(&tables.IriscmsAdminRolePriv{})
		//public的操作也要全部暴露
		if total > 0 || strings.Contains(menu.A, "public-") {
			retmenus = append(retmenus, menu)
		}
	}
	return retmenus
}

//当前位置
func (this MenuModel) CurrentPos(id int64) string {
	menu := tables.IriscmsMenu{Id: id}
	has, _ := this.Orm.Get(&menu)
	str := ""
	if !has {
		return ""
	}
	if menu.Parentid != 0 {
		str = this.CurrentPos(menu.Parentid)
	}
	return str + menu.Name + " > "
}

func (this MenuModel) GetTree(menus []tables.IriscmsMenu, parentid int64) []map[string]interface{} {
	res := []map[string]interface{}{}
	if len(menus) != 0 {
		for _, menu := range menus {
			if parentid == menu.Parentid {
				son := map[string]interface{}{
					"id":        menu.Id,
					"name":      menu.Name,
					"listorder": menu.Listorder,
					"operateid": menu.Id,
					"is_system": menu.IsSystem,
				}
				son["children"] = this.GetTree(menus, menu.Id)
				res = append(res, son)
			}
		}
	}
	return res
}

func (this MenuModel) GetAll() []tables.IriscmsMenu {
	menus := new([]tables.IriscmsMenu)
	this.Orm.Asc("listorder").Desc("id").Find(menus)
	return *menus
}

func (this MenuModel) GetRoleTree(parentid int64, roleid int64) []map[string]interface{} {
	menus := new([]tables.IriscmsMenu)
	//过滤我的面板
	err := this.Orm.Where("`parentid`=? AND `id`<>?", parentid, 1).Asc("listorder").Desc("id").Find(menus)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	res := []map[string]interface{}{}
	if len(*menus) != 0 {
		for _, v := range *menus {
			menu := map[string]interface{}{
				"id":   v.Id,
				"text": v.Name,
				"attributes": map[string]interface{}{
					"parent": this.GetParentIds(v.Id, ""),
				},
				"children": this.GetRoleTree(v.Id, roleid),
			}
			if len(menu["children"].([]map[string]interface{})) > 0 {
				menu["state"] = "closed"
			} else {
				//勾选默认菜单
				rolePriv := new([]tables.IriscmsAdminRolePriv)
				this.Orm.Where("c=? and a=? and roleid=?", v.C, v.A, roleid).Find(rolePriv)
				if len(*rolePriv) > 0 {
					menu["checked"] = true
				}
			}
			res = append(res, menu)
		}
	}
	return res
}

//获取菜单父级id
func (this MenuModel) GetParentIds(id int64, result string) string {
	menu := tables.IriscmsMenu{Id: id}
	has, _ := this.Orm.Get(&menu)
	var parentid int64 = 0
	if has {
		parentid = menu.Parentid
	}
	res := ""
	if parentid != 0 {
		if result == "" {
			res = strconv.Itoa(int(parentid))
		} else {
			res = "," + result
		}
		res = this.GetParentIds(parentid, res)
	}
	return res
}

//检查菜单名称是否存在
func (this MenuModel) CheckName(name string) bool {
	has, _ := this.Orm.Get(&tables.IriscmsMenu{Name: name})
	return has
}

//
func (this MenuModel) GetInfo(id int64) (*tables.IriscmsMenu, bool) {
	im := &tables.IriscmsMenu{Id: id}
	has, _ := this.Orm.Get(im)
	return im, has
}

//获取selectTree
func (this MenuModel) GetSelectTree(menus []tables.IriscmsMenu, parentid int64) []map[string]interface{} {
	res := []map[string]interface{}{}
	if len(menus) != 0 {
		for _, v := range menus {
			if parentid == v.Parentid {
				menu := map[string]interface{}{
					"id":       v.Id,
					"text":     v.Name,
					"children": this.GetSelectTree(menus, v.Id),
				}
				res = append(res, menu)
			}
		}
	}
	return res
}
