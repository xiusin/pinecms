package tables

type Link struct {
	Id        int64     `xorm:"int(11) autoincr not null pk 'id'" json:"id" schema:"id"`
	Linktype  int       `xorm:"tinyint(3) not null 'linktype'" json:"linktype" schema:"linktype"`
	Name      string    `xorm:"varchar(50) not null 'name'" json:"name" schema:"name"`
	Url       string    `xorm:"varchar(255) not null 'url'" json:"url" schema:"url"`
	Logo      string    `xorm:"varchar(100) not null 'logo'" json:"logo" schema:"logo"`
	Introduce string    `xorm:"varchar(255) not null 'introduce'" json:"introduce" schema:"introduce"`
	Listorder int64     `xorm:"int(11) not null 'listorder'" json:"listorder" schema:"listorder"`
	Passed    int       `xorm:"tinyint(1) not null default '0' 'passed'" json:"passed" schema:"passed"`
	Addtime   LocalTime `xorm:"datetime default 'null' 'addtime'" json:"addtime" schema:"addtime"`
}
