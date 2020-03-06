package backend

import (
	"github.com/xiusin/pine"
	"html/template"

	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type MemberController struct {
	pine.Controller
}

func (c *MemberController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY( "/user/list", "List")
	b.ANY( "/user/info", "Info")
	b.POST( "/user/edit", "Edit")
	b.ANY( "/wechat/userlist", "WechatMemberList")
	b.ANY( "/wechat/userinfo", "WechatMemberInfo")
}

func (c *MemberController) List() {
	page, _ := c.Ctx().URLParamInt64("page")
	rows, _ := c.Ctx().URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewMemberModel().GetList(page, rows)
		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("member_list_datagrid", "/b/user/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel().CurrentPos(menuid),
		"toolbar": "member_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"昵称": {"field": "nickname", "width": "30", "index": "0"},
		"账户": {"field": "account", "width": "30", "index": "1"},
		"邮箱": {"field": "email", "width": "60", "index": "2"},
		"积分": {"field": "integral", "width": "25", "index": "3"},
		"时间": {"field": "created_at", "width": "25", "index": "4"},
		"启用": {"field": "enabled", "width": "25", "index": "5", "formatter": "enabledFormatter"},
		"验证": {"field": "verify_token", "width": "25", "index": "6", "formatter": "verifyTokenFormatter"},
		"操作": {"field": "id", "width": "25", "index": "7", "formatter": "optFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/member_list.html")
}

func (c *MemberController) Info() {
	id, _ := c.Ctx().URLParamInt64("id")
	if id < 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	member := models.NewMemberModel().GetInfo(id)
	if member.Id != id {
		helper.Ajax("用户信息获取失败", 1, c.Ctx())
		return
	}
	helper.Ajax(member, 0, c.Ctx())
	return
}

func (c *MemberController) Edit() {
	var d tables.IriscmsMember
	if err := c.Ctx().BindForm(&d); err != nil || d.Id <= 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	m := models.NewMemberModel()
	if m.Edit(d.Id, &d) {
		helper.Ajax("更新用户信息成功", 0, c.Ctx())
	} else {
		helper.Ajax("更新用户信息失败", 1, c.Ctx())
	}
}

//微信用户列表(通过关注公众号获取密码的用户)
func (c *MemberController) WechatMemberList() {
	page, _ := c.Ctx().URLParamInt64("page")
	rows, _ := c.Ctx().URLParamInt64("rows")
	if page > 0 {
		list, total := models.NewWechatMemberModel().GetList(page, rows)
		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("wechat_member_list_datagrid", "/b/wechat/userlist", helper.EasyuiOptions{
		"title":   models.NewMenuModel().CurrentPos(menuid),
	}, helper.EasyuiGridfields{
		"昵称": {"field": "nickname", "width": "30", "index": "0"},
		"账户": {"field": "account", "width": "30", "index": "1"},
		"邮箱": {"field": "email", "width": "60", "index": "2"},
		"积分": {"field": "integral", "width": "25", "index": "3"},
		"时间": {"field": "created_at", "width": "25", "index": "4"},
		"启用": {"field": "enabled", "width": "25", "index": "5", "formatter": "enabledFormatter"},
		"验证": {"field": "verify_token", "width": "25", "index": "6", "formatter": "verifyTokenFormatter"},
		"操作": {"field": "id", "width": "25", "index": "7", "formatter": "optFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/wechat_member_list.html")
}

//微信用户列表(通过关注公众号获取密码的用户)
func (c *MemberController) WechatMemberInfo() {
	id, _ := c.Ctx().URLParamInt64("id")
	if id < 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	member := models.NewWechatMemberModel().GetInfo(id)
	if member.Id != id {
		helper.Ajax("用户信息获取失败", 1, c.Ctx())
		return
	}
	helper.Ajax(member, 0, c.Ctx())
	return
}
