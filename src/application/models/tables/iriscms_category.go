package tables

type IriscmsCategory struct {
	Catid                int64             `xorm:"pk" json:"Catid"`
	Parentid             int64             `json:"Parentid"`
	Catname              string            `json:"Catname"`
	Type                 int64             `json:"Type"`
	ModelId              int64             `json:"model_id"`
	ManagerContentRouter string            `json:"manager_content_router"`
	AddContentRouter     string            `json:"add_content_router"`
	EditContentRouter    string            `json:"edit_content_router"`
	Description          string            `json:"Description"`
	Thumb                string            `json:"Thumb"`
	Url                  string            `json:"Url"`
	Listorder            int64             `json:"Listorder"`
	Ismenu               int64             `json:"Ismenu"`
	IndexTpl             string            `json:"index_tpl"`
	ListTpl              string            `json:"list_tpl"`
	DetailTpl            string            `json:"detail_tpl"`
	Sons                 []IriscmsCategory `xorm:"-"`
	HasSon               bool              `xorm:"-"`
}
