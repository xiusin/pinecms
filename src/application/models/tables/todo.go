package tables

type Todo struct {
	Id      int64
	UserID  int64 `xorm:"userid"`
	Message string
	Status  bool
}
