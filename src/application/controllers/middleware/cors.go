package middleware

import (
	"github.com/xiusin/pine"
	"strings"
)

func Cors() pine.Handler {
	return func(ctx *pine.Context) {
		//ctx.Response.Header.Add("Vary", "Access-Control-Allow-Methods")
		//ctx.Response.Header.Add("Vary", "Access-Control-Allow-Headers")
		//ctx.Response.Header.Add("Vary", "Access-Control-Allow-Credentials")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", strings.TrimRight(string(ctx.RequestCtx.Referer()), "/"))
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "X-TOKEN, Content-Type, Origin, Referer, Content-Length, Access-Control-Allow-Headers")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if !ctx.IsOptions() {
			ctx.Next()
		}
	}
}
