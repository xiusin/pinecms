package router

import (
	"github.com/xiusin/pine"
	"io/ioutil"
)

func InitRouter(app *pine.Application) {

	// 前端路由注册
	//app.Handle(new(frontend.FescController))
	//app.Handle(new(frontend.IndexController))
	app.GET("/", func(ctx *pine.Context) {
		s, _ := ioutil.ReadFile("./dist/index.html")
		ctx.WriteHTMLBytes(s)
	})
}
