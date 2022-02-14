package storage

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
)

type FileUploader struct {
	host    string
	fixDir  string
	baseDir string
}

func NewFileUploader(opt map[string]string) *FileUploader {
	return &FileUploader{
		host:    opt["SITE_URL"],
		fixDir:  opt["UPLOAD_URL_PREFIX"],
		baseDir: opt["UPLOAD_DIR"],
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
	return s.GetFullUrl(storageName), nil
}

func (s *FileUploader) List(dir string) ([]File, error) {
	scanDir := filepath.Join(s.baseDir, dir)
	fs, err := ioutil.ReadDir(scanDir)
	if err != nil {
		return nil, err
	}
	var list = []File{}
	for _, f := range fs {
		list = append(list, File{
			Id:       f.Name(),
			FullPath: strings.ReplaceAll(filepath.Join(s.fixDir, dir, f.Name()), "\\", "/"),
			Name:     filepath.Base(f.Name()),
			Size:     f.Size(),
			Ctime:    f.ModTime(),
			IsDir:    f.IsDir(),
		})
	}
	return list, nil
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
	return strings.TrimRight(s.host, "/") + getAvailableUrl(filepath.Join(s.fixDir, name))
}

func (s *FileUploader) Remove(name string) error {
	return os.Remove(filepath.Join(s.baseDir, name))
}

func (s *FileUploader) Mkdir(dir string) error {
	return os.Mkdir(filepath.Join(s.baseDir, dir), os.ModePerm)
}

func (s *FileUploader) Rmdir(dir string) error {
	return os.RemoveAll(filepath.Join(s.baseDir, dir))
}

func (s *FileUploader) Content(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.baseDir, name))
}

func (s *FileUploader) Rename(oldname, newname string) error {
	return os.Rename(filepath.Join(s.baseDir, oldname), filepath.Join(s.baseDir, newname))
}

func init() {
	di.Set(fmt.Sprintf(controllers.ServiceUploaderEngine, (&FileUploader{}).GetEngineName()), func(builder di.AbstractBuilder) (interface{}, error) {
		cfg, err := config.SiteConfig()
		if err != nil {
			return nil, err
		}
		checkIsValidConf(cfg)
		return NewFileUploader(cfg), nil
	}, false)
}
