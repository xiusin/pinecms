package webssh

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/ptemplate"
	"strings"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	InitInstall(app, urlPath("static"), "webssh/static")

	engine := ptemplate.New("webssh/view", ".html", true)
	pine.RegisterViewEngine(engine)

	app.Handle(new(SshController), urlPath("ui"))
	app.Handle(new(ApiController), urlPath("v1"))
}

func urlPath(str string) string  {
	return "/webssh/" + strings.TrimPrefix(str, "/")
}