package tables

type DocumentModelDsl struct {
	Id             int64      `json:"id" xorm:"id pk autoincr"`
	Mid            int64      `json:"mid" xorm:"comment('模型ID') int(5)"`
	FieldType      int64      `json:"field_type" xorm:"comment('字段类型') int(5)"` // 字段类型ID  命名搞错了 先这样写程序吧
	FormName       string     `json:"form_name"  xorm:"comment('表单名称') varchar(50)"`
	TableField     string     `json:"table_field" xorm:"comment('表字段') varchar(50)"`
	ListOrder      int64      `xorm:"listorder comment('排序值')" json:"listorder"`
	Html           string     `json:"html" xorm:"comment('自定义html') text"`
	Required       bool       `json:"required"  xorm:"comment('是否必填') tinyint(1)"`
	Datasource     string     `json:"datasource" xorm:"comment('数据源，链接或json') text"`
	RequiredTips   string     `json:"required_tips"  xorm:"comment('必填字段信息') varchar(100)"`
	Validator      string     `json:"validator"  xorm:"comment('验证器或规则') varchar(100)"`
	Default        string     `json:"default"  xorm:"comment('默认值') varchar(100)"` //默认值
	Status         bool       `json:"status" xorm:"comment('状态 0=禁用 1=启用') tinyint(1)"`
	MainTableField bool       `json:"main_table_field" xorm:"comment('是否为主表字段') tinyint(1)"`
	Searchable     bool       `json:"searchable" xorm:"comment('是否可搜索') tinyint(1)"`
	Span           uint       `json:"span" xorm:"comment('表单span宽度') tinyint(3)"`
	Sortable       bool       `json:"sortable" xorm:"comment('是否可排序') tinyint(1)"`
	Visible        bool       `json:"visible" xorm:"comment('是否表单可见') tinyint(1)"`
	ListVisible    bool       `json:"list_visible" xorm:"comment('是否列表可见') tinyint(1)"`
	FieldLen       uint       `json:"field_len" xorm:"comment('字段长度') bigint"`
	CreatedAt      *LocalTime `xorm:"created" json:"created_at"`
	UpdatedAt      *LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt      *LocalTime `xorm:"deleted" json:"deleted_at"`
}
