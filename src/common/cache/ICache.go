package cache

type ICache interface {
	Get(string) string
	Set(string, []byte) error
	IsExist(string) bool
	Delete(string) error
	Flush() error
}
