package backend

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"
)

type AdminController struct {
	Orm     *xorm.Engine
	Ctx     iris.Context //存在则自动绑定
	Cache   cache.ICache
	Session *sessions.Session
}

func (c *AdminController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/admin/memberlist", "Memberlist")
	b.Handle("ANY", "/admin/public-editpwd", "PublicEditpwd")
	b.Handle("ANY", "/admin/public-editinfo", "PublicEditInfo")
	b.Handle("POST", "/admin/public-checkEmail", "PublicCheckEmail")
	b.Handle("POST", "/admin/public-checkName", "PubicCheckName")
	b.Handle("POST", "/admin/check-password", "PublicCheckPassword")
	b.Handle("ANY", "/admin/member-add", "MemberAdd")
	b.Handle("ANY", "/admin/member-edit", "MemberEdit")
	b.Handle("ANY", "/admin/member-delete", "MemberDelete")
	b.Handle("ANY", "/admin/rolelist", "RoleList")
	b.Handle("ANY", "/admin/role-add", "RoleAdd")
	b.Handle("ANY", "/admin/role-edit", "RoleEdit")
	b.Handle("ANY", "/admin/role-delete", "RoleDelete")
	b.Handle("ANY", "/admin/role-permission", "RolePermission")
	b.Handle("ANY", "/admin/check-rolename", "PublicCheckRoleName")

}

func (c *AdminController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
	}
}

//用于用户列表数据格式返回
type memberlist struct {
	Email         string
	Lastloginip   string
	Lastlogintime string
	Realname      string
	Rolename      string
	Userid        int64
	Username      string
}

func (c *AdminController) PublicEditInfo() {
	aid, _ := c.Ctx.Values().GetInt64("adminid") //检测是否设置过session
	if c.Ctx.Method() == "POST" {
		info := tables.IriscmsAdmin{
			Userid: aid,
		}
		has, _ := c.Orm.Get(&info) //读取用户资料
		if !has {
			helper.Ajax("用户资料已经不存在", 1, c.Ctx)
		} else {
			info.Realname = c.Ctx.PostValue("realname")
			info.Email = c.Ctx.PostValue("email")
			res, err := c.Orm.Id(aid).Update(info)
			if err != nil {
				helper.Ajax("修改资料失败"+err.Error(), 1, c.Ctx)
			} else {
				if res > 0 {
					helper.Ajax("修改资料成功", 0, c.Ctx)
				} else {
					helper.Ajax("修改资料失败", 1, c.Ctx)
				}
			}
		}
		return
	}
	menuid, _ := c.Ctx.URLParamInt64("menuid")
	currentPos := models.NewMenuModel(c.Orm).CurrentPos(menuid)
	info, err := models.NewAdminModel(c.Orm).GetUserInfo(aid)
	if err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx)
		return
	}
	c.Ctx.ViewData("title", currentPos)
	c.Ctx.ViewData("info", info)
	if err := c.Ctx.View("backend/admin_editinfo.html"); err != nil {
		c.Ctx.WriteString(err.Error())
	}
}

func (c *AdminController) Memberlist() {
	act := c.Ctx.URLParam("grid")
	if act == "datagrid" {
		page, err := c.Ctx.URLParamInt("page")
		orderField := c.Ctx.URLParam("sort")
		if orderField == "" {
			orderField = "userid"
		}
		orderType := c.Ctx.URLParam("sort")
		if orderType == "" {
			orderType = "desc"
		}

		if err != nil {
			page = 1
		}
		data := models.NewAdminModel(c.Orm).GetList("1", page, 10, orderField, orderType)
		retData := []memberlist{}
		//将数据以map的方式返回吧.
		for _, v := range data {
			item := memberlist{
				Email:         v.Email,
				Lastloginip:   v.Lastloginip,
				Lastlogintime: helper.FormatTime(v.Lastlogintime),
				Realname:      v.Realname,
				Rolename:      "",
				Userid:        v.Userid,
				Username:      v.Username,
			}
			roleInfo, err := models.NewAdminModel(c.Orm).GetRoleById(int64(v.Roleid))
			if err != nil {
				roleInfo.Rolename = ""
			}
			item.Rolename = roleInfo.Rolename
			retData = append(retData, item)
		}

		c.Ctx.JSON(map[string]interface{}{
			"total": len(retData),
			"rows":  retData,
		})
		return
	}
	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("admin_member_list_datagrid", "/b/admin/memberlist?grid=datagrid", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "admin_memberlist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"用户名":    {"field": "Username", "width": "15", "sortable": "true", "index": "0"},
		"所属角色":   {"field": "Rolename", "width": "15", "sortable": "true", "index": "1"},
		"最后登录IP": {"field": "Lastloginip", "width": "15", "sortable": "true", "index": "2"},
		"最后登录时间": {"field": "Lastlogintime", "width": "15", "sortable": "true", "formatter": "adminMemberListTimeFormatter", "index": "3"},
		"E-mail": {"field": "Email", "width": "25", "sortable": "true", "index": "4"},
		"真实姓名":   {"field": "Realname", "width": "15", "sortable": "true", "index": "5"},
		"管理操作":   {"field": "Userid", "width": "15", "sortable": "true", "formatter": "adminMemberListOperateFormatter", "index": "6"},
	})
	c.Ctx.ViewData("table", template.HTML(table))
	c.Ctx.View("backend/admin_memberlist.html")
}

func (c *AdminController) PublicEditpwd() {
	aid, _ := c.Ctx.Values().GetInt64("adminid")
	menuid, _ := c.Ctx.URLParamInt64("menuid")
	info := tables.IriscmsAdmin{Userid: int64(aid)}
	has, _ := c.Orm.Get(&info)
	if !has {
		c.Ctx.WriteString("没有找到")
		return
	}
	if c.Ctx.Method() == "POST" {
		if info.Password != helper.Password(c.Ctx.PostValue("old_password"), info.Encrypt) {
			helper.Ajax("原密码错误", 1, c.Ctx)
			return
		}
		info.Password = helper.Password(c.Ctx.PostValue("new_password"), info.Encrypt)
		res, _ := c.Orm.Id(aid).Update(info)
		if res > 0 {
			helper.Ajax("修改资料成功", 0, c.Ctx)
		} else {
			helper.Ajax("修改资料失败", 1, c.Ctx)
		}
		return
	}
	c.Ctx.ViewData("currentpos", models.NewMenuModel(c.Orm).CurrentPos(menuid))
	c.Ctx.ViewData("admin", info)
	c.Ctx.View("backend/admin_editpwd.html")
}

func (c *AdminController) PublicCheckEmail() {
	info := &tables.IriscmsAdmin{Username: c.Ctx.FormValue("name")}
	has, _ := c.Orm.Get(info)
	if !has {
		helper.Ajax("没有相同的用户名", 0, c.Ctx)
	} else {
		helper.Ajax("已经有相同的用户名,请换一个再试", 1, c.Ctx)
	}
}
func (c *AdminController) PublicCheckPassword() {
	aid, _ := c.Ctx.Values().Get("adminid").(int64)
	password := c.Ctx.FormValue("password")
	admin, err := models.NewAdminModel(c.Orm).GetUserInfo(aid)
	if err != nil {
		helper.Ajax("无法查找到相关信息", 1, c.Ctx)
		return
	}
	if admin.Password != helper.Password(password, admin.Encrypt) {
		helper.Ajax("旧密码错误", 1, c.Ctx)
		return
	}
	helper.Ajax("验证密码成功", 0, c.Ctx)
}
func (c *AdminController) PubicCheckName() {
	info := &tables.IriscmsAdmin{Username: c.Ctx.FormValue("name")}
	uid, _ := c.Ctx.URLParamInt64("id")
	has, _ := c.Orm.Get(info)
	if !has || info.Userid == uid {
		helper.Ajax("没有相同的用户名", 0, c.Ctx)
	} else {
		helper.Ajax("已经有相同的用户名,请换一个再试", 1, c.Ctx)
	}
}

func (c *AdminController) PublicCheckRoleName() {
	rolename := c.Ctx.FormValue("rolename")
	if !helper.IsAjax(c.Ctx) || rolename == "" {
		helper.Ajax("参数错误 ,"+rolename, 1, c.Ctx)
		return
	}
	defaultName := c.Ctx.FormValue("default")
	if defaultName != "" && rolename == defaultName {
		helper.Ajax("角色已存在", 1, c.Ctx)
		return
	}
	if models.NewAdminModel(c.Orm).CheckRoleName(rolename) {
		helper.Ajax("角色已存在", 1, c.Ctx)
		return
	}
	helper.Ajax("通过", 0, c.Ctx)
}
func (c *AdminController) MemberAdd() {
	if c.Ctx.FormValue("act") != "" {
		if c.Ctx.FormValue("pwdconfirm") != c.Ctx.FormValue("password") || c.Ctx.FormValue("password") == "" {
			helper.Ajax("两次密码不一致", 1, c.Ctx)
			return
		}
		if c.Ctx.FormValue("roleid") == "" {
			helper.Ajax("请选择角色", 1, c.Ctx)
			return
		}
		roleid, err := strconv.Atoi(c.Ctx.FormValue("roleid"))
		if err != nil {
			helper.Ajax("角色信息错误", 1, c.Ctx)
			return
		}
		str := string(helper.Krand(6, 3))
		newAdmin := &tables.IriscmsAdmin{
			Username: c.Ctx.FormValue("username"),
			Password: helper.Password(c.Ctx.FormValue("password"), str),
			Email:    c.Ctx.FormValue("email"),
			Encrypt:  str,
			Realname: c.Ctx.FormValue("realname"),
			Roleid:   int64(roleid),
		}
		id, err := c.Orm.Insert(newAdmin)
		if id > 0 {
			helper.Ajax("添加管理员成功", 0, c.Ctx)
			return
		}
		helper.Ajax("添加管理员失败", 1, c.Ctx)
		return
	}
	roles := models.NewAdminModel(c.Orm).GetRoleList("1", 1, 1000)
	c.Ctx.ViewData("Roles", roles)
	c.Ctx.View("backend/member_add.html")
}
func (c *AdminController) MemberEdit() {
	adminid, err := c.Ctx.URLParamInt64("id")
	if err != nil {
		c.Ctx.WriteString("参数错误 : " + err.Error())
		return
	}
	info, err := models.NewAdminModel(c.Orm).GetUserInfo(adminid)
	if err != nil {
		c.Ctx.WriteString("没有该管理员信息")
		return
	}
	if c.Ctx.Method() == "POST" {
		if c.Ctx.FormValue("password") != "" {
			if c.Ctx.FormValue("pwdconfirm") != c.Ctx.FormValue("password") {
				helper.Ajax("两次密码不一致", 1, c.Ctx)
				return
			}
		}
		if c.Ctx.FormValue("roleid") == "" {
			helper.Ajax("请选择角色", 1, c.Ctx)
			return
		}
		roleid, err := strconv.Atoi(c.Ctx.FormValue("roleid"))
		if err != nil {
			helper.Ajax("角色信息错误", 1, c.Ctx)
			return
		}
		info.Username = c.Ctx.FormValue("username")
		info.Email = c.Ctx.FormValue("email")
		info.Realname = c.Ctx.FormValue("realname")
		info.Roleid = int64(roleid)
		if c.Ctx.FormValue("password") != "" {
			info.Password = helper.Password(c.Ctx.PostValue("password"), info.Encrypt)
		}
		res, err := c.Orm.Where("userid = ?", info.Userid).Update(info)
		if err != nil {
			helper.Ajax(err.Error(), 0, c.Ctx)
			return
		}
		if res > 0 {
			helper.Ajax("修改管理员成功", 0, c.Ctx)
			return
		}
		helper.Ajax("修改管理员失败", 1, c.Ctx)
		return
	}

	roles := models.NewAdminModel(c.Orm).GetRoleList("1", 1, 1000)
	c.Ctx.ViewData("Roles", roles)
	c.Ctx.ViewData("Info", info)
	c.Ctx.View("backend/member_edit.html")
}
func (c *AdminController) MemberDelete() {
	id, err := strconv.Atoi(c.Ctx.FormValue("id"))
	if err != nil || helper.IsFalse(id) || id == 1 {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	deleteAdmin := &tables.IriscmsAdmin{Userid: int64(id)}
	res, err := c.Orm.Delete(deleteAdmin)
	if err != nil || helper.IsFalse(res) {
		helper.Ajax("删除失败", 1, c.Ctx)
		return
	}
	helper.Ajax("删除成功", 0, c.Ctx)
}

func (c *AdminController) RoleList() {
	menuid, _ := c.Ctx.URLParamInt64("menuid")
	if c.Ctx.URLParam("grid") == "datagrid" {
		page, err := c.Ctx.URLParamInt("page")
		orderField := c.Ctx.URLParam("sort")
		if orderField == "" {
			orderField = "id"
		}
		orderType := c.Ctx.URLParam("sort")
		if orderType == "" {
			orderType = "desc"
		}

		if err != nil {
			page = 1
		}

		data := models.NewAdminModel(c.Orm).GetRoleList("1", page, 1000)
		c.Ctx.JSON(map[string]interface{}{
			"total": len(data),
			"rows":  data,
		})
		return
	}

	datagrid := helper.Datagrid("admin_rolelist_list_datagrid", "/b/admin/rolelist?grid=datagrid", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "admin_rolelist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"角色名称": {"field": "Rolename", "width": "15", "index": "0"},
		"角色描述": {"field": "Description", "width": "15", "index": "1"},
		"管理操作": {"field": "Roleid", "width": "15", "formatter": "adminRoleListOperateFormatter", "index": "2"},
	})
	c.Ctx.ViewData("datagrid", template.HTML(datagrid))
	c.Ctx.View("backend/member_rolelist.html")
}

func (c *AdminController) RoleAdd() {
	if c.Ctx.Method() == "POST" {
		rolename := c.Ctx.FormValue("rolename")
		description := c.Ctx.FormValue("description")
		disabled, disabled_err := strconv.Atoi(c.Ctx.FormValue("disabled"))
		listorder, listorder_err := strconv.Atoi(c.Ctx.FormValue("listorder"))
		hasErr := helper.IsFalse(rolename, description) || helper.IsError(disabled_err, listorder_err)
		if hasErr {
			helper.Ajax("表单数据错误", 1, c.Ctx)
		} else {
			if !models.NewAdminModel(c.Orm).AddRole(rolename, description, int64(disabled), int64(listorder)) {
				helper.Ajax("添加角色失败", 1, c.Ctx)
			} else {
				helper.Ajax("添加角色成功", 0, c.Ctx)
			}
		}
		return
	}
	c.Ctx.View("backend/member_roleadd.html")
}
func (c *AdminController) RoleEdit() {
	id, err := c.Ctx.URLParamInt64("id")
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	role, err := models.NewAdminModel(c.Orm).GetRoleById(id)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	if c.Ctx.Method() == "POST" {
		rolename := c.Ctx.FormValue("rolename")
		description := c.Ctx.FormValue("description")
		disabled, disabled_err := strconv.Atoi(c.Ctx.FormValue("disabled"))
		listorder, listorder_err := strconv.Atoi(c.Ctx.FormValue("listorder"))
		hasErr := helper.IsFalse(rolename, description) || helper.IsError(disabled_err, listorder_err)
		if hasErr {
			helper.Ajax("表单数据错误", 1, c.Ctx)
		} else {
			role.Rolename = rolename
			role.Description = description
			role.Disabled = int64(disabled)
			role.Listorder = int64(listorder)
			if !models.NewAdminModel(c.Orm).UpdateRole(role) {
				helper.Ajax("修改角色失败", 1, c.Ctx)
			} else {
				clearMenuCache(c.Cache, c.Orm)
				helper.Ajax("修改角色成功", 0, c.Ctx)
			}
		}
		return
	}
	c.Ctx.ViewData("info", role)
	c.Ctx.View("backend/member_roleedit.html")
}

func (c *AdminController) RoleDelete() {
	roleid, _ := strconv.Atoi(c.Ctx.FormValue("id"))
	if roleid == 0 {
		helper.Ajax("没有选择任何角色", 1, c.Ctx)
		return
	}
	if roleid == 1 {
		helper.Ajax("不能删除ROLEID为1的角色", 1, c.Ctx)
		return
	}
	role, err := models.NewAdminModel(c.Orm).GetRoleById(int64(roleid))
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	if role.Rolename == "" {
		helper.Ajax("没有找到对应的角色信息", 1, c.Ctx)
		return
	}
	if models.NewAdminModel(c.Orm).HasAdminByRoleId(int64(roleid)) {
		helper.Ajax("该角色下有对应的管理员,无法删除", 1, c.Ctx)
		return
	}
	if models.NewAdminModel(c.Orm).DeleteRole(role) {
		clearMenuCache(c.Cache, c.Orm)
		helper.Ajax("删除角色成功", 0, c.Ctx)
	} else {
		helper.Ajax("删除角色失败", 1, c.Ctx)
	}
}

func (c *AdminController) RolePermission() {
	roleid, _ := c.Ctx.URLParamInt64("id")
	if roleid == 0 {
		helper.Ajax("没有选择任何角色", 1, c.Ctx)
		return
	}
	if c.Ctx.Method() == "POST" {
		//提交权限分配
		if c.Ctx.URLParam("dosubmit") == "1" {
			_, err := c.Orm.Where("roleid=?", roleid).Delete(&tables.IriscmsAdminRolePriv{})
			if err != nil {
				helper.Ajax("设置权限失败 "+err.Error(), 1, c.Ctx)
				return
			}
			menuIds := strings.Split(c.Ctx.FormValue("menuids"), ",")
			if len(menuIds) == 0 {
				helper.Ajax("没有选择任何权限", 1, c.Ctx)
				return
			}
			inserts := []tables.IriscmsAdminRolePriv{}
			for _, v := range menuIds {
				menuid, err := strconv.Atoi(v)
				if err != nil || menuid < 1 {
					continue
				}
				menu := tables.IriscmsMenu{Id: int64(menuid)}
				has, err := c.Orm.Get(&menu)
				if err != nil || !has {
					continue
				}
				inserts = append(inserts, tables.IriscmsAdminRolePriv{
					Roleid: roleid,
					A:      menu.A,
					C:      menu.C,
				})
			}
			if len(inserts) == 0 {
				helper.Ajax("没有选择任何权限", 1, c.Ctx)
				return
			}
			res, err := c.Orm.Insert(inserts)
			if err != nil || res == 0 {
				helper.Ajax("更新权限失败", 1, c.Ctx)
				return
			}
			clearMenuCache(c.Cache, c.Orm)
			helper.Ajax("更新权限成功", 0, c.Ctx)
			return
		}
		roleTree := models.NewMenuModel(c.Orm).GetRoleTree(0, roleid)
		c.Ctx.JSON(roleTree)
		return
	}
	c.Ctx.ViewData("roleid", roleid)
	c.Ctx.View("backend/role_permission.html")
}
