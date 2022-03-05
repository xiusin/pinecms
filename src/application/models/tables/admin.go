package tables

type Admin struct {
	Userid        int64      `xorm:"pk autoincr 'id'" json:"id"`
	Username      string     `json:"username"`
	Password      string     `json:"-"`
	Roleid        int64      `json:"roleid" xorm:"-"`
	Encrypt       string     `json:"-"`
	Lastloginip   string     `json:"lastloginip"`
	Lastlogintime int64      `json:"lastlogintime"`
	Email         string     `json:"email"`
	Realname      string     `json:"realname"`
	Avatar        string     `json:"avatar"`
	Remark        string     `json:"remark"`
	RoleName      string     `json:"roleName" xorm:"-"`
	Phone         string     `json:"phone"`
	Status        uint       `json:"status"`
	PositionId    uint       `json:"position_id"`
	LevelId       uint       `json:"level_id"`
	DepartmentId  uint       `json:"department_id"`
	Birthday      *LocalTime `json:"birthday"`

	RoleIdList []int64 `json:"roleIdList" xorm:"json roles"`
}
