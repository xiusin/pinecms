package manager

import (
	"fmt"
	"github.com/xiusin/pinecms/src/application/plugins/task/table"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/robfig/cron/v3"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/common/helper"
)

var tm *taskManager

type TaskFunc func(*xorm.Engine) (string, error)

type taskManager struct {
	sync.Once
	sync.Mutex

	pool       map[int64]TaskFunc
	orm        *xorm.Engine
	logger     cron.Logger
	cron       *cron.Cron
	scriptPath string
	ExecFn     string
}

func TaskManager() *taskManager {
	if tm == nil {
		tm = &taskManager{
			logger:     &logger{pine.Logger()},
			orm:        di.MustGet(&xorm.Engine{}).(*xorm.Engine),
			scriptPath: "tasks",
			pool:       map[int64]TaskFunc{},
			ExecFn:     "tasks.Run",
		}
		_ = os.Mkdir(tm.scriptPath, os.ModePerm)
	}
	return tm
}

func (tm *taskManager) CheckErr(err error, msg string, panicErr bool, extra ...interface{}) {
	if err != nil {
		tm.logger.Error(err, msg, extra...)
		if panicErr {
			panic(err)
		}
	}
}

func (tm *taskManager) initYaegi() (i *interp.Interpreter, err error) {
	i = interp.New(interp.Options{})
	err = i.Use(stdlib.Symbols)
	tm.CheckErr(err, "导入标准库失败", false)
	err = i.Use(interp.Exports{
		"pinecms/pinecms": {
			"DB": reflect.ValueOf(xorm.Engine{}), //
			//"Table"
		},
	})
	tm.CheckErr(err, "导出自定义包异常", false)
	return i, err

}

func (tm *taskManager) Cron() {
	tm.Do(func() {
		cron.Recover(tm.logger)
		tm.cron = cron.New(
			cron.WithSeconds(),
			cron.WithLocation(helper.GetLocation()),
		)
		var tasks []table.TaskInfo
		_ = tm.orm.Table(&table.TaskInfo{}).Where("status = ?", 1).Find(&tasks)
		for _, task := range tasks {
			RegisterTask(task.Id, &task)
		}
		go tm.cron.Start()
	})
}

func taskScript(service string) string {
	return filepath.Join(tm.scriptPath, service+".gsh")
}

//RegisterTask 注册任务到任务管理对象
func RegisterTask(id int64, task *table.TaskInfo) {
	tm.Lock()
	defer tm.Unlock()
	if task == nil {
		task = &table.TaskInfo{}
		exist, _ := tm.orm.Where("id = ?", id).Get(task)
		if !exist {
			tm.logger.Error(fmt.Errorf("id:%d", id), "任务不存在")
			return
		}
	}
	if task.Status == 0 {
		return
	}
	cronExpr := task.Cron
	if task.TaskType == 1 && task.Every > 0 {
		cronExpr = fmt.Sprintf("@every %ds", task.Every)
	}
	entityId, err := tm.cron.AddFunc(cronExpr, TaskJobFunc(task))
	msg := ""
	if err != nil {
		msg = err.Error()
		tm.CheckErr(err, "注册定时任务失败", false, task)
	}
	_, err = tm.orm.Where("id = ?", task.Id).Cols("entity_id", "error").Update(&table.TaskInfo{
		EntityId: int(entityId),
		Error:    msg,
	})
	tm.CheckErr(err, "注册定时任务失败", false, task)
}

func TaskJobFunc(info *table.TaskInfo) func() {
	id := info.Id
	return func() {
		var dur time.Duration
		defer func() {
			if err := recover(); err != nil {
				_, _ = tm.orm.InsertOne(&table.TaskLog{
					TaskId:   id,
					Status:   false,
					Detail: fmt.Sprintf("%s", err),
					ExecTime: dur.Milliseconds(),
				})
			}
		}()
		task := &table.TaskInfo{}
		exist, err := tm.orm.Where("id = ?", id).Get(task)
		if err != nil || !exist {
			return
		}

		// 暂停的任务移出管理器
		if task.Status == 0 {
			RemoveTask(task)
			return
		}
		// 限制运行 仅 every生效
		if task.TaskType == 1 { // 步减
			if task.Limit == 0 {
				return
			}
			_, _ = tm.orm.Table(task).Decr("limit").Update(&table.TaskInfo{})
		}

		if task.StartDate != nil && task.StartDate.After(time.Now()) {
			return
		}

		if task.EndDate != nil && task.EndDate.Before(time.Now()) {
			return
		}

		entity := tm.cron.Entry(cron.EntryID(id))

		go func() {
			_, _ = tm.orm.Where("id = ?", id).Update(&table.TaskInfo{
				NextRunTime: &entity.Next,
			})
		}()
		tm.Lock()
		defer tm.Unlock()
		start := time.Now()

		taskSh := taskScript(info.Service)
		var fn TaskFunc
		if fn, exist = tm.pool[info.Id]; !exist { // todo 动态载入脚本修改 (是否需要一个延迟)
			if engine, err := tm.initYaegi(); err != nil {
				panic(err)
			} else {
				_, err = engine.EvalPath(taskSh)
				tm.CheckErr(err, "脚本语法错误", true)
				v, err := engine.Eval(tm.ExecFn)
				tm.CheckErr(err, "执行脚本错误", true)
				tm.pool[info.Id] = v.Interface().(func(*xorm.Engine) (string, error))
			}
		}
		msg, err := fn(tm.orm)
		tm.CheckErr(err, "脚本结果异常", true)
		dur = time.Now().Sub(start)
		go func() {
			_, err = tm.orm.InsertOne(&table.TaskLog{
				TaskId:   id,
				Status:   true,
				Detail:   msg,
				ExecTime: dur.Milliseconds(),
			})
		}()
	}
}

//RemoveTask 移除任务
func RemoveTask(task *table.TaskInfo) {
	tm.Lock()
	defer tm.Unlock()
	id := task.Id
	tm.pool[id] = nil
	delete(tm.pool, id)
	tm.cron.Remove(cron.EntryID(id))
	tm.orm.Where("id = ?", id).Cols("entity_id", "error").Table(task).Update(&table.TaskInfo{})
}
