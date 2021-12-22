package router

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/config"
	"io/ioutil"
	"path/filepath"
)


func InitStatics(app *pine.Application)  {
	for _, static := range config.App().Statics {
		app.Static(static.Route, filepath.FromSlash(static.Path), 1)
	}
}

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
