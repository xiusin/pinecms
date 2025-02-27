package router

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend/filemanager"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh"
	"github.com/xiusin/pinecms/src/application/controllers/backend/wechat"
)

func InitModuleRouter(backendRouter *pine.Router, app *pine.Application) {
	filemanager.InitRouter(app, backendRouter)
	webssh.InitRouter(app, backendRouter)
}

func InitSubModuleRouter(app *pine.Application, admin *pine.Router) {
	wechat.InitRouter(app, admin)
}
