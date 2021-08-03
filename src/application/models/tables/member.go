package tables

type Member struct {
	Id          int64      `xorm:"pk autoincr" json:"id"`
	Account     string     `json:"account" xorm:"comment('账号') varchar(40)"`
	Password    string     `json:"password" xorm:"comment('密码') varchar(32)"`
	Avatar      string     `json:"avatar" xorm:"comment('头像') varchar(40)"`
	Nickname    string     `json:"nickname" xorm:"comment('昵称') varchar(40)"`
	Integral    uint       `json:"integral" xorm:"comment('积分') int(11)"`
	Telphone    string     `json:"telphone" xorm:"comment('电话') varchar(30)"`
	Qq          string     `json:"qq" xorm:"comment('QQ') varchar(15)"`
	Description string     `json:"description" xorm:"comment('个人简介')"`
	CreatedAt   LocalTime  `json:"created" xorm:"created"`
	UpdatedAt   *LocalTime `json:"updated" xorm:"updated"`
	LoginTime   LocalTime  `json:"login_time" xorm:"datetime comment('最后登录时间')"`
	LoginIp     string     `json:"login_ip" xorm:"varchar(15) comment('最后登录IP')"`
	Email       string     `json:"email" xorm:"comment('邮箱') varchar(30)"`
	Status      uint       `json:"status" xorm:"comment('状态: 0=禁用 1=待验证 2=正常')"`
	Sex         uint       `json:"sex" xorm:"comment('性别: 0=保密 1=男 2=女') tinyint(3)"`
	GroupId     uint       `json:"group_id" xorm:"comment('分组ID') int(6)"`
	VerifyToken string     `json:"-" xorm:"comment('验证token')"`
}

type MemberGroup struct {
	Id          int64      `xorm:"pk autoincr" json:"id"`
	Name        string     `json:"name" xorm:"comment('名称') varchar(40)"`
	Description string     `json:"description" xorm:"comment('介绍')"`
	Status      uint       `json:"status" xorm:"comment('状态: 0=禁用 1=正常')"`
	Listorder   uint       `json:"listorder"`
	CreatedAt   LocalTime  `json:"created" xorm:"created"`
	UpdatedAt   *LocalTime `json:"updated" xorm:"updated"`
}
