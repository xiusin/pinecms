package backend

import (
	"github.com/go-xorm/xorm"
	"iriscms/models"
	"iriscms/controllers/backend/helper"
	"html/template"
	"time"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type SystemController struct {
	Ctx iris.Context
	Orm *xorm.Engine
	Session *sessions.Session
}


func (c *SystemController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY","/system/menulist", "MenuList")
	b.Handle("ANY","/system/menu-edit", "MenuEdit")
	b.Handle("ANY","/system/menu-add", "MenuAdd")
	b.Handle("ANY","/system/menu-tree", "MenuSelectTree")
	b.Handle("POST","/system/menu-check", "MenuCheck")
	b.Handle("ANY","/system/loglist", "LogList")
	b.Handle("ANY","/system/log-delete", "LogDelete")
}


func (this *SystemController) MenuList() {
	if this.Ctx.URLParam("grid") == "treegrid" {
		this.Ctx.JSON(models.NewMenuModel(this.Orm).GetTree(models.NewMenuModel(this.Orm).GetAll(),0))
		return
	}
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	table := helper.Treegrid("system_menu_list", "/b/system/menulist?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar":   "system_menulist_treegrid_toolbar",
		"idField":   "operateid",
		"treeField": "name",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "10", "align": "center", "formatter": "systemMenuOrderFormatter", "index":"0"},
		"菜单名称": {"field": "name", "width" : "130", "index":"1"},
		"管理操作": {"field": "operateid", "width": "25", "align": "center", "formatter": "systemMenuOperateFormatter", "index":"2"},
	})
	this.Ctx.ViewData("treegrid", template.HTML(table))
	this.Ctx.View("backend/system_menulist.html")
}

func (this *SystemController) MenuAdd() {
	this.Ctx.View("backend/system_menuadd.html")
}

func (this *SystemController) MenuEdit() {

	id, _ := this.Ctx.URLParamInt64("id")
	if id == 0 {
		this.Ctx.WriteString("参数错误")
		return
	}
	menu, has := models.NewMenuModel(this.Orm).GetInfo(id)
	if !has {
		this.Ctx.WriteString("没有菜单")
	}
	this.Ctx.ViewData("data", menu)
	this.Ctx.View("backend/system_menuedit.html");
}

func (this *SystemController) MenuSelectTree() {
	tree := []map[string]interface{}{}
	tree = append(tree, map[string]interface{}{
		"id":0,
		"text":"作为一级菜单",
		"children":models.NewMenuModel(this.Orm).GetSelectTree(models.NewMenuModel(this.Orm).GetAll(),0),
	})
	this.Ctx.JSON(tree)
}

func (this *SystemController) MenuCheck() {
	name := this.Ctx.FormValue("name")
	if name == "" {
		helper.Ajax("用户名为空", 1, this.Ctx)
		return
	}
	if models.NewMenuModel(this.Orm).CheckName(name) {
		helper.Ajax("用户名已存在", 1, this.Ctx)
		return
	}

	helper.Ajax("正常", 0, this.Ctx)
}

func (this *SystemController)LogList() {
	page, _ := this.Ctx.URLParamInt64("page")
	rows, _ := this.Ctx.URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewLogModel(this.Orm).GetList(page, rows)
		this.Ctx.JSON(map[string]interface{}{"rows":list, "total":total})
		return
	}

	menuid, _ := this.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("system_menu_logList", "/b/system/loglist", helper.EasyuiOptions{
		"title":     models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar":   "system_loglist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"用户名":   {"field": "Username", "width": "20", "index":"0"},
		"模块": {"field": "Controller", "width" : "15", "index":"1"},
		"方法": {"field": "Action", "width" : "15", "index":"2"},
		"URL": {"field": "Querystring", "width" : "100", "formatter":"systemLogViewFormatter", "index":"3"},
		"时间": {"field": "Time", "width" : "30", "index":"4"},
		"IP": {"field": "Ip", "width" : "25", "index":"5"},
	})
	this.Ctx.ViewData("dataGrid", template.HTML(table))
	this.Ctx.View("backend/system_loglist.html")

}

func (this *SystemController)LogDelete() {
	//删除日志
	//删除一个月前的日志
	date := helper.FormatTime(time.Now().AddDate(0, -1, 0).Unix())
	if models.NewLogModel(this.Orm).DeleteBeforeByDate(date) {
		helper.Ajax("删除日志成功", 0, this.Ctx)
		return
	}
	helper.Ajax("删除日志失败", 1, this.Ctx)
}