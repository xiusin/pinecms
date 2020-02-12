package config

import (
	"github.com/flosch/pongo2"
	"github.com/kataras/iris/v12/view"
	"github.com/xiusin/iriscms/src/application/controllers"
)

func registerDjangoFunc(engine *view.DjangoEngine)  {
	pongo2.RegisterTag("list", tagListTagParser)
}

func registerHtmlFunc(engine *view.HTMLEngine)  {
	engine.AddFunc("GetInMap", controllers.GetInMap)
}


type tagListTag struct {
	content string
}

func (t *tagListTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	writer.WriteString(t.content)
	panic("implement me")
}

func tagListTagParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (tag pongo2.INodeTag, e *pongo2.Error) {
	ttNode := &tagListTag{}
	return ttNode, nil
}