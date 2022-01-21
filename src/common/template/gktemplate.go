package template

import (
	"fmt"
	"github.com/gokeeptech/gktemplate"
	"io"
	"strings"
)

type GkTemplate struct {
	*GkTemplateConf
}

type GkTemplateConf struct {
	NameSpace string
	L         string
	R         string
	Path      string
	Ext       string
}

func NewGkTemplate(conf *GkTemplateConf) *GkTemplate {
	if conf == nil {
		conf = &GkTemplateConf{"pinecms", "{", "}", "resources/themes", ".html"}
	}
	if err := gktemplate.LoadDir(strings.TrimRight(conf.Path, "/") + "/*" + conf.Ext); err != nil {
		panic(err)
	}
	return &GkTemplate{conf}
}

func (g GkTemplate) Ext() string {
	return g.GkTemplateConf.Ext
}

func (g GkTemplate) AddFunc(s string, i interface{}) {
	switch i := i.(type) {
	case gktemplate.TagFunc:
		fun := map[string]gktemplate.TagFunc{s: i}
		gktemplate.ExtFuncs(&fun)
	}
}

func (g GkTemplate) HTML(writer io.Writer, name string, binding map[string]interface{}) error {
	var str string
	var err error
	if str, err = gktemplate.ParseFile(name, binding); err == nil {
		_, err = fmt.Fprintf(writer, str)
	}
	return err
}
