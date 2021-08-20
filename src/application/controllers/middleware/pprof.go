package middleware

import (
	"github.com/valyala/fasthttp/pprofhandler"
	"github.com/xiusin/pine"
	"strings"
)

func Pprof() pine.Handler {
	return func(ctx *pine.Context) {
		p := ctx.Path()
		if !strings.HasPrefix(ctx.Path(), "statsviz") && strings.HasPrefix(p, "/debug") {
			pprofhandler.PprofHandler(ctx.RequestCtx)
			ctx.Stop()
		}
		ctx.Next()
	}
}
