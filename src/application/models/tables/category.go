package tables

// todo cat 对应的表名字段冗余
type Category struct {
	Catid       int64          `xorm:"pk autoincr" json:"Catid"`
	Parentid    int64          `json:"Parentid"`
	Topid       int64          `json:"topid"`
	ModelId     int64          `json:"model_id"`
	Catname     string         `json:"Catname"`
	Type        int64          `json:"Type"`
	Keywords    string         `json:"Keywords"`
	Description string         `json:"Description"`
	Content     string         `xorm:"-"`
	Thumb       string         `json:"Thumb"`
	Dir         string         `json:"dir"`
	Url         string         `json:"Url"`
	Listorder   int64          `json:"Listorder"`
	Ismenu      int64          `json:"Ismenu"`
	ListTpl     string         `json:"list_tpl"`
	DetailTpl   string         `json:"detail_tpl"`
	UrlPrefix   string         `xorm:"-" json:"url_prefix"`
	Active      bool           `xorm:"-"`
	HasSon      bool           `xorm:"-"`
	Model       *DocumentModel `xorm:"-" json:"model"`
	*Page       `xorm:"-"`
}
