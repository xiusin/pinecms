package tables

type Department struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" schema:"name" xorm:"varchar(25)"`
	LeaderName  string    `json:"leader_name" xorm:"varchar(25)"`
	LeaderPhone string    `json:"leader_phone" xorm:"varchar(35)"`
	Email       string    `json:"email" xorm:"varchar(100)"`
	Status      bool      `json:"status" schema:"status"` // 状态
	Listorder   uint      `json:"listorder"`
	ParentId    uint      `json:"parent_id"`
	CreatedAt   LocalTime `json:"created_at"`
}
