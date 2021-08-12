package tables

type District struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
	Initial  string `json:"initial"`
	Initials string `json:"initials"`
	PinYin   string `json:"pinyin" xorm:"pinyin"`
	Extra    string `json:"extra"`
	Suffix   string `json:"suffix"`
	Code     string `json:"code"`
	AreaCode string `json:"area_code" xorm:"area_code"`
	OrderNum uint   `json:"order" xorm:"order"`
}
