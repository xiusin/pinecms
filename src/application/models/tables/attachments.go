package tables

type AttachmentType struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt LocalTime `json:"created_at" xorm:"created"`
}
