package table

import (
	"time"
)

type TaskInfo struct {
	Id          int64      `json:"id" xorm:"autoincr"`
	EntityId    int        `json:"entity_id" xorm:"comment('任务ID')"`
	RepeatConf  string     `json:"repeatConf" xorm:"comment('任务配置') text"`
	Name        string     `json:"name" xorm:"comment('任务名称') varchar(50)"`
	Cron        string     `json:"cron" xorm:"comment('cron配置') varchar(50)"`
	Every       uint       `json:"every" xorm:"comment('间隔执行时间,task_type=1时生效')"`
	Limit       uint       `json:"limit" xorm:"comment('最大执行次数 不传为无限次')"`
	Remark      string     `json:"remark" xorm:"comment('备注') text"`
	Status      uint       `json:"status" xorm:"comment('状态 0:停止 1：运行')"`
	StartDate   *time.Time `json:"startDate" xorm:"comment('开始时间')"`
	EndDate     *time.Time `json:"endDate" xorm:"comment('结束时间')"`
	Data        string     `json:"data" xorm:"comment('数据') text"`
	Service     string     `json:"service" xorm:"comment('执行的service实例ID') varchar(100)"`
	Type        uint8      `json:"type" xorm:"comment('类型 0:系统 1：用户')"`
	NextRunTime *time.Time `json:"nextRunTime" xorm:"comment('下一次执行时间')"`
	TaskType    uint       `json:"taskType" xorm:"comment('状态 0:cron 1：时间间隔')"`
	Error       string     `json:"error"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskLog struct {
	Id        int64     `json:"id"  xorm:"autoincr"`
	TaskId    int64     `json:"taskId" xorm:"comment('任务ID')"`
	Status    bool      `json:"status" xorm:"comment('状态 0:失败 1：成功')"`
	Detail    string    `json:"detail" xorm:"text comment('详情')"`
	ExecTime  int64     `json:"exec_time" xorm:"int(11) default 0 comment('执行时长')"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
