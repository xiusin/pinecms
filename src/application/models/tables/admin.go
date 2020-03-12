package tables

type Admin struct {
	Userid        int64  `xorm:"pk"`
	Username      string
	Password      string
	Roleid        int64
	Encrypt       string
	Lastloginip   string
	Lastlogintime int64
	Email         string
	Realname      string
}
