package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/common/helper"
	"html/template"
)

/**
1. 文档模型管理
*/
type DocumentController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (c *DocumentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/model/list", "ModelList")
	b.Handle("ANY", "/model/add", "ModelAdd")
	b.Handle("ANY", "/model/edit", "ModelEdit")
	b.Handle("POST", "/model/delete", "ModelDelete")
}

func (c *DocumentController) ModelList() {
	page, _ := c.Ctx.URLParamInt64("page")
	rows, _ := c.Ctx.URLParamInt64("rows")

	if page > 0 {
		list, total := models.NewDocumentModel(c.Orm).GetList(page, rows)
		c.Ctx.JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("model_list_datagrid", "/b/model/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "model_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"模型名称": {"field": "name", "width": "30", "index": "0"},
		"模型ID": {"field": "table", "width": "30", "index": "1"},
		"启用": {"field": "enabled", "width": "25", "index": "2", "formatter": "enabledFormatter"},
		"操作": {"field": "id", "width": "25", "index": "3", "formatter": "optFormatter"},
	})
	c.Ctx.ViewData("dataGrid", template.HTML(table))
	c.Ctx.View("backend/model_list.html")
}

func (c *DocumentController) ModelAdd() {
	currentPos := models.NewMenuModel(c.Orm).CurrentPos(64)
	// 查找数据库模型字段
	list, _ := models.NewDocumentModelFieldModel(c.Orm).GetList(1, 1000)
	c.Ctx.ViewData("list", list)
	c.Ctx.ViewData("title", currentPos)
	c.Ctx.View("backend/model_add.html")
}

func (c *DocumentController) ModelEdit() {
	c.Ctx.View("backend/model_edit.html")
}

func (c *DocumentController) ModelDelete() {
	c.Ctx.WriteString("hello model111")
}
