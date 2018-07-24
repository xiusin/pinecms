package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func FrontendGlobalViewData(app *iris.Application) func(ctx context.Context) {
	return func(ctx context.Context) {
		ctx.ViewData("setting", nil)
		app.Logger().Println("后端使用 groupcache 缓存前端需要的全局数据,在这里也可以添加公共函数和缓存 header")
		ctx.Next()
	}
}
