package api

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iriscms/application/controllers"
	"iriscms/application/models"
	"iriscms/application/models/tables"
	"encoding/json"
	"iriscms/common/helper"
)

type CategoryController struct {
	Orm *xorm.Engine
	Ctx iris.Context
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
		"/api/v1/free/video/list": 32,
		"/api/v1/free/book/list":  31,
		"/api/v1/paid/video/list": 30,
		"/api/v1/paid/book/list":  29,
	}
	path := c.Ctx.Path()
	topCatId := categoryPathMapToId[path]
	cacheKey := fmt.Sprintf(controllers.CACHE_CATEGORY_FORMAT, topCatId)
	datas, err := redis.String(client.Do("GET", cacheKey))
	var cats []tables.IriscmsCategory
	if err != nil {
		fmt.Println("不走缓存," + err.Error())
		cats = models.NewCategoryModel(c.Orm).GetNextCategory(topCatId)
		if len(cats) > 0 {
			client.Do("SET", cacheKey, helper.JsonEncode(cats), "EX", controllers.CACHE_EX)
		}
	} else {
		json.Unmarshal([]byte(datas), &cats)
	}
	c.Ctx.JSON(ReturnApiData{Data: cats, Msg: "获取列表成功",Status:true})
}

func  (c *CategoryController) FreeVideoList()  {
	c.CategoryList()
}

func  (c *CategoryController) FreeBookList()  {
	c.CategoryList()
}

func  (c *CategoryController) PaidVideoList()  {
	c.CategoryList()
}

func  (c *CategoryController) PaidBookList()  {
	c.CategoryList()
}