package main

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/cobra"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/plugins/task/cmd"
	"github.com/xiusin/pinecms/src/application/plugins/task/controller"
	"github.com/xiusin/pinecms/src/application/plugins/task/manager"
	"github.com/xiusin/pinecms/src/application/plugins/task/table"
	"sync"
)

type Task struct {
	sync.Once
	orm       *xorm.Engine
	app       *pine.Application
	prefix    string
	isInstall bool
}

func (p *Task) Name() string {
	return "任务管理"
}

func (p *Task) Sign() string {
	return "product_plugin"
}

func (p *Task) Version() string {
	return "dev 0.0.1"
}

func (p *Task) Author() string {
	return "xiusin"
}

func (p *Task) Description() string {
	return "实现任务管理功能"
}

func (p *Task) Course() string {
	return `# 任务管理模块教程

## 安装 
1. 下载编译完成的**task.so**放置到程序同级目录**plugins**下
2. 程序每隔10秒扫描一次目录,自动注册程序信息到插件系统
3. 本模块需要在开发环境时导入, 需要自动导出**pages.zip** 放到前端开发目录下使用

## 修改路由前缀
1. 默认路由前缀和插件名称一致,如果修改需要重启程序以载入.so实现.  
2. 需要重新导入router.ts注入修改后的路由地址 (不建议修改)

## 结束语
如有问题望请各位PR或提交到个人邮箱或电话
`
}

func (p *Task) Contact() string {
	return "18888888888"
}

func (p *Task) View() string {
	return `{}`
}

func (p *Task) Install() {
	if err := p.orm.Sync2(&table.TaskInfo{}, &table.TaskLog{}); err != nil {
		pine.Logger().Error("安装task插件数据库失败", err)
	}
	p.orm.Cols("entity_id", "error").Update(&table.TaskInfo{})

	manager.TaskManager().Cron()
}

func (p *Task) Prefix(prefix string) {
	p.prefix = prefix
}

func (p *Task) Upgrade() {
}

func (p *Task) Init(
	app *pine.Application,
	backend *pine.Router,
	rootCmd *cobra.Command,
) {
	p.Do(func() {
		rootCmd.AddCommand(cmd.TaskCmd)

		p.orm = di.MustGet(&xorm.Engine{}).(*xorm.Engine)
		if len(p.prefix) == 0 {
			p.prefix = "/task"
		}
		backend.Handle(new(controller.TaskController), p.prefix)
	})
}

// TaskPlugin 导出插件可执行变量
var TaskPlugin = Task{}
