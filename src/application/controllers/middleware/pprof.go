package middleware

import (
	"strings"

	"github.com/valyala/fasthttp/pprofhandler"
	"github.com/xiusin/pine"
)

func Pprof() pine.Handler {
	return func(ctx *pine.Context) {
		p := ctx.Path()
		if !strings.Contains(ctx.Path(), "statsviz") && strings.HasPrefix(p, "/debug") {
			pprofhandler.PprofHandler(ctx.RequestCtx)
			ctx.Stop()
			return
		}
		ctx.Next()
	}
}
