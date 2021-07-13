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
}

func getAvailableUrl(path string) string {
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}

type UploadRegister struct {
	Name string
	Inst Uploader
}

var storages []UploadRegister

func register(name string, ins Uploader) {
	storages = append(storages, UploadRegister{
		Name: name,
		Inst: ins,
	})
}
