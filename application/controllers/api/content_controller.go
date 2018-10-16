package api

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iriscms/application/models/tables"
)

type ContentController struct {
	Orm *xorm.Engine
	Ctx iris.Context
}

func (c *ContentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/RngNb/list", "ContentList")
	b.Handle(iris.MethodGet, "/content/info/:id", "ContentInfo")
}

func (c *ContentController) ContentList() {
	catid, _ := c.Ctx.URLParamInt("catid")
	pageNo, _ := c.Ctx.URLParamInt("pageNo")
	pageSize, _ := c.Ctx.URLParamInt("pageSize")
	six := c.Ctx.URLParam("six")
	sort := c.Ctx.URLParam("sort")
	var arts []tables.IriscmsContent
	var offset = (pageNo - 1) * pageSize
	q := c.Orm.Where("catid = ? and deleted_at = 0 and status = 1", catid).Limit(pageSize, offset)
	if sort == "desc" {
		q.Desc(six)
	} else {
		q.Asc(six)
	}
	q.Find(&arts)
	if len(arts) == 0 {
		arts = []tables.IriscmsContent{}
	}
	c.Ctx.JSON(ReturnApiData{Status: true, Msg: "成功", Data: arts})
}

func (c *ContentController) ContentInfo() {
	id, err := c.Ctx.Params().GetInt64("id")
	if err != nil {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}
	var content tables.IriscmsContent
	ok, _ := c.Orm.Id(id).Get(&content)
	if !ok {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}

	if content.DeletedAt > 0 || content.Status != 1 {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}
	c.Ctx.JSON(ReturnApiData{Status: true, Msg: "获取资源成功", Data: content})
}
