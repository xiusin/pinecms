package main

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
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
	return "77975e7f-de8b-4f26-be90-38c24fcd7c7d"
}

func (p *Task) Menu(table interface{}, pluginId int) {
	exist, _ := p.orm.Table(table).Where("plugin_id = ?", pluginId).Exist()
	if !exist {
		p.orm.Table(table).Insert(map[string]interface{}{
			"plugin_id":  pluginId,
			"name":       p.Name(),
			"parentid":   0,
			"c":          "task",
			"a":          "list",
			"router":     "/sys/plugin/task",
			"icon":       "icon-menu",
			"keep_alive": 1,
			"type":       1,
			"display":    1,
			"is_system":  1,
			"view_path":  "cool/modules/task/views/task.vue",
		})
	}
}

func (p *Task) View() string {
	return `[
  {
    "label": "accessKeyId",
    "prop": "accessKeyId",
    "component": {
      "name": "el-input",
      "attrs": {
        "placeholder": "阿里云accessKeyId"
      }
    },
    "props": {
      "label-width": "130px"
    },
    "rules": {
      "required": true,
      "message": "值不能为空"
    }
  },
  {
    "label": "accessKeySecret",
    "prop": "accessKeySecret",
    "component": {
      "name": "el-input",
      "attrs": {
        "placeholder": "阿里云accessKeySecret"
      }
    },
    "props": {
      "label-width": "130px"
    },
    "rules": {
      "required": true,
      "message": "值不能为空"
    }
  },
  {
    "label": "bucket",
    "prop": "bucket",
    "component": {
      "name": "el-input",
      "attrs": {
        "placeholder": "阿里云oss的bucket"
      }
    },
    "props": {
      "label-width": "130px"
    },
    "rules": {
      "required": true,
      "message": "值不能为空"
    }
  },
  {
    "label": "endpoint",
    "prop": "endpoint",
    "component": {
      "name": "el-input",
      "attrs": {
        "placeholder": "阿里云oss的endpoint"
      }
    },
    "props": {
      "label-width": "130px"
    },
    "rules": {
      "required": true,
      "message": "值不能为空"
    }
  },
  {
    "label": "timeout",
    "prop": "timeout",
    "value": "3600s",
    "component": {
      "name": "el-input",
      "attrs": {
        "placeholder": "阿里云oss的timeout"
      }
    },
    "props": {
      "label-width": "130px"
    },
    "rules": {
      "required": true,
      "message": "值不能为空"
    }
  }
]`
}

func (p *Task) Uninstall() {

}

func (p *Task) Install() {
	if !p.isInstall {
		if err := p.orm.Sync2(&table.TaskInfo{}, &table.TaskLog{}); err != nil {
			pine.Logger().Error("安装task插件数据库失败", err)
		}
		pine.Logger().Print("[plugin:task] 启动定时任务")
		_, _ = p.orm.Cols("entity_id", "error").Update(&table.TaskInfo{})
		manager.TaskManager().Cron()
	}
	p.isInstall = true
}

func (p *Task) Prefix(prefix string) {
	p.prefix = prefix
}

func (p *Task) Upgrade() {
}

func (p *Task) Init(
	app *pine.Application,
	backend *pine.Router,
) {
	p.Do(func() {
		//rootCmd.AddCommand(cmd.TaskCmd)
		p.app = app
		p.orm = di.MustGet(&xorm.Engine{}).(*xorm.Engine)
		p.Install()
		if len(p.prefix) == 0 {
			p.prefix = "/task"
		}
		backend.Handle(new(controller.TaskController), p.prefix)
	})
}

// TaskPlugin 导出插件可执行变量
var TaskPlugin = Task{}
