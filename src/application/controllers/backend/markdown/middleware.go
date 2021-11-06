package markdown

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
)

func InjectVar() pine.Handler {
	return func(ctx *pine.Context) {

		cacher := di.MustGet(controllers.ServiceICache).(cache.AbstractCache)

		ctx.Render().ViewData("title", config.Website.Title)
		ctx.Render().ViewData("keywords", config.Website.Keywords)
		ctx.Render().ViewData("description", config.Website.Description)

		version := ctx.Params().Get("version")
		doc := ""
		if len(version) == 0 && strings.Contains(ctx.Path(), "docs") {
			version = config.Doc.Version
			doc = config.Doc.Default
			ctx.Redirect(fmt.Sprintf("%s/docs/%s/%s", urlPrefix, version, doc), 302)
			ctx.Stop()
			return
		}
		if len(version) == 0 {
			version = config.Doc.Version
			doc = config.Doc.Default
		}

		cacheKey := strings.TrimSuffix(version, ".md")

		byts, _ := cacher.Get(cacheKey)
		byts = []byte{}
		if len(byts) > 0 {
			ctx.WriteHTMLBytes(byts)
			ctx.Stop()
			return
		}

		docfiles := strings.Split(version, "/")
		if len(docfiles) == 1 {
			version = docfiles[0]
			doc = config.Doc.Default
		} else {
			version = docfiles[0]
			doc = strings.TrimSuffix(docfiles[1], ".md") + ".md"
		}

		ctx.Render().ViewData("version", version)

		ctx.Render().ViewData("doc", doc)

		ctx.SetUserValue("version", version)
		ctx.SetUserValue("doc", doc)

		byts, _ = json.Marshal(config)
		ctx.Render().ViewData("config", template.JS(string(byts)))
		ctx.Next()
		// 记录缓存
		// cacher.Set(cacheKey, ctx.Response.Body())
	}
}
