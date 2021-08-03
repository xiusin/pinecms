package tables

type DocumentModel struct {
	Id              int64     `json:"id"`
	Name            string    `json:"table_name" xorm:"comment('模型名') varchar(100)"`
	Table           string    `json:"table" xorm:"comment('模型表名') varchar(100)"`
	Enabled         int       `json:"enabled"`
	ModelType       int64     `json:"model_type" xorm:"comment('模型类型: 1=系统模型 0=普通模型')"`
	FeTplIndex      string    `json:"fe_tpl_index" xorm:"comment('前端列表主页模板') varchar(70)"`
	FeTplList       string    `json:"fe_tpl_list" xorm:"comment('前端列表模板') varchar(70)"`
	FeTplDetail     string    `json:"fe_tpl_detail" xorm:"comment('前端详情模板') varchar(70)"`
	Remark          string    `json:"remark" api:"remark:备注" xorm:"comment('备注') text"`
	CreatedAt       LocalTime `json:"created"`
	UpdatedAt       LocalTime `json:"updated"`
	FieldShowInList string    `json:"field_show_in_list" xorm:"-"`
	FeSearchFields  string    `json:"fe_search_fields" xorm:"-"`
	Formatters      string    `json:"formatters" xorm:"-"`
	Execed          bool      `json:"execed" xorm:"comment('是否已执行') tinyint(1)"`
	DeletedAt       LocalTime `xorm:"deleted" json:"-"`
}
