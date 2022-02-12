package webssh

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/ptemplate"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/middleware"
	"strings"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	InitInstall(app, urlPath("static"), "webssh/static")
	engine := ptemplate.New("webssh/view", ".html", true)

	pine.RegisterViewEngine(engine)

	webssh := app.Group(urlPath(), middleware.Auth())
	{
		webssh.Handle(new(SshController), "/ui")
		webssh.Handle(new(ApiController), "/v1")
	}
}

func urlPath(str ...string) string  {
	if len(str) == 0 {
		return "/webssh"
	}
	return "/webssh/" + strings.TrimPrefix(str[0], "/")
}