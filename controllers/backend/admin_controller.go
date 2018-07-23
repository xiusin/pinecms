package backend

import (
	"html/template"
	"iriscms/common/helper"
	"iriscms/models"
	"github.com/go-xorm/xorm"
	"strconv"
	"log"
	"strings"
	"iriscms/models/tables"
	"fmt"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type AdminController struct {
	Orm *xorm.Engine
	Ctx iris.Context	//存在则自动绑定
	Session *sessions.Session
}


func (c *AdminController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY","/admin/memberlist", "Memberlist")
	b.Handle("ANY","/admin/public-editpwd", "PublicEditpwd")
	b.Handle("ANY","/admin/public-editinfo", "PublicEditInfo")
	b.Handle("POST","/admin/public-checkEmail", "PublicCheckEmail")
	b.Handle("POST","/admin/public-checkName", "PubicCheckName")
	b.Handle("POST","/admin/check-password", "PublicCheckPassword")
	b.Handle("ANY","/admin/member-add", "MemberAdd")
	b.Handle("ANY","/admin/member-edit", "MemberEdit")
	b.Handle("ANY","/admin/member-delete", "MemberDelete")
	b.Handle("ANY","/admin/rolelist", "RoleList")
	b.Handle("ANY","/admin/role-add", "RoleAdd")
	b.Handle("ANY","/admin/role-edit", "RoleEdit")
	b.Handle("ANY","/admin/role-delete", "RoleDelete")
	b.Handle("ANY","/admin/role-permission", "RolePermission")
	b.Handle("ANY","/admin/check-rolename", "PublicCheckRoleName")

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

func (this *AdminController) PublicEditInfo() {
	aid, _ := this.Ctx.Values().GetInt64("adminid") //检测是否设置过session
	if this.Ctx.Method() == "POST" {
		info := tables.IriscmsAdmin{
			Userid:aid,
		}
		has, _ := this.Orm.Get(&info)        //读取用户资料
		if !has {
			helper.Ajax("用户资料已经不存在", 1, this.Ctx)
		} else {
			info.Realname = this.Ctx.PostValue("realname")
			info.Email = this.Ctx.PostValue("email")
			res, err := this.Orm.Id(aid).Update(info)
			if err != nil {
				helper.Ajax("修改资料失败" + err.Error() , 1, this.Ctx)
			} else {
				if res > 0 {
					helper.Ajax("修改资料成功", 0, this.Ctx)
				} else {
					helper.Ajax("修改资料失败" , 1, this.Ctx)
				}
			}
		}
		return
	}
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	currentPos := models.NewMenuModel(this.Orm).CurrentPos(menuid)
	info, err := models.NewAdminModel(this.Orm).GetUserInfo(aid)
	if err != nil {
		helper.Ajax(err.Error(), 1, this.Ctx)
		return
	}
	this.Ctx.ViewData("title", currentPos)
	this.Ctx.ViewData("info", info)
	if err := this.Ctx.View("backend/admin_editinfo.html"); err != nil {
		this.Ctx.WriteString(err.Error())
	}
}

func (this *AdminController) Memberlist() {
	act := this.Ctx.URLParam("grid")
	if act == "datagrid" {
		page, err := this.Ctx.URLParamInt("page")
		orderField := this.Ctx.URLParam("sort")
		if orderField == "" {
			orderField = "userid"
		}
		orderType := this.Ctx.URLParam("sort")
		if orderType == "" {
			orderType = "desc"
		}

		if err != nil {
			page = 1
		}
		data := models.NewAdminModel(this.Orm).GetList("1", page, 10, orderField, orderType)
		retData := []memberlist{}
		//将数据以map的方式返回吧.
		for _, v := range data {
			item := memberlist{
				Email:v.Email,
				Lastloginip : v.Lastloginip,
				Lastlogintime : helper.FormatTime(v.Lastlogintime),
				Realname : v.Realname,
				Rolename : "",
				Userid : v.Userid,
				Username : v.Username,
			}
			roleInfo, err := models.NewAdminModel(this.Orm).GetRoleById(int64(v.Roleid))
			if err != nil {
				roleInfo.Rolename = ""
			}
			item.Rolename = roleInfo.Rolename
			retData = append(retData, item)
		}

		this.Ctx.JSON(map[string]interface{}{
			"total": len(retData),
			"rows":  retData,
		})
		return
	}
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("admin_member_list_datagrid", "/b/admin/memberlist?grid=datagrid", helper.EasyuiOptions{
		"title":   models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar": "admin_memberlist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"用户名":    {"field": "Username", "width": "15", "sortable": "true", "index":"0"},
		"所属角色":   {"field": "Rolename", "width": "15", "sortable": "true", "index":"1"},
		"最后登录IP": {"field": "Lastloginip", "width": "15", "sortable": "true", "index":"2"},
		"最后登录时间": {"field": "Lastlogintime", "width": "15", "sortable": "true", "formatter": "adminMemberListTimeFormatter", "index":"3"},
		"E-mail": {"field": "Email", "width": "25", "sortable": "true", "index":"4"},
		"真实姓名":   {"field": "Realname", "width": "15", "sortable": "true", "index":"5"},
		"管理操作":   {"field": "Userid", "width": "15", "sortable": "true", "formatter": "adminMemberListOperateFormatter", "index":"6"},
	})
	this.Ctx.ViewData("table", template.HTML(table))
	this.Ctx.View("backend/admin_memberlist.html")
}

func (this *AdminController) PublicEditpwd() {
	aid, _ := this.Ctx.Values().GetInt64("adminid")
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	info := tables.IriscmsAdmin{Userid: int64(aid)}
	has, _ := this.Orm.Get(&info)
	if !has {
		this.Ctx.WriteString("没有找到")
		return
	}
	if this.Ctx.Method() == "POST" {
		if info.Password != helper.Password(this.Ctx.PostValue("old_password"), info.Encrypt) {
			helper.Ajax("原密码错误", 1, this.Ctx)
			return
		}
		info.Password = helper.Password(this.Ctx.PostValue("new_password"), info.Encrypt)
		res, _ := this.Orm.Id(aid).Update(info)
		if res > 0 {
			helper.Ajax("修改资料成功", 0, this.Ctx)
		} else {
			helper.Ajax("修改资料失败", 1, this.Ctx)
		}
		return
	}
	this.Ctx.ViewData("currentpos", models.NewMenuModel(this.Orm).CurrentPos(menuid))
	this.Ctx.ViewData("admin", info)
	this.Ctx.View("backend/admin_editpwd.html")
}

func (this *AdminController) PublicCheckEmail() {
	info := &tables.IriscmsAdmin{Username:this.Ctx.FormValue("name")}
	has, _ := this.Orm.Get(info)
	if !has {
		helper.Ajax("没有相同的用户名", 0, this.Ctx)
	} else {
		helper.Ajax("已经有相同的用户名,请换一个再试", 1, this.Ctx)
	}
}
func (this *AdminController) PublicCheckPassword() {
	aid, _ := this.Ctx.Values().Get("adminid").(int64)
	password := this.Ctx.FormValue("password")
	admin, err := models.NewAdminModel(this.Orm).GetUserInfo(aid)
	if err != nil {
		helper.Ajax("无法查找到相关信息", 1, this.Ctx)
		return
	}
	if admin.Password != helper.Password(password, admin.Encrypt) {
		helper.Ajax("旧密码错误", 1, this.Ctx)
		return
	}
	helper.Ajax("验证密码成功", 0, this.Ctx)
}
func (this *AdminController) PubicCheckName() {
	info := &tables.IriscmsAdmin{Username:this.Ctx.FormValue("name")}
	uid, _ := this.Ctx.URLParamInt64("id")
	has, _ := this.Orm.Get(info)
	fmt.Println(uid, info.Userid)
	if !has || info.Userid == uid {
		helper.Ajax("没有相同的用户名", 0, this.Ctx)
	} else {
		helper.Ajax("已经有相同的用户名,请换一个再试", 1, this.Ctx)
	}
}

func (this *AdminController) PublicCheckRoleName() {
	rolename := this.Ctx.FormValue("rolename")
	if !helper.IsAjax(this.Ctx) || rolename == "" {
		helper.Ajax("参数错误 ," + rolename, 1, this.Ctx)
		return
	}
	defaultName := this.Ctx.FormValue("default")
	if defaultName != "" && rolename == defaultName {
		helper.Ajax("角色已存在", 1, this.Ctx)
		return
	}
	if models.NewAdminModel(this.Orm).CheckRoleName(rolename) {
		helper.Ajax("角色已存在", 1, this.Ctx)
		return
	}
	helper.Ajax("通过", 0, this.Ctx)
}
func (this *AdminController) MemberAdd() {
	if this.Ctx.FormValue("act") != "" {
		if this.Ctx.FormValue("pwdconfirm") != this.Ctx.FormValue("password") || this.Ctx.FormValue("password") == "" {
			helper.Ajax("两次密码不一致", 1, this.Ctx)
			return
		}
		if this.Ctx.FormValue("roleid") == "" {
			helper.Ajax("请选择角色", 1, this.Ctx)
			return
		}
		roleid, err := strconv.Atoi(this.Ctx.FormValue("roleid"))
		if err != nil {
			helper.Ajax("角色信息错误", 1, this.Ctx)
			return
		}
		str := string(helper.Krand(6, 3))
		newAdmin := &tables.IriscmsAdmin{
			Username:this.Ctx.FormValue("username"),
			Password:helper.Password(this.Ctx.FormValue("password"), str),
			Email:this.Ctx.FormValue("email"),
			Encrypt:str,
			Realname:this.Ctx.FormValue("realname"),
			Roleid:int64(roleid),
		}
		id, err := this.Orm.Insert(newAdmin)
		if id > 0 {
			helper.Ajax("添加管理员成功", 0, this.Ctx)
			return
		}
		helper.Ajax("添加管理员失败", 1, this.Ctx)
		return
	}
	roles := models.NewAdminModel(this.Orm).GetRoleList("1", 1, 1000)
	this.Ctx.ViewData("Roles", roles)
	this.Ctx.View("backend/member_add.html")
}
func (this *AdminController) MemberEdit() {
	adminid, err := this.Ctx.URLParamInt64("id")
	if err != nil {
		this.Ctx.WriteString("参数错误 : " + err.Error())
		return
	}
	info, err := models.NewAdminModel(this.Orm).GetUserInfo(adminid)
	if err != nil {
		this.Ctx.WriteString("没有该管理员信息")
		return
	}
	if this.Ctx.Method() == "POST" {
		if this.Ctx.FormValue("password") != "" {
			if this.Ctx.FormValue("pwdconfirm") != this.Ctx.FormValue("password") {
				helper.Ajax("两次密码不一致", 1, this.Ctx)
				return
			}
		}
		if this.Ctx.FormValue("roleid") == "" {
			helper.Ajax("请选择角色", 1, this.Ctx)
			return
		}
		roleid, err := strconv.Atoi(this.Ctx.FormValue("roleid"))
		if err != nil {
			helper.Ajax("角色信息错误", 1, this.Ctx)
			return
		}
		info.Username = this.Ctx.FormValue("username")
		info.Email = this.Ctx.FormValue("email")
		info.Realname = this.Ctx.FormValue("realname")
		info.Roleid = int64(roleid)
		if this.Ctx.FormValue("password") != "" {
			info.Password = helper.Password(this.Ctx.PostValue("password"), info.Encrypt)
		}
		res , err := this.Orm.Where("userid = ?",info.Userid).Update(info)
		if err != nil {
			helper.Ajax(err.Error(), 0, this.Ctx)
			return
		}
		if res > 0 {
			helper.Ajax("修改管理员成功", 0, this.Ctx)
			return
		}
		helper.Ajax("修改管理员失败", 1, this.Ctx)
		return
	}

	roles := models.NewAdminModel(this.Orm).GetRoleList("1", 1, 1000)
	this.Ctx.ViewData("Roles", roles)
	this.Ctx.ViewData("Info", info)
	this.Ctx.View("backend/member_edit.html")
}
func (this *AdminController) MemberDelete() {
	id, err := strconv.Atoi(this.Ctx.FormValue("id"))
	if err != nil || helper.IsFalse(id) || id == 1 {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	deleteAdmin := &tables.IriscmsAdmin{Userid:int64(id)}
	res, err := this.Orm.Delete(deleteAdmin)
	if err != nil || helper.IsFalse(res) {
		helper.Ajax("删除失败", 1, this.Ctx)
		return
	}
	helper.Ajax("删除成功", 0, this.Ctx)
}

func (this *AdminController) RoleList() {
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	if this.Ctx.URLParam("grid") == "datagrid" {
		page, err := this.Ctx.URLParamInt("page")
		orderField := this.Ctx.URLParam("sort")
		if orderField == "" {
			orderField = "id"
		}
		orderType := this.Ctx.URLParam("sort")
		if orderType == "" {
			orderType = "desc"
		}

		if err != nil {
			page = 1
		}

		data := models.NewAdminModel(this.Orm).GetRoleList("1", page, 1000)
		this.Ctx.JSON(map[string]interface{}{
			"total": len(data),
			"rows": data,
		})
		return
	}

	datagrid := helper.Datagrid("admin_rolelist_list_datagrid", "/b/admin/rolelist?grid=datagrid", helper.EasyuiOptions{
		"title":models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar": "admin_rolelist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"角色名称": {"field": "Rolename", "width": "15", "index":"0"},
		"角色描述": {"field": "Description", "width": "15", "index":"1"},
		"管理操作":   {"field": "Roleid", "width": "15", "formatter": "adminRoleListOperateFormatter", "index":"2"},
	})
	this.Ctx.ViewData("datagrid", template.HTML(datagrid))
	this.Ctx.View("backend/member_rolelist.html")
}

func (this *AdminController) RoleAdd() {
	if this.Ctx.Method() == "POST" {
		rolename := this.Ctx.FormValue("rolename")
		description := this.Ctx.FormValue("description")
		disabled, disabled_err := strconv.Atoi(this.Ctx.FormValue("disabled"))
		listorder, listorder_err := strconv.Atoi(this.Ctx.FormValue("listorder"))
		hasErr := helper.IsFalse(rolename, description) || helper.IsError(disabled_err, listorder_err)
		if hasErr {
			helper.Ajax("表单数据错误", 1, this.Ctx)
		} else {
			if !models.NewAdminModel(this.Orm).AddRole(rolename, description, int64(disabled), int64(listorder)) {
				helper.Ajax("添加角色失败", 1, this.Ctx)
			} else {
				helper.Ajax("添加角色成功", 0, this.Ctx)
			}
		}
		return
	}
	this.Ctx.View("backend/member_roleadd.html")
}
func (this *AdminController) RoleEdit() {
	id, err := this.Ctx.URLParamInt64("id")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	role, err := models.NewAdminModel(this.Orm).GetRoleById(id)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	if this.Ctx.Method() == "POST" {
		rolename := this.Ctx.FormValue("rolename")
		description := this.Ctx.FormValue("description")
		disabled, disabled_err := strconv.Atoi(this.Ctx.FormValue("disabled"))
		listorder, listorder_err := strconv.Atoi(this.Ctx.FormValue("listorder"))
		hasErr := helper.IsFalse(rolename, description) || helper.IsError(disabled_err, listorder_err)
		if hasErr {
			helper.Ajax("表单数据错误", 1, this.Ctx)
		} else {
			role.Rolename = rolename
			role.Description = description
			role.Disabled = int64(disabled)
			role.Listorder = int64(listorder)
			log.Println(role)
			if !models.NewAdminModel(this.Orm).UpdateRole(role) {
				helper.Ajax("修改角色失败", 1, this.Ctx)
			} else {
				helper.Ajax("修改角色成功", 0, this.Ctx)
			}
		}
		return
	}
	this.Ctx.ViewData("info", role)
	this.Ctx.View("backend/member_roleedit.html")
}

func (this *AdminController) RoleDelete() {
	roleid, _ := strconv.Atoi(this.Ctx.FormValue("id"))
	if roleid == 0 {
		helper.Ajax("没有选择任何角色", 1, this.Ctx)
		return
	}
	if roleid == 1 {
		helper.Ajax("不能删除ROLEID为1的角色", 1, this.Ctx)
		return
	}
	role, err := models.NewAdminModel(this.Orm).GetRoleById(int64(roleid))
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	if role.Rolename == "" {
		helper.Ajax("没有找到对应的角色信息", 1, this.Ctx)
		return
	}
	if models.NewAdminModel(this.Orm).HasAdminByRoleId(int64(roleid)) {
		helper.Ajax("该角色下有对应的管理员,无法删除", 1, this.Ctx)
		return
	}
	if models.NewAdminModel(this.Orm).DeleteRole(role) {
		helper.Ajax("删除角色成功", 0, this.Ctx)
	} else {
		helper.Ajax("删除角色失败", 1, this.Ctx)
	}
}

func (this *AdminController) RolePermission() {
	roleid, _ := this.Ctx.URLParamInt64("id")
	if roleid == 0 {
		helper.Ajax("没有选择任何角色", 1, this.Ctx)
		return
	}
	if this.Ctx.Method() == "POST" {
		//提交权限分配
		if this.Ctx.URLParam("dosubmit") == "1" {
			_, err := this.Orm.Where("roleid=?", roleid).Delete(&tables.IriscmsAdminRolePriv{})
			if err != nil {
				helper.Ajax("设置权限失败 " + err.Error(), 1, this.Ctx)
				return
			}
			menuIds := strings.Split(this.Ctx.FormValue("menuids"), ",")
			if len(menuIds) == 0 {
				helper.Ajax("没有选择任何权限", 1, this.Ctx)
				return
			}
			inserts := []tables.IriscmsAdminRolePriv{}
			for _, v := range menuIds {
				menuid, err := strconv.Atoi(v)
				if err != nil || menuid < 1 {
					continue
				}
				menu := tables.IriscmsMenu{Id:int64(menuid)}
				has, err := this.Orm.Get(&menu)
				if err != nil || !has {
					continue
				}
				inserts = append(inserts, tables.IriscmsAdminRolePriv{
					Roleid:roleid,
					A:menu.A,
					C:menu.C,
				})
			}
			if len(inserts) == 0 {
				helper.Ajax("没有选择任何权限", 1, this.Ctx)
				return
			}
			res, err := this.Orm.Insert(inserts)
			if err != nil || res == 0 {
				helper.Ajax("更新权限失败", 1, this.Ctx)
				return
			}
			helper.Ajax("更新权限成功", 0, this.Ctx)
			return
		}
		roleTree := models.NewMenuModel(this.Orm).GetRoleTree(0, roleid)
		this.Ctx.JSON(roleTree)
		return
	}
	this.Ctx.ViewData("roleid", roleid)
	this.Ctx.View("backend/role_permission.html")
}