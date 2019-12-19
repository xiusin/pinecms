package config

import "github.com/kataras/iris/v12/context"

func initRouter() {
	/// 这里注册相关路由
	mvcApp.Router.Get("/", func(context context.Context) {
		context.Text("%s", "hello world")
	})
}
