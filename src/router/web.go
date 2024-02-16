package router

import (
	"path/filepath"

	"github.com/xiusin/pinecms/src/application/controllers/frontend"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/config"
)

func InitStatics(app *pine.Application) {
	for _, static := range config.App().Statics {
		app.Static(static.Route, filepath.FromSlash(static.Path), 1)
	}
}

func InitRouter(app *pine.Application) {
	app.Handle(new(frontend.IndexController))
}
