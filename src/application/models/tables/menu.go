package tables

type Menu struct {
	Id             int64  `xorm:"pk autoincr" json:"id" api:"remark:菜单ID"`
	Name           string `json:"name" api:"remark:菜单名称"`
	Parentid       int64  `json:"parentId" api:"remark:父菜单ID"`
	PluginId       int64  `json:"plugin_id" api:"remark: 插件ID,用于标记插件创建菜单,卸载时使用"`
	Listorder      int64  `json:"orderNum" api:"remark:排序号"`
	Display        bool   `json:"isShow" api:"remark:是否显示"`
	Type           int64  `json:"type" api:"remark:菜单类型" xorm:"comment('类型 0：目录 1：菜单 2：按钮')"`
	Children       []Menu `json:"children" xorm:"-" form:"-"`
	Icon           string `json:"icon" api:"remark:图标"`
	ViewPath       string `json:"viewPath" api:"remark:视图路径"`
	KeepAlive      bool   `json:"keepAlive" api:"remark:路由缓存"`
	Router         string `json:"router" api:"remark:路由地址"`
	Perms          string `json:"perms"  xorm:"comment('权限标识')"`
	Identification string `json:"identification" xorm:"comment('权限标识, 查询时唯一索引') unique"`

	//C              string `json:"c" form:"c"  api:"remark:控制器（旧版）"`
	//A              string `json:"a" form:"a"  api:"remark:操作（旧版）"`
	//Data           string `json:"data" form:"data" api:"remark:菜单附加参数"`
	//IsSystem       uint64 `json:"is_system" form:"is_system" api:"remark:是否为系统菜单"`
}
