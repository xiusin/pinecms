package tables

type Plugin struct {
	Id          int64                    `json:"id"`
	Name        string                   `json:"name" schema:"name" xorm:"varchar(100) comment('插件名称')"`
	Author      string                   `json:"author" xorm:"varchar(100) comment('作者')"`
	Contact     string                   `json:"contact" xorm:"varchar(100) comment('联系方式')"`
	Description string                   `json:"description" xorm:"text comment('功能描述')"`
	Version     string                   `json:"version" xorm:"varchar(100) comment('版本号')"`
	Sign        string                   `json:"sign" xorm:"comment('标志') unique"`
	Path        string                   `json:"path" xorm:"comment('插件本地路径') unique"`
	Enable      bool                     `json:"enable" xorm:"comment('是否启用 0：否 1：是')"`
	Status      uint                     `json:"status" xorm:"comment('状态 0:缺少配置 1:可用 2: 配置错误 3:未知错误')"`
	View        []map[string]interface{} `json:"view" xorm:"json comment('页面配置信息')"`
	Page        string                   `json:"page" xorm:"-"`
	Config      map[string]interface{}   `json:"config" xorm:"json comment('插件绑定配置')"`
	Prefix      string                   `json:"prefix" xorm:"comment('插件访问前缀')"`
	NoInstall   bool                     `json:"no_install" xorm:"-"`
	ErrMsg      string                   `json:"errmsg" xorm:"-"`
	CreatedAt   LocalTime                `json:"created_at"`
	UpdatedAt   LocalTime                `json:"updated_at"`
}
