package api

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

type CategoryController struct {
	Orm *xorm.Engine
	Ctx iris.Context
}

func (c *CategoryController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/free/video/list", "CategoryList")
	b.Handle(iris.MethodGet, "/free/book/list", "CategoryList")
	b.Handle(iris.MethodGet, "/paid/video/list", "CategoryList")
	b.Handle(iris.MethodGet, "/paid/book/list", "CategoryList")
}

func (c *CategoryController) CategoryList() {
	//path => topCateId
	categoryPathMapToId := map[string]int64{
		"/api/v1/free/video/list": 11,
		"/api/v1/free/book/list":  12,
		"/api/v1/paid/video/list": 2,
		"/api/v1/paid/book/list":  8,
	}

	path := c.Ctx.Path()

	topCatId := categoryPathMapToId[path]

	c.Ctx.JSON(ReturnApiData{Msg: strconv.Itoa(int(topCatId))})
}
