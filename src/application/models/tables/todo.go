package tables

type Todo struct {
	Id          int64     `xorm:"int(11) autoincr not null pk 'id'" json:"id"`
	Type        string    `xorm:"varchar(30) not null default '1' 'type' comment('字符串多选:1=外部链接,2=内部链接,3=通用链接')" json:"type" validate:"required"`
	Name        string    `xorm:"varchar(50) not null 'name' comment('普通输入框')" json:"name" validate:"required"`
	Introduce   string    `xorm:"varchar(255) not null 'introduce' comment('普通多行输入框::cms-textarea')" json:"introduce" validate:"required"`
	Listorder   int64     `xorm:"int(11) not null 'listorder' comment('不可为空数字')" json:"listorder" validate:"required"`
	Status      int       `xorm:"tinyint(1) not null default 0 'status' comment('tinyint单选:0=待审核,1=通过,2=拒绝:cms-radio')" json:"status" validate:"required"`
	PutDate     LocalTime `xorm:"date default null 'put_date' comment('日期')" json:"put_date" validate:"required"`
	PutDatetime LocalTime `xorm:"datetime default null 'put_datetime' comment('时间日期')" json:"put_datetime" validate:"required"`
	StartTime   LocalTime `xorm:"datetime default null 'start_time' comment('开始时间$end=end_time')" json:"start_time" validate:"required"`
	EndTime     LocalTime `xorm:"datetime default null 'end_time' comment('结束时间被引用隐藏到代码区间选择器')" json:"end_time" validate:"required"`
	Logo        string    `xorm:"varchar(30) default null 'logo' comment('单图上传')" json:"logo" validate:"required"`
	Logos       string    `xorm:"varchar(255) default null 'logos' comment('多图上传')" json:"logos" validate:"required"`
}
