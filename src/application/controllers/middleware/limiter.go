package middleware

import (
	"time"

	"github.com/valyala/fasthttp"
	"github.com/xiusin/pine"
	"golang.org/x/time/rate"
)

func Limiter(maxBurstSize int) pine.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Second), maxBurstSize)
	return func(ctx *pine.Context) {
		if limiter.Allow() {
			ctx.Next()
		} else {
			ctx.Abort(fasthttp.StatusServiceUnavailable, "The service is temporarily unavailable")
		}
	}
}
