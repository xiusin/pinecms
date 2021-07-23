package main

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"sync"
)

type echo struct {
	sync.Once
	orm       *xorm.Engine
	app       *pine.Application
	prefix    string
	isInstall bool
}

func (p *echo) Name() string {
	return ""
}

func (p *echo) Sign() string {
	return "4255f575-2903-4717-913d-1c8a160ff72c"
}

func (p *echo) Version() string {
	return "dev 0.0.1"
}

func (p *echo) Author() string {
	return ""
}

func (p *echo) Description() string {
	return ""
}

func (p *echo) Course() string {
	return ""
}

func (p *echo) Contact() string {
	return ""
}

func (p *echo) View() string {
	return ""
}

func (p *echo) Install() {

}

func (p *echo) Prefix(prefix string) {
	p.prefix = prefix
}

func (p *echo) Upgrade() {
}

func (p *echo) Init(app *pine.Application, backend *pine.Router) {
	
}

// echoPlugin 导出插件可执行变量 不可删除
var echoPlugin = echo{}