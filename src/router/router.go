package router

import (
	"io/ioutil"

	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application) {

	// 前端路由注册
	//app.Handle(new(frontend.FescController))
	//app.Handle(new(frontend.IndexController))
	app.GET("/", func(ctx *pine.Context) {
		ctx.WriteString(string(ctx.RequestCtx.URI().Scheme()) + "://" + string(ctx.Host()) + "/admin/")
	})

	app.GET("/admin", func(ctx *pine.Context) {
		if byts, err := ioutil.ReadFile("admin/dist/index.html"); err != nil {
			ctx.Abort(500, err.Error())
		} else {
			_ = ctx.WriteHTMLBytes(byts)
		}
	})

}
