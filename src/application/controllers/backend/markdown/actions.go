package markdown

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	template2 "html/template"

	"github.com/88250/lute"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/template"
)

var config *Config

func init() {
	if f, err := os.ReadFile(configFilePath); err != nil {
		panic(err)
	} else {
		json.Unmarshal(f, &config)
	}
	t := template.New(viewsDir, ".tpl", true)
	t.AddFunc("url", func(url string) string {
		if strings.HasPrefix(url, "/quickui") {
			return config.Homepage + "/" + url
		}
		return config.Homepage + urlPrefix + "/" + strings.TrimLeft(url, "/")
	})
	t.AddFunc("doc", func(url string) string {
		return config.Homepage + urlPrefix + "/docs/" + strings.TrimLeft(url, "/")
	})
	pine.RegisterViewEngine(t)

}

func Docs(ctx *pine.Context) {
	version, doc := ctx.UserValue("version").(string), ctx.UserValue("doc").(string)

	mdfile := filepath.Join(docsDir, version, doc)
	mdFileIndex := filepath.Join(docsDir, version, config.Doc.Index)

	var err error
	var documentationByts []byte
	var willReadDocumentBytes []byte

	luteEngine := lute.New()

	if documentationByts, err = os.ReadFile(mdFileIndex); err == nil {
		ctx.Render().ViewData("index", template2.HTML(luteEngine.Markdown(mdFileIndex, documentationByts)))
	}
	if willReadDocumentBytes, err = os.ReadFile(mdfile); err == nil {
		ctx.Render().ViewData("content", template2.HTML(luteEngine.Markdown(mdfile, willReadDocumentBytes)))
	} else {
		ctx.Render().ViewData("content", template2.HTML(DocMiss))
	}
	ctx.Render().HTML("docs_quick.tpl")
}

func Views(ctx *pine.Context) {
	ctx.Render().HTML("index.tpl")
}
