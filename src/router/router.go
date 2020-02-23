package router

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application) {
	/// 这里注册相关路由
	app.GET("/", func(context *pine.Context) {})
}
