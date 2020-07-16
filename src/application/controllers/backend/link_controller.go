package backend

import (
	"github.com/xiusin/pine"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type LinkController struct {
	pine.Controller
}

func (c *LinkController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY( "/link/list", "List")
	b.ANY( "/link/add", "Add")
	b.ANY( "/link/edit", "Edit")
	b.ANY( "/link/delete", "Delete")
	b.ANY( "/link/order", "Order")
}

func (c *LinkController) List() {
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")

	if page > 0 {
		list, total := models.NewLinkModel().GetList(page, rows)
		//c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})

		helper.Ajax(map[string]interface{}{"rows": list, "total": total},0, c.Ctx())

		return
	}

	menuid, _ := c.Ctx().GetInt64("menuid")
	table := helper.Datagrid("link_list_datagrid", "/b/link/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel().CurrentPos(menuid),
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
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/link_list.html")
}

func (c *LinkController) Add() {
	if c.Ctx().IsPost() {
		if c.Ctx().PostString("passed") == "on" {
			c.Ctx().PostArgs().Set("passed", "1")
		}
		var d tables.Link
		if err := c.Ctx().BindForm(&d); err != nil {
			helper.Ajax("参数错误："+err.Error(), 1, c.Ctx())
			return
		}
		d.Addtime = time.Now().In(helper.GetLocation())
		if models.NewLinkModel().Add(&d) > 0 {
			helper.Ajax("添加友链成功", 0, c.Ctx())
		} else {
			helper.Ajax("添加友链失败", 1, c.Ctx())
		}
		return
	}
	imageUploader := template.HTML(helper.SiginUpload("logo", "", false, "Logo", "", ""))
	c.ViewData("imageUploader", imageUploader)
	c.Ctx().Render().HTML("backend/link_add.html")
}

func (c *LinkController) Edit() {
	if c.Ctx().IsPost() {
		if c.Ctx().PostString("passed") == "on" {
			c.Ctx().PostArgs().Set("passed", "1")
		}
		var d tables.Link
		if err := c.Ctx().BindForm(&d); err != nil || d.Linkid < 1 {
			helper.Ajax(fmt.Sprintf("参数错误:%s", err), 1, c.Ctx())
			return
		}
		d.Addtime = time.Now().In(helper.GetLocation())
		if models.NewLinkModel().Update(&d) {
			helper.Ajax("修改友链成功", 0, c.Ctx())
		} else {
			helper.Ajax("修改友链失败", 1, c.Ctx())
		}
		return
	}
	linkId, _ := c.Ctx().GetInt64("id")
	if linkId < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	links := models.NewLinkModel().Get(linkId)
	if links == nil {
		helper.Ajax("链接不存在", 1, c.Ctx())
		return
	}
	c.Ctx().Render().ViewData("link", links)
	imageUploader := template.HTML(helper.SiginUpload("logo", links.Logo, false, "Logo", "", ""))
	c.ViewData("imageUploader", imageUploader)
	c.Ctx().Render().HTML("backend/link_edit.html")
}

func (c *LinkController) Delete() {
	linkId, _ := c.Ctx().GetInt64("id")
	if linkId < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	if models.NewLinkModel().Delete(linkId) {
		helper.Ajax("删除链接成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除链接失败", 1, c.Ctx())
	}
}

func (c *LinkController) Order() {
	data := c.Ctx().PostData()
	order, ok := data["order"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	id, ok := data["id"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx())
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
		c.Ctx().Value("orm").(*xorm.Engine).ID(linkid).MustCols("listorder").Update(&tables.Link{Listorder: int64(orderNum)})
	}
	helper.Ajax("更新排序值成功", 0, c.Ctx())
}
