package tables

type Admin struct {
	Userid        int64  `xorm:"pk autoincr" json:"userid"`
	Username      string `json:"username"`
	Password      string `json:"-"`
	Roleid        int64  `json:"roleid"`
	Encrypt       string `json:"-"`
	Lastloginip   string `json:"lastloginip"`
	Lastlogintime int64  `json:"lastlogintime"`
	Email         string `json:"email"`
	Realname      string `json:"realname"`
	Avatar        string `json:"avatar"`
}
