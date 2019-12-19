package backend

import (
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type SystemController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Cache   cache.ICache
	Session *sessions.Session
}

func (c *SystemController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/system/menulist", "MenuList")
	b.Handle("ANY", "/system/menu-edit", "MenuEdit")
	b.Handle("ANY", "/system/menu-add", "MenuAdd")
	b.Handle("POST", "/system/menu-delete", "MenuDelete")
	b.Handle("POST", "/system/menu-order", "MenuOrder")
	b.Handle("ANY", "/system/menu-tree", "MenuSelectTree")
	b.Handle("POST", "/system/menu-check", "MenuCheck")
	b.Handle("ANY", "/system/loglist", "LogList")
	b.Handle("ANY", "/system/log-delete", "LogDelete")
}

func (this *SystemController) MenuList() {
	if this.Ctx.URLParam("grid") == "treegrid" {
		this.Ctx.JSON(models.NewMenuModel(this.Orm).GetTree(models.NewMenuModel(this.Orm).GetAll(), 0))
		return
	}
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	table := helper.Treegrid("system_menu_list", "/b/system/menulist?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar":   "system_menulist_treegrid_toolbar",
		"idField":   "operateid",
		"treeField": "name",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "10", "align": "center", "formatter": "systemMenuOrderFormatter", "index": "0"},
		"菜单名称": {"field": "name", "width": "130", "index": "1"},
		"管理操作": {"field": "operateid", "width": "25", "align": "center", "formatter": "systemMenuOperateFormatter", "index": "2"},
	})
	this.Ctx.ViewData("treegrid", template.HTML(table))
	this.Ctx.View("backend/system_menulist.html")
}

func (this *SystemController) MenuAdd() {
	if this.Ctx.Method() == "POST" {
		parentid := this.Ctx.PostValueInt64Default("parentid", 0)
		name := this.Ctx.PostValueDefault("name", "")
		c := this.Ctx.PostValueDefault("c", "")
		a := this.Ctx.PostValueDefault("a", "")
		data := this.Ctx.PostValueDefault("data", "")
		display := this.Ctx.PostValueInt64Default("display", 1)

		menu := &tables.IriscmsMenu{
			Parentid: parentid,
			Name:     name,
			C:        c,
			A:        a,
			Data:     data,
			Display:  display,
		}
		newid, _ := this.Orm.InsertOne(menu)
		if newid > 0 {
			clearMenuCache(this.Cache, this.Orm)
			helper.Ajax("新增菜单成功", 0, this.Ctx)
		} else {
			helper.Ajax("新增菜单失败", 1, this.Ctx)
		}
		return
	}
	this.Ctx.View("backend/system_menuadd.html")
}

func (this *SystemController) MenuDelete() {
	id := this.Ctx.PostValueInt64Default("id", 0)
	if id < 1 {
		helper.Ajax("参数失败", 1, this.Ctx)
		return
	}
	// 查找是否有下级菜单
	exists, err := this.Orm.Where("parentid = ?", id).Count(&tables.IriscmsMenu{})
	if err != nil {
		helper.Ajax("删除菜单失败,异常错误", 1, this.Ctx)
		return
	}
	if exists > 0 {
		helper.Ajax("删除菜单失败,有下级菜单", 1, this.Ctx)
		return
	}
	aff, _ := this.Orm.Id(id).Delete(&tables.IriscmsMenu{})
	if aff > 0 {
		clearMenuCache(this.Cache, this.Orm)
		helper.Ajax("删除菜单成功", 0, this.Ctx)
	} else {
		helper.Ajax("删除菜单失败", 1, this.Ctx)
	}
}

func (this *SystemController) MenuOrder() {
	posts := this.Ctx.FormValues()
	data := map[int64]int64{}
	for k, v := range posts {
		k = strings.Replace(k, "order[", "", 1)
		k = strings.Replace(k, "]", "", 1)
		s, e := strconv.Atoi(k)
		if e != nil {
			continue
		}
		sort, e := strconv.Atoi(v[0])
		if e != nil {
			continue
		}
		data[int64(s)] = int64(sort)
	}
	var flag int64
	for id, val := range data {
		menu := new(tables.IriscmsMenu)
		menu.Listorder = val
		affected, _ := this.Orm.Id(id).Update(menu)
		if affected > 0 {
			flag++
		}
	}
	if flag > 0 {
		clearMenuCache(this.Cache, this.Orm)
		helper.Ajax("排序更新成功", 0, this.Ctx)
	} else {
		helper.Ajax("排序规则没有发生任何改变", 1, this.Ctx)
	}
}

func (this *SystemController) MenuEdit() {
	id, _ := this.Ctx.URLParamInt64("id")
	if id == 0 {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	menu, has := models.NewMenuModel(this.Orm).GetInfo(id)
	if !has {
		helper.Ajax("没有此菜单", 1, this.Ctx)
		return
	}
	if this.Ctx.Method() == "POST" {
		parentid := this.Ctx.PostValueInt64Default("parentid", 0)
		name := this.Ctx.PostValueDefault("name", "")
		c := this.Ctx.PostValueDefault("c", "")
		a := this.Ctx.PostValueDefault("a", "")
		data := this.Ctx.PostValueDefault("data", "")
		display := this.Ctx.PostValueInt64Default("display", 1)

		menu := &tables.IriscmsMenu{
			Parentid: parentid,
			Name:     name,
			C:        c,
			A:        a,
			Data:     data,
			Display:  display,
		}
		newid, _ := this.Orm.Id(id).Update(menu)
		if newid > 0 {
			clearMenuCache(this.Cache, this.Orm)
			helper.Ajax("修改菜单成功", 0, this.Ctx)
		} else {
			helper.Ajax("修改菜单失败", 1, this.Ctx)
		}
		return
	}

	this.Ctx.ViewData("data", menu)
	this.Ctx.View("backend/system_menuedit.html")
}

func (this *SystemController) MenuSelectTree() {
	var tree []map[string]interface{}
	tree = append(tree, map[string]interface{}{
		"id":       0,
		"text":     "作为一级菜单",
		"children": models.NewMenuModel(this.Orm).GetSelectTree(models.NewMenuModel(this.Orm).GetAll(), 0),
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

func (this *SystemController) LogList() {
	page, _ := this.Ctx.URLParamInt64("page")
	rows, _ := this.Ctx.URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewLogModel(this.Orm).GetList(page, rows)
		this.Ctx.JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := this.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("system_menu_logList", "/b/system/loglist", helper.EasyuiOptions{
		"title":   models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar": "system_loglist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"用户名": {"field": "Username", "width": "20", "index": "0"},
		"模块":  {"field": "Controller", "width": "15", "index": "1"},
		"方法":  {"field": "Action", "width": "15", "index": "2"},
		"URL": {"field": "Querystring", "width": "100", "formatter": "systemLogViewFormatter", "index": "3"},
		"时间":  {"field": "Time", "width": "30", "index": "4"},
		"IP":  {"field": "Ip", "width": "25", "index": "5"},
	})
	this.Ctx.ViewData("dataGrid", template.HTML(table))
	this.Ctx.View("backend/system_loglist.html")

}

func (this *SystemController) LogDelete() {
	//删除日志
	//删除一个月前的日志
	date := helper.FormatTime(time.Now().AddDate(0, -1, 0).Unix())
	if models.NewLogModel(this.Orm).DeleteBeforeByDate(date) {
		helper.Ajax("删除"+date+"前的日志成功", 0, this.Ctx)
		return
	}
	helper.Ajax("删除"+date+"前的日志失败", 1, this.Ctx)
}
