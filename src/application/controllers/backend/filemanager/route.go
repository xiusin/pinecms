package filemanager

import (
	"embed"

	"github.com/xiusin/pine"
)

////go:embed dist
var assets embed.FS

func InitRouter(app *pine.Application, router *pine.Router) {
	InitInstall(app, "/uploads/", "./resources/assets/uploads/")
	app.StaticFS("/fm/ui", assets, "dist", "index.html")
	//app.Static("/fm/ui", "src/application/controllers/backend/filemanager/dist", 3)

	app.Handle(new(FileManagerController), "/filemanager")
}
