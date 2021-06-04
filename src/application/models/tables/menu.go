package tables

type Menu struct {
	Id        int64  `xorm:"pk autoincr" form:"id" json:"id"`
	Name      string `json:"name" form:"name"`
	Parentid  int64  `json:"parentId" form:"parentid"`
	C         string `json:"c" form:"c"`
	A         string `json:"a" form:"a"`
	Data      string `json:"data" form:"data"`
	IsSystem  int64  `json:"is_system" form:"is_system"`
	Listorder int64  `json:"orderNum" form:"listorder"`
	Display   int64  `json:"isShow" form:"display"`
	Type      int64  `json:"type"`
	Children  []Menu `json:"children" xorm:"-" form:"-"`
	Icon      string `json:"icon"`
	ViewPath  string `json:"viewPath"`
	KeepAlive string `json:"keepAlive"`
	Router    string `json:"router"`
}
