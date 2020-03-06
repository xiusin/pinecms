package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/xiusin/pinecms/src/common/helper"
)

func ViewRequestPath(app *iris.Application, path string) func(ctx context.Context) {
	return func(ctx context.Context) {
		app.Logger().SetOutput(helper.NewLogFile(path))
		ctx.Next()
	}
}
