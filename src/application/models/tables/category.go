package tables

// todo cat 对应的表名字段冗余
type Category struct {
	Catid                int64          `xorm:"pk autoincr" json:"Catid"`
	Parentid             int64          `json:"Parentid"`
	Catname              string         `json:"Catname"`
	Type                 int64          `json:"Type"`
	ModelId              int64          `json:"model_id"`
	ManagerContentRouter string         `json:"manager_content_router"`
	AddContentRouter     string         `json:"add_content_router"`
	EditContentRouter    string         `json:"edit_content_router"`
	Description          string         `json:"Description"`
	Thumb                string         `json:"Thumb"`
	Dir                  string         `json:"dir"`
	Url                  string         `json:"Url"`
	Listorder            int64          `json:"Listorder"`
	Ismenu               int64          `json:"Ismenu"`
	IndexTpl             string         `json:"index_tpl"`
	ListTpl              string         `json:"list_tpl"`
	DetailTpl            string         `json:"detail_tpl"`
	UrlPrefix            string         `xorm:"-" json:"url_prefix"`
	Sons                 []Category     `xorm:"-"`
	HasSon               bool           `xorm:"-"`
	Model                *DocumentModel `xorm:"-" json:"model"`
}
