package tables

type Menu struct {
	Id        int64  `xorm:"pk autoincr" form:"id" json:"id" api:"remark:菜单ID"`
	Name      string `json:"name" form:"name" api:"remark:菜单名称"`
	Parentid  int64  `json:"parentId" form:"parentid" api:"remark:父菜单ID"`
	C         string `json:"c" form:"c"  api:"remark:控制器（旧版）"`
	A         string `json:"a" form:"a"  api:"remark:操作（旧版）"`
	Data      string `json:"data" form:"data" api:"remark:菜单附加参数"`
	IsSystem  uint64 `json:"is_system" form:"is_system" api:"remark:是否为系统菜单"`
	Listorder int64  `json:"orderNum" form:"listorder" api:"remark:排序号"`
	Display   bool   `json:"isShow" form:"display" api:"remark:是否显示"`
	Type      int64  `json:"type" api:"remark:菜单类型"`
	Children  []Menu `json:"children" xorm:"-" form:"-"`
	Icon      string `json:"icon" api:"remark:图标"`
	ViewPath  string `json:"viewPath" api:"remark:视图路径"`
	KeepAlive bool   `json:"keepAlive" api:"remark:路由缓存"`
	Router    string `json:"router" api:"remark:路由地址"`
}
