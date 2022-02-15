package storage

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/xiusin/pine"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
)

type FtpUploader struct {
	host    string
	fixDir  string
	baseDir string
	client  *ftp.ServerConn
}

func NewFtpUploader(opt map[string]string) *FtpUploader {
	c, err := ftp.Dial(opt["FTP_SERVER_URL"]+":"+opt["FTP_SERVER_PORT"], ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	if err := c.Login(opt["FTP_USER_NAME"], opt["FTP_USER_PWD"]); err != nil {
		panic(err)
	}
	pine.RegisterOnInterrupt(func() {
		c.Quit()
	})
	return &FtpUploader{host: opt["SITE_URL"], fixDir: opt["UPLOAD_URL_PREFIX"], baseDir: opt["UPLOAD_DIR"], client: c}
}

func (s *FtpUploader) GetEngineName() string {
	return "FTP存储"
}

func (s *FtpUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	if err := s.client.Stor(storageName, LocalFile); err != nil {
		return "", err
	}
	return storageName, nil
}

func (s *FtpUploader) List(dir string) ([]File, error) {
	entities, err := s.client.List(dir)
	if err != nil {
		return nil, err
	}
	fileArr := []File{}
	for _, entity := range entities {
		f := File{
			Id:       filepath.Join(dir, entity.Name),
			FullPath: filepath.Join(dir, entity.Name),
			Name:     entity.Name,
			Size:     int64(entity.Size),
			Ctime:    entity.Time,
		}
		if entity.Type.String() == "folder" {
			f.IsDir = true
		}
		fileArr = append(fileArr, f)
	}
	return fileArr, nil
}

func (s *FtpUploader) Exists(name string) (bool, error) {
	_, err := s.client.FileSize(name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *FtpUploader) GetFullUrl(name string) string {
	// 链接地址需要本地提供服务, 先下载到本地再发送给客户端预览

	return strings.TrimRight(s.host, "/") + getAvailableUrl(filepath.Join(s.fixDir, name))
}

func (s *FtpUploader) Remove(name string) error {
	return s.client.Delete(name)
}

func (s *FtpUploader) Mkdir(dir string) error {
	return s.client.MakeDir(dir) // 创建目录
}

func (s *FtpUploader) Rmdir(dir string) error {
	return s.client.RemoveDirRecur(dir)
}

func (s *FtpUploader) Content(name string) ([]byte, error) {
	resp, err := s.client.Retr(name)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	return ioutil.ReadAll(resp)
}

func (s *FtpUploader) Rename(oldname, newname string) error {
	return s.client.Rename(oldname, newname)
}

func init() {
	di.Set(fmt.Sprintf(controllers.ServiceUploaderEngine, (&FtpUploader{}).GetEngineName()), func(builder di.AbstractBuilder) (interface{}, error) {
		cfg, err := config.SiteConfig()
		if err != nil {
			return nil, err
		}
		return NewFtpUploader(cfg), nil
	}, false)
}
