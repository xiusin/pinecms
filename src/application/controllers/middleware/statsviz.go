package middleware

import (
	"fmt"
	"github.com/arl/statsviz"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/xiusin/pine"
	"runtime"
	"time"
)

type stats struct {
	Mem          runtime.MemStats
	NumGoroutine int
}

var upgrader = websocket.FastHTTPUpgrader{}

func StatesViz(app *pine.Application) pine.Handler {
	indexHandler := fasthttpadaptor.NewFastHTTPHandler(statsviz.Index)
	app.GET("/debug/statsviz/*filepath", func(ctx *pine.Context) {
		if ctx.Params().Get("filepath") == "ws" {
			err := upgrader.Upgrade(ctx.RequestCtx, func(ws *websocket.Conn) {
				defer ws.Close()
				tick := time.NewTicker(time.Second)
				defer tick.Stop()
				var stats stats
				for range tick.C {
					runtime.ReadMemStats(&stats.Mem)
					stats.NumGoroutine = runtime.NumGoroutine()
					if err := ws.WriteJSON(stats); err != nil {
						fmt.Println("发送数据失败", err)
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
