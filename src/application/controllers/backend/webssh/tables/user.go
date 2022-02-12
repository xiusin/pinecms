package tables

type SSHUser struct {
	Id       int64   `xorm:"pk autoincr"`
	Phone    int64   `xorm:"unique"`
	Email    *string `xorm:"unique"`
	Password string
	Servers  []SSHServer
}
