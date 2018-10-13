package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
	"iriscms/application/models"
	"iriscms/common/helper"
)

type MemberController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (c *MemberController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/user/list", "List")
}

func (c *MemberController) List() {
	page, _ := c.Ctx.URLParamInt64("page")
	rows, _ := c.Ctx.URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewMemberModel(c.Orm).GetList(page, rows)
		c.Ctx.JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("system_menu_logList", "/b/user/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "system_loglist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"昵称": {"field": "nickname", "width": "30", "index": "0"},
		"账户":  {"field": "account", "width": "30", "index": "1"},
		"邮箱":  {"field": "email","width": "60",  "index": "2"},
		"积分": {"field": "integral", "width": "25",  "index": "3"},
		"时间":  {"field": "created_at", "width": "25", "index": "4"},
		"启用":  {"field": "enabled", "width": "25", "index": "5","formatter": "enabledFormatter"},
		"验证":  {"field": "verify_token", "width": "25", "index": "6","formatter": "verifyTokenFormatter"},
		"操作":  {"field": "id", "width": "25", "index": "7","formatter": "optFormatter"},
	})
	c.Ctx.ViewData("dataGrid", template.HTML(table))
	c.Ctx.View("backend/member_list.html")
}
