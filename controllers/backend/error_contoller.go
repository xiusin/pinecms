package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

type ErrorController struct {
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (_ ErrorController) ServerError(this context.Context) {
	if this.IsAjax() {
		this.JSON(map[string]interface{}{"error": 1})
	} else {
		this.ViewData("message", "系统发生错误 : "+this.Path())
		this.View("error.html")
	}
}

func (_ ErrorController) StatusNotFound(ctx context.Context)  {
	if ctx.IsAjax() {
		ctx.JSON(map[string]interface{}{"error": 1})
	} else {
		ctx.ViewData("message", "无法找到路由 : "+ctx.Path()+"\r\n")
		ctx.View("error.html")
	}
}
