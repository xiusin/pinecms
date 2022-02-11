package tables

type SSHUser struct {
	Id       int64
	Phone    int     `xorm:"unique"`
	Email    *string `xorm:"unique"`
	Password string
	Servers  []SSHServer
}
