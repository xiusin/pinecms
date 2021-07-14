package tables

//Position 职位管理
type Position struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" schema:"name" xorm:"varchar(25)"`
	Code      string `json:"code" xorm:"varchar(25)"`
	Listorder uint   `json:"listorder"`
	Status    bool   `json:"status"`
	Remark    string `json:"remark" xorm:"text"`
}
