package tables

type Todo struct {
	Id int64 `xorm:"int autoincr not null pk 'id'" json:"id" schema:"id"`
	Userid int64 `xorm:"int not null 'userid'" json:"userid" schema:"userid"`
	Message string `xorm:"longtext not null 'message'" json:"message" schema:"message" validate:"required"`
	Status int64 `xorm:"int not null 'status'" json:"status" schema:"status" validate:"required,gte=1"`
}
