package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/xiusin/iriscms/src/common/helper"
)

func ViewRequestPath(app *iris.Application, path string) func(ctx context.Context) {
	return func(ctx context.Context) {
		app.Logger().SetOutput(helper.NewLogFile(path))
		ctx.Next()
	}
}
