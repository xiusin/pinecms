package backend

import "github.com/xiusin/pine"

type TestController struct {
	BaseController
}

//todo 通过前置控制器方法处理

func (t *TestController) Construct()  {

}

func (c *TestController) RegisterRoute(b pine.IRouterWrapper) {
	b.GET("/list", "List")
	b.POST("/add", "Add")
	b.ANY("/edit", "Edit")
	b.ANY("/order", "Order")
	b.ANY("/delete", "Delete")
}