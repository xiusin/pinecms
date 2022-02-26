package router

import (
	"github.com/xiusin/pinecms/src/application/controllers/frontend"
	"path/filepath"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/config"
)

func InitStatics(app *pine.Application) {
	for _, static := range config.App().Statics {
		pine.Logger().Debug("注册静态资源路由", static.Route, "->", filepath.FromSlash(static.Path))
		app.Static(static.Route, filepath.FromSlash(static.Path), 1)
	}
}

func InitRouter(app *pine.Application) {
	// 前端路由注册
	app.Handle(new(frontend.IndexController))
}
