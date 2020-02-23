package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"html/template"
	"strconv"
	"strings"

	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type SystemController struct {
	pine.Controller
}

func (c *SystemController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY( "/system/menulist", "MenuList")
	b.ANY( "/system/menu-edit", "MenuEdit")
	b.ANY( "/system/menu-add", "MenuAdd")
	b.POST( "/system/menu-delete", "MenuDelete")
	b.POST( "/system/menu-order", "MenuOrder")
	b.ANY( "/system/menu-tree", "MenuSelectTree")
	b.POST( "/system/menu-check", "MenuCheck")
	b.ANY( "/system/loglist", "LogList")
	b.ANY( "/system/log-delete", "LogDelete")
}

func (c *SystemController) MenuList() {
	if c.Ctx().URLParam("grid") == "treegrid" {
		c.Ctx().Render().JSON(models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetTree(models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetAll(), 0))
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Treegrid("system_menu_list", "/b/system/menulist?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).CurrentPos(menuid),
		"toolbar":   "system_menulist_treegrid_toolbar",
		"idField":   "operateid",
		"treeField": "name",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "10", "align": "center", "formatter": "systemMenuOrderFormatter", "index": "0"},
		"菜单名称": {"field": "name", "width": "130", "index": "1"},
		"管理操作": {"field": "operateid", "width": "25", "align": "center", "formatter": "systemMenuOperateFormatter", "index": "2"},
	})
	c.Ctx().Render().ViewData("treegrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/system_menulist.html")
}

func (c *SystemController) MenuAdd(iCache cache.ICache) {
	if c.Ctx().IsPost() {
		parentid,_:= c.Ctx().PostInt64("parentid", 0)
		name := c.Ctx().PostString("name", "")
		ctrl := c.Ctx().PostString("c", "")
		a:= c.Ctx().PostString("a", "")
		data := c.Ctx().PostString("data", "")
		display,_ := c.Ctx().PostInt64("display", 1)

		menu := &tables.IriscmsMenu{
			Parentid: parentid,
			Name:     name,
			C:        ctrl,
			A:        a,
			Data:     data,
			Display:  display,
		}
		newid, _ := c.Ctx().Value("orm").(*xorm.Engine).InsertOne(menu)
		if newid > 0 {
			clearMenuCache(iCache, c.Ctx().Value("orm").(*xorm.Engine))
			helper.Ajax("新增菜单成功", 0, c.Ctx())
		} else {
			helper.Ajax("新增菜单失败", 1, c.Ctx())
		}
		return
	}
	c.Ctx().Render().HTML("backend/system_menuadd.html")
}

func (c *SystemController) MenuDelete(iCache cache.ICache) {
	id,_ := c.Ctx().PostInt64("id", 0)
	if id < 1 {
		helper.Ajax("参数失败", 1, c.Ctx())
		return
	}
	// 查找是否有下级菜单
	exists, err := c.Ctx().Value("orm").(*xorm.Engine).Where("parentid = ?", id).Count(&tables.IriscmsMenu{})
	if err != nil {
		helper.Ajax("删除菜单失败,异常错误", 1, c.Ctx())
		return
	}
	if exists > 0 {
		helper.Ajax("删除菜单失败,有下级菜单", 1, c.Ctx())
		return
	}
	aff, _ := c.Ctx().Value("orm").(*xorm.Engine).Id(id).Delete(&tables.IriscmsMenu{})
	if aff > 0 {
		clearMenuCache(iCache, c.Ctx().Value("orm").(*xorm.Engine))
		helper.Ajax("删除菜单成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除菜单失败", 1, c.Ctx())
	}
}

func (c *SystemController) MenuOrder(iCache cache.ICache) {
	posts := c.Ctx().PostData()
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
		affected, _ := c.Ctx().Value("orm").(*xorm.Engine).Id(id).Update(menu)
		if affected > 0 {
			flag++
		}
	}
	if flag > 0 {
		clearMenuCache(iCache, c.Ctx().Value("orm").(*xorm.Engine))
		helper.Ajax("排序更新成功", 0, c.Ctx())
	} else {
		helper.Ajax("排序规则没有发生任何改变", 1, c.Ctx())
	}
}

func (c *SystemController) MenuEdit(iCache cache.ICache) {
	id, _ := c.Ctx().URLParamInt64("id")
	if id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	menu, has := models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetInfo(id)
	if !has {
		helper.Ajax("没有此菜单", 1, c.Ctx())
		return
	}
	if c.Ctx().IsPost() {
		parentid,_:= c.Ctx().PostInt64("parentid", 0)
		name := c.Ctx().PostString("name", "")
		ctrl := c.Ctx().PostString("c", "")
		a:= c.Ctx().PostString("a", "")
		data := c.Ctx().PostString("data", "")
		display,_ := c.Ctx().PostInt64("display", 1)

		menu := &tables.IriscmsMenu{
			Parentid: parentid,
			Name:     name,
			C:        ctrl,
			A:        a,
			Data:     data,
			Display:  display,
		}
		newid, _ := c.Ctx().Value("orm").(*xorm.Engine).Id(id).Update(menu)
		if newid > 0 {
			clearMenuCache(iCache, c.Ctx().Value("orm").(*xorm.Engine))
			helper.Ajax("修改菜单成功", 0, c.Ctx())
		} else {
			helper.Ajax("修改菜单失败", 1, c.Ctx())
		}
		return
	}

	c.Ctx().Render().ViewData("data", menu)
	c.Ctx().Render().HTML("backend/system_menuedit.html")
}

func (c *SystemController) MenuSelectTree() {
	var tree []map[string]interface{}
	tree = append(tree, map[string]interface{}{
		"id":       0,
		"text":     "作为一级菜单",
		"children": models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetSelectTree(models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetAll(), 0),
	})
	c.Ctx().Render().JSON(tree)
}

func (c *SystemController) MenuCheck() {
	name := c.Ctx().FormValue("name")
	if name == "" {
		helper.Ajax("用户名为空", 1, c.Ctx())
		return
	}
	if models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).CheckName(name) {
		helper.Ajax("用户名已存在", 1, c.Ctx())
		return
	}

	helper.Ajax("正常", 0, c.Ctx())
}

func (c *SystemController) LogList() {
	page, _ := c.Ctx().URLParamInt64("page")
	rows, _ := c.Ctx().URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewLogModel(c.Ctx().Value("orm").(*xorm.Engine)).GetList(page, rows)
		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("system_menu_logList", "/b/system/loglist", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).CurrentPos(menuid),
		"toolbar": "system_loglist_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"用户名": {"field": "Username", "width": "20", "index": "0"},
		"模块":  {"field": "Controller", "width": "15", "index": "1"},
		"方法":  {"field": "Action", "width": "15", "index": "2"},
		"URL": {"field": "Querystring", "width": "100", "formatter": "systemLogViewFormatter", "index": "3"},
		"时间":  {"field": "Time", "width": "30", "index": "4"},
		"IP":  {"field": "Ip", "width": "25", "index": "5"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/system_loglist.html")

}

func (c *SystemController) LogDelete() {
	if models.NewLogModel(c.Ctx().Value("orm").(*xorm.Engine)).DeleteAll() {
		helper.Ajax("清空日志成功", 0, c.Ctx())
	} else {
		helper.Ajax("清空日志失败", 1, c.Ctx())
	}
}
