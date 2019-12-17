package frontend

import (
	"github.com/go-xorm/xorm"
	"github.com/gorilla/sessions"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type IndexController struct {
	Orm     *xorm.Engine
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/", "Index")
	b.Handle("ANY", "/site/close", "Close")
	b.Handle("ANY", "/resume", "Resume")
}

func (c *IndexController) Index() {
	c.Ctx.View("frontend/index.html")
}


func (c *IndexController) Close() {
	c.Ctx.HTML("站点关闭,敬请谅解")
}


func (c *IndexController) Resume() {
	c.Ctx.ViewData("site", map[string]interface{}{
		"name":      "陈成彬",
		"location":  "商丘",
		"company":   "网家科技",
		"position":  "PHP工程师",
		"github":    "https://github.com/xiusin/",
		"facebook":  "",
		"phone":     "17610053500",
		"email":     "chenchengbin92@gmail.com",
		"baseurl":   "/resume",
		"permalink": "/:year/:month/:day/:title/",
		"exclude":   []string{"README.md"},
		"markdown":  "kramdown",
	})

	c.Ctx.ViewLayout("frontend/_layouts/default.html")
	c.Ctx.View("frontend/index_resume.html")
}
