package main

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/cobra"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/plugins/email/controller"
	"sync"
)

type Email struct {
	sync.Once
	isCliMode bool
	orm       *xorm.Engine
	app       *pine.Application
	prefix    string
	isInstall bool
}

func (p *Email) Name() string {
	return "邮箱管理"
}

func (p *Email) Sign() string {
	return "e79f5561-5653-4641-ae36-d3fb44fc58cf"
}

func (p *Email) Menu(table interface{}, pluginId int) {
	exist, _ := p.orm.Table(table).Where("plugin_id = ?", pluginId).Exist()
	if !exist {
		p.orm.Table(table).Insert(map[string]interface{}{
			"plugin_id":  pluginId,
			"name":       p.Name(),
			"parentid":   0,
			"c":          "email",
			"a":          "manager",
			"router":     "/sys/plugin/email",
			"icon":       "icon-menu",
			"keep_alive": 1,
			"type":       1,
			"display":    1,
			"is_system":  1,
			"view_path":  "cool/modules/task/views/plugin.vue",
		})
	}
}

func (p *Email) View() string {
	return `单独启动一个plugin.vue用于单独动态接受页面参数`
}

func (p *Email) Uninstall() {

}

func (p *Email) Install() {
	p.isInstall = true
}

func (p *Email) Prefix(prefix string) {
	p.prefix = prefix
}

func (p *Email) Upgrade() {
}

func (p *Email) Init(
	isCliMode bool,
	app *pine.Application,
	backend *pine.Router,
	rootCmd *cobra.Command,
) {
	p.Do(func() {
		p.isCliMode = isCliMode
		p.app = app
		p.orm = di.MustGet(&xorm.Engine{}).(*xorm.Engine)
		p.Install()
		if len(p.prefix) == 0 {
			p.prefix = "/email"
		}
		backend.Handle(new(controller.EmailController), p.prefix)
	})
}

// EmailPlugin 导出插件可执行变量
var EmailPlugin = Email{}
