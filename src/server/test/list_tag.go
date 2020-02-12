package test

import (
	"fmt"

	"github.com/flosch/pongo2"
)

type tagListTag struct {
	content string
}

func (t *tagListTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	writer.WriteString("hello world")
	return nil
}

func tagListTagParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (tag pongo2.INodeTag, e *pongo2.Error) {
	fmt.Println(doc, start, arguments)
	ttNode := &tagListTag{}
	return ttNode, nil
}