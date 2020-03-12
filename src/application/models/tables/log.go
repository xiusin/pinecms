package tables

type Log struct {
	Logid       int64 `xorm:"pk"`
	Controller  string
	Action      string
	Querystring string
	Userid      int64
	Username    string
	Ip          string
	Time        string
}
