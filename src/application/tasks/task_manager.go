package tasks

import (
	"github.com/robfig/cron/v3"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"sync"
)

type TaskIntf interface {
	Exec() error
}

var tm = newTaskManager()

type taskManager struct {
	sync.Once
	sync.RWMutex
	tasks map[uint]TaskIntf
	cron  *cron.Cron
}

func newTaskManager() *taskManager {
	tm := &taskManager{}
	tm.Cron()
	return tm
}

func (tm *taskManager) Cron() {
	tm.Do(func() {
		tm.cron = cron.New(
			cron.WithLocation(helper.GetLocation()),
			cron.WithParser(cron.NewParser(cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)),
			cron.WithLogger(&logger{pine.Logger()}),
		)
	})
	tm.cron.Start()
}

//RegisterTask 注册任务到任务管理对象
func RegisterTask(id uint, t TaskIntf) {
	tm.Lock()
	defer tm.Unlock()
	tm.tasks[id] = t
}

//RemoveTask 移除任务
func RemoveTask(id uint) {
	tm.Lock()
	defer tm.Unlock()
	tm.tasks[id] = nil
	delete(tm.tasks, id)
}
