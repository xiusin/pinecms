package tables

type Category struct {
	Catid       int64          `xorm:"pk autoincr id" json:"id"`
	Parentid    int64          `json:"parentId"`
	Topid       int64          `json:"topid"`
	ModelId     int64          `json:"model_id"`
	Catname     string         `json:"name"`
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
