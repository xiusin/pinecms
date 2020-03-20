package storage

import (
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

// Upload 上传图片
// storageName 云端路径名.
// LocalFile 要上传的文件名
func (s *FileUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	//检测是否可以生成目录
	originName := storageName
	storageName = filepath.Join(s.baseDir, storageName)
	uploadDir := filepath.Dir(storageName)
	f, err := os.Open(uploadDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	f.Close()
	out, err := os.OpenFile(storageName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, LocalFile)
	if err != nil {
		return "", err
	}
	return filepath.Join(s.fixDir, originName), nil
}

func (s *FileUploader) List(dir string) ([]string, string, error) {
	scandir := filepath.Join(s.baseDir, dir)
	fs, err := ioutil.ReadDir(scandir)
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
	fullname := filepath.Join(s.baseDir, name)
	_, err := os.Stat(fullname)
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
	return filepath.Join(s.fixDir, name)
}

func (s *FileUploader) Remove(name string) error {
	return os.Remove(filepath.Join(s.baseDir, name))
}
