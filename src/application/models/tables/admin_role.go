package tables

type AdminRole struct {
	Roleid      int64  `xorm:"pk autoincr" json:"roleid"`
	Rolename    string `json:"rolename"`
	Description string `json:"description"`
	Listorder   int64  `json:"listorder"`
	Disabled    int64  `json:"disabled"`
}
