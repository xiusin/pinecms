package tables

// Category 分类
type Category struct {
	Catid       int64          `xorm:"pk autoincr id" json:"id"`
	Parentid    int64          `json:"parentId" xorm:"comment('所属栏目ID')"`
	Topid       int64          `json:"topid" xorm:"comment('顶级栏目ID')"`
	ModelId     int64          `json:"model_id" xorm:"comment('绑定模型ID')"`
	Catname     string         `json:"name" xorm:"comment('分类ID')"`
	Type        int64          `json:"type"`
	Keywords    string         `json:"keywords"`
	Description string         `json:"description"`
	Content     string         `xorm:"-"`
	Thumb       string         `json:"thumb"`
	Dir         string         `json:"dir"`
	Url         string         `json:"url"`
	Listorder   int64          `json:"listorder"`
	Ismenu      bool           `json:"ismenu"`
	ListTpl     string         `json:"list_tpl"`
	DetailTpl   string         `json:"detail_tpl"`
	UrlPrefix   string         `xorm:"-" json:"url_prefix"`
	Active      bool           `xorm:"-"`
	HasSon      bool           `xorm:"-"`
	Model       *DocumentModel `xorm:"-" json:"model"`
	Page        *Page          `xorm:"-"`
}
