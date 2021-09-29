package mywebsql

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)

func InitRouter(app *pine.Application, router *pine.Router) {
	// 注册静态地址
	app.Static("/mywebsql/", helper.GetRootPath("mywebsql"), 1)

	// 注册路由
	app.ANY("/mywebsql/cache", Cache)
	app.Handle(new(IndexController), "/mywebsql/index")
}
