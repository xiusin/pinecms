package tables

type District struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
	Initial  string `json:"initial"`
	Initials string `json:"initials"`
	PinYin   string `json:"pinyin"`
	Extra    string `json:"extra"`
	Suffix   string `json:"suffix"`
	Code     string `json:"code"`
	AreaCode string `json:"area_code"`
	OrderNum uint   `json:"order"`
}
