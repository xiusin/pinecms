package tables

type CategoryPriv struct {
	Catid   int64 `xorm:"pk"`
	Roleid  int64
	IsAdmin int64
	Action  string
}
