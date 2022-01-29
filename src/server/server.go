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

	router.InitApiRouter(app)
	router.InitStatics(app)
	router.InitRouter(app)

	go plugins.Init()

	app.Run(
		pine.Addr(fmt.Sprintf("%s:%d", "127.0.0.1", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))),
		pine.WithServerName("xiusin/pinecms"),
		pine.WithCookie(true),
	)
}
