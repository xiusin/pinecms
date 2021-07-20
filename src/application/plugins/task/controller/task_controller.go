package controller

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/plugins/task/manager"
	"github.com/xiusin/pinecms/src/application/plugins/task/table"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

// TaskController 控制器区域
type TaskController struct {
	backend.BaseController
}

func (c *TaskController) Construct() {
	c.BaseController.Construct()
	c.Table = &table.TaskInfo{}
	c.Entries = &[]table.TaskInfo{}
	c.OpBefore = c.before
}

func (c *TaskController) before(act int, query interface{}) error {
	if act == backend.OpList {
		sess := query.(*xorm.Session)
		status := c.Input().GetInt("status")
		sess.Where("status = ?", status)
		if status == 1 {
			sess.Where("`type` = ?", c.Input().GetInt("type"))
		}
	}
	return nil
}

// PostStop 暂停任务
func (c *TaskController) PostStop() {
	c.save(0, "停止")
	task := &table.TaskInfo{}
	c.Orm.Where(c.TableKey+"=?", c.Input().GetInt("id")).Get(task)
	manager.RemoveTask(task)
}

func (c *TaskController) PostStart() {
	c.save(1, "开启")
	task := &table.TaskInfo{}
	c.Orm.Where(c.TableKey+"=?", c.Input().GetInt("id")).Get(task)
	manager.RegisterTask(task.Id, task)
}

func (c *TaskController) PostOnce() {
	task := &table.TaskInfo{}
	c.Orm.Where(c.TableKey+"=?", c.Input().GetInt("id")).Get(task)

	manager.TaskJobFunc(task)()
	helper.Ajax("任务触发成功", 0, c.Ctx())
}

func (c *TaskController) GetLog() {
	id, _ := c.Ctx().GetInt("id", 0)
	page, _ := c.Ctx().GetInt("page", 0)
	pageSize, _ := c.Ctx().GetInt("size", 0)
	status, _ := c.Ctx().GetInt("status", -1)
	sess := c.Orm.Desc("id")

	if id > 0 {
		sess.Where("task_id = ?", id)
	}
	if status >= 0 {
		sess.Where("status = ?", status)
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	sess.Limit((page-1)*pageSize, pageSize)
	var logs []table.TaskLog
	count, _ := sess.FindAndCount(&logs)
	helper.Ajax(pine.H{
		"pagination": pine.H{
			"page":  page,
			"size":  pageSize,
			"total": count,
		},
		"list": logs,
	}, 0, c.Ctx())

}

func (c *TaskController) save(status uint, act string) {
	exist, err := c.Orm.Where(c.TableKey+"=?", c.Input().GetInt("id")).Get(c.Table)
	if err != nil || !exist {
		helper.Ajax("任务不存在", 1, c.Ctx())
		return
	}
	if rest, _ := c.Orm.Where(c.TableKey+"=?", c.Input().GetInt("id")).
		Cols("status", "updated_at").
		Update(&table.TaskInfo{
			Status:    status,
			UpdatedAt: time.Now(),
		}); rest > 0 {
		helper.Ajax(act+"任务成功", 0, c.Ctx())
	} else {
		helper.Ajax(act+"任务失败", 1, c.Ctx())
	}
}
