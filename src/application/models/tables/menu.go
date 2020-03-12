package tables

type Menu struct {
	Id        int64 `xorm:"pk"`
	Name      string
	Parentid  int64
	C         string
	A         string
	Data      string
	IsSystem  int64
	Listorder int64
	Display   int64
}
