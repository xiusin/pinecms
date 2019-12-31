package router

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(mvcApp *mvc.Application) {
	/// 这里注册相关路由
	mvcApp.Router.Get("/", func(context context.Context) {
		///context.Text("%s", "hello world")
	})
}
