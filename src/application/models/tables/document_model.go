package tables

import "time"

type DocumentModel struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Table           string    `json:"table"`
	Enabled         int       `json:"enabled"`
	ModelType       int       `json:"model_type"`
	FeTplIndex      string    `json:"fe_tpl_index"`
	FeTplList       string    `json:"fe_tpl_list"`
	FeTplDetail     string    `json:"fe_tpl_detail"`
	DeletedAt       time.Time `xorm:"deleted_at" json:"deleted_at"`
	FieldShowInList string    `json:"field_show_in_list"`
	FeSearchFields  string    `json:"fe_search_fields"`
	Formatters      string    `json:"formatters"`
	Execed          int       `json:"execed"`
}
