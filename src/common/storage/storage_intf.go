package storage

import (
	"io"
	"runtime"
	"strings"
	"time"
)

type Uploader interface {
	Upload(storageName string, LocalFile io.Reader) (string, error)
	List(dir string) (list []File, prefix string, err error)
	Exists(name string) (bool, error)
	GetFullUrl(name string) string
	Remove(name string) error
	GetEngineName() string
}

type File struct {
	Id       string    `json:"id"`
	FullPath string    `json:"full_path"`
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	Ctime    time.Time `json:"ctime"`
}

func getAvailableUrl(path string) string {
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}
