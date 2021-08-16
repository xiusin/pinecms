package tables

type Todo struct {
	Id int64 `xorm:"int(11) autoincr not null pk 'id'" json:"id"`
	Linktype int `xorm:"tinyint(3) not null 'linktype'" json:"linktype" validate:"required"`
	Name string `xorm:"varchar(50) not null 'name'" json:"name" validate:"required"`
	Url string `xorm:"varchar(255) not null 'url'" json:"url" validate:"required"`
	Logo string `xorm:"varchar(100) not null 'logo'" json:"logo" validate:"required"`
	Introduce string `xorm:"varchar(255) not null 'introduce'" json:"introduce" validate:"required"`
	Listorder int64 `xorm:"int(11) not null 'listorder'" json:"listorder" validate:"required"`
	Passed int `xorm:"tinyint(1) not null default '0' 'passed'" json:"passed" validate:"required"`
	Addtime LocalTime `xorm:"datetime default 'null' 'addtime'" json:"addtime" validate:"required"`
}
