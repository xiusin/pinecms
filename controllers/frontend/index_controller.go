package frontend

import (
	"github.com/kataras/iris/context"
	"github.com/go-xorm/xorm"
)

type IndexController struct {
	Orm *xorm.Engine
}

func (ctx *IndexController) Index(this context.Context) {
	this.View("frontend/index_index.html")
}
