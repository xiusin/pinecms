package backend

import (
	"html/template"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type LinkController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (c *LinkController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/link/list", "List")
	b.Handle("ANY", "/link/add", "Add")
	b.Handle("ANY", "/link/edit", "Edit")
	b.Handle("ANY", "/link/delete", "Delete")
	b.Handle("ANY", "/link/order", "Order")
}

func (c *LinkController) List() {
	page, _ := c.Ctx.URLParamInt64("page")
	rows, _ := c.Ctx.URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewLinkModel(c.Orm).GetList(page, rows)
		c.Ctx.JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("link_list_datagrid", "/b/link/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "link_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "20", "index": "0", "formatter": "linkListOrderFormatter"},
		"名称":   {"field": "name", "width": "30", "index": "1"},
		"网址":   {"field": "url", "width": "50", "index": "2"},
		"Logo": {"field": "logo", "width": "30", "index": "3", "formatter": "linkListLogoFormatter"},
		"描述":   {"field": "introduce", "width": "40", "index": "4"},
		"启用":   {"field": "passed", "width": "20", "index": "5", "formatter": "linkListEnabledFormatter"},
		"操作":   {"field": "linkid", "index": "6", "formatter": "linkListOptFormatter"},
	})
	c.Ctx.ViewData("dataGrid", template.HTML(table))
	c.Ctx.View("backend/link_list.html")
}

func (c *LinkController) Add() {
	if c.Ctx.Method() == "POST" {
		var d tables.IriscmsLink
		if err := c.Ctx.ReadForm(&d); err != nil {
			helper.Ajax("参数错误："+err.Error(), 1, c.Ctx)
			return
		}
		d.Addtime = time.Now().In(helper.GetLocation())
		if models.NewLinkModel(c.Orm).Add(&d) > 0 {
			helper.Ajax("添加友链成功", 0, c.Ctx)
		} else {
			helper.Ajax("添加友链失败", 1, c.Ctx)
		}
		return
	}
	c.Ctx.View("backend/link_add.html")
}

func (c *LinkController) Edit() {
	if c.Ctx.Method() == "POST" {
		var d tables.IriscmsLink
		if err := c.Ctx.ReadForm(&d); err != nil || d.Linkid < 1 {
			helper.Ajax("参数错误", 1, c.Ctx)
			return
		}
		d.Addtime = time.Now().In(helper.GetLocation())
		if models.NewLinkModel(c.Orm).Update(&d) {
			helper.Ajax("修改友链成功", 0, c.Ctx)
		} else {
			helper.Ajax("修改友链失败", 1, c.Ctx)
		}
		return
	}
	linkId, _ := c.Ctx.URLParamInt64("id")
	if linkId < 1 {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	link := models.NewLinkModel(c.Orm).Get(linkId)
	if link == nil {
		helper.Ajax("链接不存在", 1, c.Ctx)
		return
	}
	c.Ctx.ViewData("link", link)
	c.Ctx.View("backend/link_edit.html")
}

func (c *LinkController) Delete() {
	linkId, _ := c.Ctx.URLParamInt64("id")
	if linkId < 1 {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	if models.NewLinkModel(c.Orm).Delete(linkId) {
		helper.Ajax("删除链接成功", 0, c.Ctx)
	} else {
		helper.Ajax("删除链接失败", 1, c.Ctx)
	}
}

func (c *LinkController) Order() {
	data := c.Ctx.FormValues()
	order, ok := data["order"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	id, ok := data["id"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	for i := 0; i < len(order); i++ {
		linkid, err := strconv.Atoi(id[i])
		if err != nil {
			continue
		}
		orderNum, err := strconv.Atoi(order[i])
		if err != nil {
			continue
		}
		c.Orm.ID(linkid).MustCols("listorder").Update(&tables.IriscmsLink{Listorder: int64(orderNum)})
	}
	helper.Ajax("更新排序值成功", 0, c.Ctx)
}
