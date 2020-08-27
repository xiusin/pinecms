package tables

type Menu struct {
	Id        int64  `xorm:"pk autoincr" form:"id" json:"id"`
	Name      string `json:"name" form:"name"`
	Parentid  int64  `json:"parentid" form:"parentid"`
	C         string `json:"c" form:"c"`
	A         string `json:"a" form:"a"`
	Data      string `json:"data" form:"data"`
	IsSystem  int64  `json:"is_system" form:"is_system"`
	Listorder int64  `json:"listorder" form:"listorder"`
	Display   int64  `json:"display" form:"display"`
	Children  []Menu `json:"children" xorm:"-" form:"-"`
}
