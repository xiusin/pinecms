package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type ErrorController struct {
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (_ ErrorController) ServerError(this iris.Context) {
	fmt.Println("ServerError")
	if this.IsAjax() {
		this.JSON(map[string]interface{}{"error": 1})
	} else {
		this.ViewData("message", "系统发生错误 : "+this.Path())
		this.View("error.html")
	}
}

func (_ ErrorController) StatusNotFound(ctx iris.Context) {
	fmt.Println("StatusNotFound", ctx.Path(), "method", ctx.Method())
	if ctx.IsAjax() {
		ctx.JSON(map[string]interface{}{"error": 1})
	} else {
		ctx.ViewData("message", "无法找到路由 : "+ctx.Path()+"\r\n")
		ctx.View("error.html")
	}
}
