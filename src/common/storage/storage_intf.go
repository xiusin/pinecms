package storage

import (
	"io"
	"runtime"
	"strings"
)

type Uploader interface {
	Upload(storageName string, LocalFile io.Reader) (string, error)
	List(dir string) (list []string, prefix string, err error)
	Exists(name string) (bool, error)
	GetFullUrl(name string) string
	Remove(name string) error
	GetEngineName() string
}

func getAvailableUrl(path string) string {
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}
