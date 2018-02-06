package main

import (
	//"fmt"
	 . "iriscms/config"
	//"os"
	//"plugin"
	//"runtime"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app:= newApp()
	go app.Run(iris.Addr("0.0.0.0:8089"))
	StartApplication()

}

/**
测试服务
 */
func newApp() *iris.Application {
	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		ctx.WriteString("Hello World")
	})

	return app
}

