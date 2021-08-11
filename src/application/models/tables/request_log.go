package tables

type RequestLog struct {
	Id       int64     `json:"id"`
	Params   string    `json:"params" xorm:"comment('参数') varchar(255)"`
	Uri      string    `json:"uri" xorm:"comment('请求uri') varchar(255)"`
	Userid   int64     `json:"userid" xorm:"comment('操作用户ID') int(6)"`
	Username string    `json:"username" xorm:"comment('操作用户名') varchar(100)"`
	Ip       string    `json:"ip" xorm:"comment('操作IP') varchar(15)"`
	Time     LocalTime `json:"time" xorm:"comment('操作时间')"`
	Method   string    `json:"method" xorm:"comment('请求方法') varchar(10)"`
}
