package middleware

import (
	"github.com/arl/statsviz"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/xiusin/pine"
	"runtime"
	"time"
)

func StatesViz(app *pine.Application) pine.Handler {
	type stats struct {
		Mem          runtime.MemStats
		NumGoroutine int
	}
	upgrade := websocket.FastHTTPUpgrader{}
	indexHandler := fasthttpadaptor.NewFastHTTPHandler(statsviz.Index)
	app.GET("/debug/statsviz/*filepath", func(ctx *pine.Context) {
		if ctx.Params().Get("filepath") == "ws" {
			err := upgrade.Upgrade(ctx.RequestCtx, func(ws *websocket.Conn) {
				defer ws.Close()
				tick := time.NewTicker(time.Second)
				defer tick.Stop()
				var stats stats
				for range tick.C {
					runtime.ReadMemStats(&stats.Mem)
					stats.NumGoroutine = runtime.NumGoroutine()
					if err := ws.WriteJSON(stats); err != nil {
						pine.Logger().Print(err)
						break
					}
				}
			})
			if err != nil {
				if _, ok := err.(websocket.HandshakeError); ok {
					pine.Logger().Error(err)
				}
				return
			}

		} else {
			indexHandler(ctx.RequestCtx)
		}
	})
	return func(ctx *pine.Context) {
		ctx.Next()
	}
}
