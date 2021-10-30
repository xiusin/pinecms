package main

import (
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/plugins/task/controller"
	"github.com/xiusin/pinecms/src/application/plugins/task/manager"
	"github.com/xiusin/pinecms/src/application/plugins/task/table"
)

type Task struct {
	sync.Once
	orm        *xorm.Engine
	isInstall  bool
	status     bool
	controller pine.IController
}

func (p *Task) Name() string {
	return "任务管理"
}

func (p *Task) Sign() string {
	return "77975e7f-de8b-4f26-be90-38c24fcd7c7d"
}

func (p *Task) Prefix() string {
	return "/pinecms-task-manager"
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

func (p *Task) SetStatus(status bool) {
	p.status = status
}

func (p *Task) IsInstall() bool {
	return p.isInstall
}

func (p *Task) Status() bool {
	return p.status
}

func (p *Task) Uninstall() {
	p.isInstall = false // 卸载标注为false
	p.status = false
	// TODO 删除一些数据表和菜单相关
}

func (p *Task) Install() {
	if !p.isInstall {
		if err := p.orm.Sync2(&table.TaskInfo{}, &table.TaskLog{}); err != nil {
			pine.Logger().Error("安装task插件数据库失败", err)
		}
		manager.TaskManager().Cron()
	}
	p.isInstall = true
}

func (p *Task) Upgrade() {

}

func (p *Task) GetController() pine.IController {
	return p.controller
}

func (p *Task) Init(services di.AbstractBuilder) {
	p.Do(func() {
    p.status = true // 默认开启
		p.orm = services.MustGet(&xorm.Engine{}).(*xorm.Engine)
		p.Install()
		_, _ = p.orm.Cols("entity_id", "error").Update(&table.TaskInfo{}) // 更新一些错误信息
		p.controller = new(controller.TaskController)
	})
}

// TaskPlugin 导出插件可执行变量
var TaskPlugin = Task{}
