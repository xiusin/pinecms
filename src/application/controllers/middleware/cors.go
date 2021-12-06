package middleware

import (
	"github.com/valyala/fasthttp"
	"strings"

	"github.com/xiusin/pine"
)

func Cors(app *pine.Application) pine.Handler {
	app.AddRoute(fasthttp.MethodOptions, "/*any", func(ctx *pine.Context) { // 匹配所有录音的options
		return
	})

	return func(ctx *pine.Context) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", strings.TrimRight(string(ctx.RequestCtx.Referer()), "/"))
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "X-TOKEN, Content-Type, Origin, Referer, Content-Length, Access-Control-Allow-Headers, Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if !ctx.IsOptions() {
			ctx.Next()
		} else {
			ctx.Stop()
		}
	}
}
