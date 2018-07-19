package frontend

import (
	"github.com/go-xorm/xorm"
	"github.com/gorilla/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Orm     *xorm.Engine
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/", "Index")
	b.Handle("ANY", "/jianli", "JianLi")
}

func (c *IndexController) Index() {
	c.Ctx.View("frontend/index_index.html")
}

func (c *IndexController) JianLi() {
	c.Ctx.ViewData("site", map[string]interface{}{
		"name":      "Lazy007",
		"location":  "河南 * 商丘",
		"company":   "北京网家科技&东易日盛集团",
		"position":  "PHP开发工程师",
		"github":    "https://github.com/lazy007/",
		"facebook":  "",
		"phone":     "17610053500",
		"email":     "chenchengbin92@gmail.com",
		"baseurl":   "/jianli",
		"permalink": "/:year/:month/:day/:title/",
		"exclude":   []string{"README.md"},
		"markdown":  "kramdown",
	})

	c.Ctx.ViewLayout("frontend/_layouts/default.html")
	c.Ctx.View("frontend/index_jianli.html")
}
