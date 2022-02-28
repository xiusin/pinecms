package middleware

import (
	"github.com/xiusin/pine"
	"strings"
)

func Cors() pine.Handler {
	return func(ctx *pine.Context) {
		hostOrigin := strings.TrimRight(string(ctx.RequestCtx.Referer()), "/")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", hostOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "X-TOKEN, Content-Type, Origin, Referer, Content-Length, Access-Control-Allow-Headers, Authorization, x-requested-with")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if !ctx.IsOptions() {
			ctx.Next()
		} else {
			ctx.Stop()
		}
	}
}
