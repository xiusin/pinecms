package cache

type ICache interface {
	SetTable(string)
	GetTable() string
	Get(string) string
	Set(string, []byte) error
	IsExist(string) bool
	Delete(string) error
	Flush() error
}
