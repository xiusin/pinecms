package filemanager

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	InitInstall(app, "/uploads/", "./resources/assets/uploads/")
	//app.StaticFS("/fm/ui", assets, "dist", "index.html")
	app.Static("/fm/ui", "src/application/controllers/backend/filemanager/dist", 2)
	app.Handle(new(FileManagerController), "/filemanager")
}
