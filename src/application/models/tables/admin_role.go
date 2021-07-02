package tables

type AdminRole struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	Rolename    string `json:"rolename"`
	Description string `json:"description"`
	Listorder   int64  `json:"listorder"`
	Disabled    int64  `json:"disabled"`

	MenuIdList []int64 `json:"menuIdList" xorm:"json menu_ids"`
}
