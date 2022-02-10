package filemanager

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	InitInstall(app, "/uploads/", "./resources/assets/uploads/")
	app.Handle(new(FileManagerController), "/filemanager")
}
