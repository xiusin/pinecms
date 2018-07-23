package tables

type IriscmsAdminRole struct {
	Roleid      int64 `xorm:"pk"`
	Rolename    string `xorm:"unique"`
	Description string
	Listorder   int64
	Disabled    int64
}