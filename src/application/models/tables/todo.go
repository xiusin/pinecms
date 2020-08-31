package tables

type Todo struct {
	Id int64 `xorm:"int autoincr not null pk 'id'" json:"id" schema:"id"`
	Userid int64 `xorm:"int not null 'userid'" json:"userid" schema:"userid" validate:"required"`
	Message string `xorm:"text not null 'message'" json:"message" schema:"message" validate:"required"`
	Status int64 `xorm:"int not null 'status'" json:"status" schema:"status" validate:"required"`
	SetStatus string `xorm:"set('1','2','3') default '1' 'set_status'" json:"set_status" schema:"set_status" validate:"required"`
	EnumStatus int `xorm:"enum('0','1','2') default '0' 'enum_status'" json:"enum_status" schema:"enum_status" validate:"required"`
	Image string `xorm:"varchar(255) default 'null' 'image'" json:"image" schema:"image" validate:"required"`
	Images string `xorm:"varchar(255) default 'null' 'images'" json:"images" schema:"images" validate:"required"`
	File string `xorm:"varchar(255) default 'null' 'file'" json:"file" schema:"file" validate:"required"`
	Files string `xorm:"varchar(255) default 'null' 'files'" json:"files" schema:"files" validate:"required"`
	CityId int64 `xorm:"int default 'null' 'city_id'" json:"city_id" schema:"city_id" validate:"required"`
	Content string `xorm:"text 'content'" json:"content" schema:"content" validate:"required"`
}
