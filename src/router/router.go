package router

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application) {

	// 前端路由注册
	//app.Handle(new(frontend.FescController))
	//app.Handle(new(frontend.IndexController))
	app.GET("/", func(ctx *pine.Context) {
		ctx.WriteString(string(ctx.RequestCtx.URI().Scheme()) + "://" + string(ctx.Host()) + "/admin/")
	})

}
