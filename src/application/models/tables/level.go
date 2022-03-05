package tables

// Level 职级
type Level struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" schema:"name" xorm:"varchar(25)"`
	Listorder uint   `json:"listorder"`
	Status    bool   `json:"status"`
	CreatedAt   LocalTime  `json:"created" xorm:"created"`
	UpdatedAt   *LocalTime `json:"updated" xorm:"updated"`
}
