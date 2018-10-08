package tables

import "time"

type IriscmsMember struct {
	Id           int64     `xorm:"pk" json:"id"`
	Account      string    `json:"account"`
	Password     string    `json:"-"`
	Avatar       string    `json:"avatar"`
	Nickname     string    `json:"nickname"`
	Integral     int64     `json:"integral"`
	SaleIntegral int64     `json:"sale_integral"`
	DrawAccount  string    `json:"draw_account"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
