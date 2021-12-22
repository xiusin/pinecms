package frontend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
)

type FescController struct {
	pine.Controller
}

func (c *FescController) RegisterRoute(b pine.IRouterWrapper) {
	// 必须放到最后 否则搜索路由时会优先被此路由拦截到
	b.GET("/fesc/down-list.go", "DownList")
}

func (c *FescController) DownList() {
	aid, _ := c.Ctx().GetInt("id", 0)
	tid, _ := c.Ctx().GetInt64("tid", 0)
	cat, _ := models.NewCategoryModel().GetCategoryFByIdForBE(tid)
	if cat == nil {
		c.Ctx().Abort(404)
		return
	}
	sess := getOrmSess(cat.Model).Where("id = ?", aid).Limit(1)
	data, _ := sess.QueryString()
	c.ViewData("data", data[0])
	c.View(template("down_list.jet"))
}
