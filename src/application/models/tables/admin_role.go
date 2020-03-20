package tables

type AdminRole struct {
	Roleid      int64  `xorm:"pk"`
	Rolename    string
	Description string
	Listorder   int64
	Disabled    int64
}
