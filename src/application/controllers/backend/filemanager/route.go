package filemanager

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	InitInstall(app, "/uploads/", "./resources/assets/uploads/")
	FileMangerWebRouter(app)
	app.Handle(new(FileManagerController), "/filemanager")
}

////go:embed dist
//var assets embed.FS

func FileMangerWebRouter(app *pine.Application)  {
	//app.StaticFS("/fm/ui", assets, "index.html")
}
