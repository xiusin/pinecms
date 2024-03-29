package server

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/plugins"
	"github.com/xiusin/pinecms/src/router"
)

func Server() {
	InitCache()

	pine.SetControllerDefaultAction("Index")
	router.InitApiRouter(app)
	router.InitStatics(app)

	go plugins.Init()

	// 内部托管任意路由
	router.InitRouter(app)

	app.Run(
		pine.Addr(fmt.Sprintf("%s:%d", "127.0.0.1", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))),
		pine.WithServerName("xiusin/pinecms"),
		pine.WithoutStartupLog(true),
		pine.WithCookie(true),
		pine.WithMaxMultipartMemory(100 * 1024 * 1024),
	)
}
