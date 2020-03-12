package tables

import "time"

type Member struct {
	Id           int64     `xorm:"pk" json:"id"`
	Account      string    `json:"account"`
	Password     string    `json:"-"`
	Avatar       string    `json:"avatar"`
	Nickname     string    `json:"nickname"`
	Integral     int64     `json:"integral"`
	SaleIntegral int64     `json:"sale_integral"`
	DrawAccount  string    `json:"draw_account"`
	Telphone     string    `json:"telphone"`
	Qq           string    `json:"qq"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"-"`
	Email        string    `json:"email"`
	Enabled      int64     `json:"enabled"`
	VerifyToken  string    `json:"-"`
}
