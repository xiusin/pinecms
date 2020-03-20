package storage

import "io"

type Uploader interface {
	Upload(storageName string, LocalFile io.Reader) (string, error)
	List(dir string) (list []string,prefix string, err error)
	Exists(name string) (bool, error)
	GetFullUrl(name string) string
	Remove(name string) error
}
