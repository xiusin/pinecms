package tables

type Log struct {
	Id      int64       `json:"id"`
	Level   uint8       `json:"level" xorm:"comment('日志类型') tinyint(3)"`
	Uri     string      `json:"uri" xorm:"comment('请求uri') varchar(255)"`
	Method  string      `json:"method" xorm:"comment('请求方法') varchar(10)"`
	Params  interface{} `json:"params" xorm:"comment('请求参数') json"`
	Message string      `json:"message" xorm:"comment('操作用户名') text"`
	Stack   string      `json:"stack" xorm:"comment('调用堆栈') text"`
	Ip      string      `json:"ip" xorm:"comment('操作IP') varchar(15)"`
	Time    LocalTime   `json:"time" xorm:"comment('操作时间')"`
}
