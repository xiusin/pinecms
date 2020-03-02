package router

import (
	"github.com/xiusin/pine"
)

func InitRouter(app *pine.Application) {
	/// 这里注册相关路由
	app.GET("/", func(ctx *pine.Context) {
		ctx.Render().ViewData("link", func(s string) string { return s + " world" })
		ctx.Render().ViewData("link_slice", func(s string) []string { return []string{s, "world"} })
		ctx.Render().ViewData("link_map", func(s string) map[string]string { return map[string]string{s: "world"} })
		ctx.Render().ViewData("link_map_struct", func(s string) map[string]struct {
			FirstName string
			LastName  string
		} {
			return map[string]struct {
				FirstName string
				LastName  string
			}{
				s: {FirstName: "chen", LastName: "chengbin"},
			}
		})

		ctx.Render().HTML("frontend/index.jet")
	})
}
