package router

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/frontend"
)

func InitRouter(app *pine.Application) {

	// 前端路由注册
	app.Handle(new(frontend.FescController))
	app.Handle(new(frontend.IndexController))

}
