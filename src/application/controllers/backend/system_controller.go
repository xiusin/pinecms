package backend

import (
	"html/template"
	"strconv"
	"strings"

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

func (c *SystemController) MenuList() {
	if c.Ctx.URLParam("grid") == "treegrid" {
		c.Ctx.JSON(models.NewMenuModel(c.Orm).GetTree(models.NewMenuModel(c.Orm).GetAll(), 0))
		return
	}
	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Treegrid("system_menu_list", "/b/system/menulist?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar":   "system_menulist_treegrid_toolbar",
		"idField":   "operateid",
		"treeField": "name",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "10", "align": "center", "formatter": "systemMenuOrderFormatter", "index": "0"},
		"菜单名称": {"field": "name", "width": "130", "index": "1"},
		"管理操作": {"field": "operateid", "width": "25", "align": "center", "formatter": "systemMenuOperateFormatter", "index": "2"},
	})
	c.Ctx.ViewData("treegrid", template.HTML(table))
	c.Ctx.View("backend/system_menulist.html")
}

func (c *SystemController) MenuAdd() {
	if c.Ctx.Method() == "POST" {
		parentid := c.Ctx.PostValueInt64Default("parentid", 0)
		name := c.Ctx.PostValueDefault("name", "")
		ctrl := c.Ctx.PostValueDefault("c", "")
		a := c.Ctx.PostValueDefault("a", "")
		data := c.Ctx.PostValueDefault("data", "")
		display := c.Ctx.PostValueInt64Default("display", 1)

		menu := &tables.IriscmsMenu{
			Parentid: parentid,
			Name:     name,
			C:        ctrl,
			A:        a,
			Data:     data,
			Display:  display,
		}
		newid, _ := c.Orm.InsertOne(menu)
		if newid > 0 {
			clearMenuCache(c.Cache, c.Orm)
			helper.Ajax("新增菜单成功", 0, c.Ctx)
		} else {
			helper.Ajax("新增菜单失败", 1, c.Ctx)
		}
		return
	}
	c.Ctx.View("backend/system_menuadd.html")
}

func (c *SystemController) MenuDelete() {
	id := c.Ctx.PostValueInt64Default("id", 0)
	if id < 1 {
		helper.Ajax("参数失败", 1, c.Ctx)
		return
	}
	// 查找是否有下级菜单
	exists, err := c.Orm.Where("parentid = ?", id).Count(&tables.IriscmsMenu{})
	if err != nil {
		helper.Ajax("删除菜单失败,异常错误", 1, c.Ctx)
		return
	}
	if exists > 0 {
		helper.Ajax("删除菜单失败,有下级菜单", 1, c.Ctx)
		return
	}
	aff, _ := c.Orm.Id(id).Delete(&tables.IriscmsMenu{})
	if aff > 0 {
		clearMenuCache(c.Cache, c.Orm)
		helper.Ajax("删除菜单成功", 0, c.Ctx)
	} else {
		helper.Ajax("删除菜单失败", 1, c.Ctx)
	}
}

func (c *SystemController) MenuOrder() {
	posts := c.Ctx.FormValues()
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
		affected, _ := c.Orm.Id(id).Update(menu)
		if affected > 0 {
			flag++
		}
	}
	if flag > 0 {
		clearMenuCache(c.Cache, c.Orm)
		helper.Ajax("排序更新成功", 0, c.Ctx)
	} else {
		helper.Ajax("排序规则没有发生任何改变", 1, c.Ctx)
	}
}

func (c *SystemController) MenuEdit() {
	id, _ := c.Ctx.URLParamInt64("id")
	if id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	menu, has := models.NewMenuModel(c.Orm).GetInfo(id)
	if !has {
		helper.Ajax("没有此菜单", 1, c.Ctx)
		return
	}
	if c.Ctx.Method() == "POST" {
		parentid := c.Ctx.PostValueInt64Default("parentid", 0)
		name := c.Ctx.PostValueDefault("name", "")
		ctrl := c.Ctx.PostValueDefault("c", "")
		a := c.Ctx.PostValueDefault("a", "")
		data := c.Ctx.PostValueDefault("data", "")
		display := c.Ctx.PostValueInt64Default("display", 1)

		menu := &tables.IriscmsMenu{
			Parentid: parentid,
			Name:     name,
			C:        ctrl,
			A:        a,
			Data:     data,
			Display:  display,
		}
		newid, _ := c.Orm.Id(id).Update(menu)
		if newid > 0 {
			clearMenuCache(c.Cache, c.Orm)
			helper.Ajax("修改菜单成功", 0, c.Ctx)
		} else {
			helper.Ajax("修改菜单失败", 1, c.Ctx)
		}
		return
	}

	c.Ctx.ViewData("data", menu)
	c.Ctx.View("backend/system_menuedit.html")
}

func (c *SystemController) MenuSelectTree() {
	var tree []map[string]interface{}
	tree = append(tree, map[string]interface{}{
		"id":       0,
		"text":     "作为一级菜单",
		"children": models.NewMenuModel(c.Orm).GetSelectTree(models.NewMenuModel(c.Orm).GetAll(), 0),
	})
	c.Ctx.JSON(tree)
}

func (c *SystemController) MenuCheck() {
	name := c.Ctx.FormValue("name")
	if name == "" {
		helper.Ajax("用户名为空", 1, c.Ctx)
		return
	}
	if models.NewMenuModel(c.Orm).CheckName(name) {
		helper.Ajax("用户名已存在", 1, c.Ctx)
		return
	}

	helper.Ajax("正常", 0, c.Ctx)
}

func (c *SystemController) LogList() {
	page, _ := c.Ctx.URLParamInt64("page")
	rows, _ := c.Ctx.URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewLogModel(c.Orm).GetList(page, rows)
		c.Ctx.JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("system_menu_logList", "/b/system/loglist", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "system_loglist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"用户名": {"field": "Username", "width": "20", "index": "0"},
		"模块":  {"field": "Controller", "width": "15", "index": "1"},
		"方法":  {"field": "Action", "width": "15", "index": "2"},
		"URL": {"field": "Querystring", "width": "100", "formatter": "systemLogViewFormatter", "index": "3"},
		"时间":  {"field": "Time", "width": "30", "index": "4"},
		"IP":  {"field": "Ip", "width": "25", "index": "5"},
	})
	c.Ctx.ViewData("dataGrid", template.HTML(table))
	c.Ctx.View("backend/system_loglist.html")

}

func (c *SystemController) LogDelete() {
	if models.NewLogModel(c.Orm).DeleteAll() {
		helper.Ajax("清空日志成功", 0, c.Ctx)
	} else {
		helper.Ajax("清空日志失败", 1, c.Ctx)
	}
}
