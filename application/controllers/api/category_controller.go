package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiusin/iriscms/application/controllers"
	"github.com/xiusin/iriscms/application/models"
	"github.com/xiusin/iriscms/application/models/tables"
	"github.com/xiusin/iriscms/common/helper"

	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type CategoryController struct {
	Orm       *xorm.Engine
	Ctx       iris.Context
	RedisPool *redis.Pool
}

func (c *CategoryController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/free/video/list", "FreeVideoList")
	b.Handle(iris.MethodGet, "/free/book/list", "FreeBookList")
	b.Handle(iris.MethodGet, "/paid/video/list", "PaidVideoList")
	b.Handle(iris.MethodGet, "/paid/book/list", "PaidBookList")
}

func (c *CategoryController) CategoryList() {
	client := c.RedisPool.Get()
	defer client.Close()

	//path => topCateId
	categoryPathMapToId := map[string]int64{
		"/api/v1/free/video/list": 27,
		"/api/v1/free/book/list":  28,
		"/api/v1/paid/video/list": 29,
		"/api/v1/paid/book/list":  30,
	}
	path := c.Ctx.Path()
	topCatId := categoryPathMapToId[path]
	cacheKey := fmt.Sprintf(controllers.CacheCategoryFormat, topCatId)
	datas, err := redis.String(client.Do("GET", cacheKey))
	var cats []tables.IriscmsCategory
	if err != nil {
		fmt.Println("不走缓存," + err.Error())
		cats = models.NewCategoryModel(c.Orm).GetNextCategory(topCatId)
		fmt.Println(cats)
		if len(cats) > 0 {
			client.Do("SET", cacheKey, helper.JsonEncode(cats), "EX", controllers.CacheEx)
		}
	} else {
		_ = json.Unmarshal([]byte(datas), &cats)
	}
	c.Ctx.JSON(ReturnApiData{Data: cats, Msg: "获取列表成功", Status: true})
}

func (c *CategoryController) FreeVideoList() {
	c.CategoryList()
}

func (c *CategoryController) FreeBookList() {
	c.CategoryList()
}

func (c *CategoryController) PaidVideoList() {
	c.CategoryList()
}

func (c *CategoryController) PaidBookList() {
	c.CategoryList()
}
