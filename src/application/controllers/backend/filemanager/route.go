package filemanager

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	app.Handle(new(FileManagerController), "/filemanager")
}
