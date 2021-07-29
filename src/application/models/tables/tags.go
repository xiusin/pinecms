package tables

type Tags struct {
	Id        int64     `xorm:"pk autoincr id" json:"id"`
	Name      string    `json:"name"`
	RefNum    uint      `json:"ref_num"`
	Listorder uint      `json:"listorder"`
	Status    uint      `json:"status"`
	CreatedAt LocalTime `json:"created_at"`
}
