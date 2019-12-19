package backend

import (
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type ErrorController struct {
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (_ ErrorController) ServerError(this iris.Context) {
	if this.IsAjax() {
		this.JSON(iris.Map{"errcode": 500, "errmsg": "系统发生错误: " + this.Path()})
	} else {
		this.ViewData("message", "系统发生错误 : "+this.Path())
		this.View("error.html")
	}
}

func (_ ErrorController) StatusNotFound(ctx iris.Context) {
	if ctx.IsAjax() {
		ctx.JSON(iris.Map{
			"errocde": http.StatusNotFound,
			"errmsg": "地址不存在: " + ctx.Path()})
	} else {
		ctx.ViewData("message", "无法找到路由 : "+ctx.Path()+"\r\n")
		ctx.View("error.html")
	}
}
