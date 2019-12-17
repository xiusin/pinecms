package api

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
)

type CategoryController struct {
	Orm   *xorm.Engine
	Ctx   iris.Context
	cache *boltdb.Database
}

//func (c *CategoryController) BeforeActivation(b mvc.BeforeActivation) {
//	b.Handle(iris.MethodGet, "/free/video/list", "FreeVideoList")
//	b.Handle(iris.MethodGet, "/free/book/list", "FreeBookList")
//	b.Handle(iris.MethodGet, "/paid/video/list", "PaidVideoList")
//	b.Handle(iris.MethodGet, "/paid/book/list", "PaidBookList")
//}

//func (c *CategoryController) CategoryList() {
//	client := c.cache.Service
//	//path => topCateId
//	categoryPathMapToId := map[string]int64{
//		"/api/v1/free/video/list": 27,
//		"/api/v1/free/book/list":  28,
//		"/api/v1/paid/video/list": 29,
//		"/api/v1/paid/book/list":  30,
//	}
//	path := c.Ctx.Path()
//	topCatId := categoryPathMapToId[path]
//	cacheKey := fmt.Sprintf(controllers.CacheCategoryFormat, topCatId)
//	datas, err := redis.String(client.Do("GET", cacheKey))
//	var cats []tables.IriscmsCategory
//	if err != nil {
//		fmt.Println("不走缓存," + err.Error())
//		cats = models.NewCategoryModel(c.Orm).GetNextCategory(topCatId)
//		fmt.Println(cats)
//		if len(cats) > 0 {
//			client.Do("SET", cacheKey, helper.JsonEncode(cats), "EX", controllers.CacheEx)
//		}
//	} else {
//		_ = json.Unmarshal([]byte(datas), &cats)
//	}
//	c.Ctx.JSON(ReturnApiData{Data: cats, Msg: "获取列表成功", Status: true})
//}
//
//func (c *CategoryController) FreeVideoList() {
//	c.CategoryList()
//}
//
//func (c *CategoryController) FreeBookList() {
//	c.CategoryList()
//}
//
//func (c *CategoryController) PaidVideoList() {
//	c.CategoryList()
//}
//
//func (c *CategoryController) PaidBookList() {
//	c.CategoryList()
//}
