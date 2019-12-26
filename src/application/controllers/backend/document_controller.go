package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
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
	c.Ctx.ViewData("dataGrid", "template.HTML(table)")
	c.Ctx.View("backend/model_list.html")
}

func (c *DocumentController) ModelAdd() {
	c.Ctx.View("backend/model_add.html")
}

func (c *DocumentController) ModelEdit() {
	c.Ctx.View("backend/model_edit.html")
}

func (c *DocumentController) ModelDelete() {
	c.Ctx.WriteString("hello model111")
}
