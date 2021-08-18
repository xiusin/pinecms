package storage

import (
	"fmt"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileUploader struct {
	fixDir  string
	baseDir string
}

func NewFileUploader(fixDir, uploadDir string) *FileUploader {
	return &FileUploader{
		fixDir:  fixDir,
		baseDir: uploadDir,
	}
}

func (s *FileUploader) GetEngineName() string {
	return "本地存储"
}

func (s *FileUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	saveFile := filepath.Join(s.baseDir, storageName)
	_ = os.MkdirAll(filepath.Dir(saveFile), os.ModePerm)
	out, err := os.OpenFile(saveFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, LocalFile)
	if err != nil {
		return "", err
	}
	return getAvailableUrl(filepath.Join(s.fixDir, storageName)), nil
}

func (s *FileUploader) List(dir string) ([]string, string, error) {
	scanDir := filepath.Join(s.baseDir, dir)
	fs, err := ioutil.ReadDir(scanDir)
	if err != nil {
		return nil, "", err
	}
	var list []string
	for _, f := range fs {
		list = append(list, filepath.Join(s.fixDir, dir, f.Name()))
	}
	return list, s.fixDir, nil
}

func (s *FileUploader) Exists(name string) (bool, error) {
	_, err := os.Stat(filepath.Join(s.baseDir, name))
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}

func (s *FileUploader) GetFullUrl(name string) string {
	return getAvailableUrl(filepath.Join(s.fixDir, name))
}

func (s *FileUploader) Remove(name string) error {
	return os.Remove(filepath.Join(s.baseDir, name))
}

func init()  {
	di.Set(fmt.Sprintf(controllers.ServiceUploaderEngine, (&FileUploader{}).GetEngineName()), func(builder di.AbstractBuilder) (interface{}, error) {
		cfg, err := config.SiteConfig()
		if err != nil {
			return nil, err
		}
		uploadDir := cfg["UPLOAD_DIR"]
		urlPrefixDir := cfg["UPLOAD_URL_PREFIX"]
		return NewFileUploader(urlPrefixDir, uploadDir), nil
	}, false)
}
