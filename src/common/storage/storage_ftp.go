package storage

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/xiusin/pine"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
)

type FtpUploader struct {
	sync.Mutex

	host    string
	fixDir  string
	baseDir string
	client  *ftp.ServerConn
}

func NewFtpUploader(opt map[string]string) *FtpUploader {
	c, err := ftp.Dial(opt["FTP_SERVER_URL"]+":"+opt["FTP_SERVER_PORT"],
		ftp.DialWithTimeout(5*time.Second),
		ftp.DialWithDisabledUTF8(true),
		//ftp.DialWithDebugOutput(os.Stdout),
		ftp.DialWithDisabledEPSV(true),
	)
	if err != nil {
		panic(err)
	}
	if err := c.Login(opt["FTP_USER_NAME"], opt["FTP_USER_PWD"]); err != nil {
		panic(err)
	}
	pine.RegisterOnInterrupt(func() {
		c.Logout()
		c.Quit()
	})
	return &FtpUploader{host: opt["SITE_URL"], fixDir: opt["UPLOAD_URL_PREFIX"], baseDir: opt["FTP_UPLOAD_DIR"], client: c}
}

func (s *FtpUploader) GetEngineName() string {
	return "FTP存储"
}

func (s *FtpUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	s.Lock()
	defer s.Unlock()
	if err := s.client.Stor(storageName, LocalFile); err != nil {
		return "", err
	}
	return storageName, nil
}

func (s *FtpUploader) List(dir string) ([]File, error) {
	s.Lock()
	entities, err := s.client.List(dir)
	s.Unlock()

	if err != nil {
		return nil, err
	}
	fileArr := []File{}
	for _, entity := range entities {
		if entity.Name == "." || entity.Name == ".." {
			continue
		}
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
	s.Lock()
	_, err := s.client.FileSize(name)
	s.Unlock()
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
	s.Lock()
	defer s.Unlock()
	return s.client.Delete(name)
}

func (s *FtpUploader) Mkdir(dir string) error {
	s.Lock()
	defer s.Unlock()

	dirs := strings.Split(strings.Trim(dir, "/"), "/")
	section := make([]string, len(dirs))
	for _, dir := range dirs {
		section = append(section, dir)
		if err := s.client.MakeDir(strings.Join(section, "/")); err != nil {
			return err
		}
	}
	return nil
}

func (s *FtpUploader) Rmdir(dir string) error {
	s.Lock()
	defer s.Unlock()

	return s.client.RemoveDirRecur(dir)
}

func (s *FtpUploader) Content(name string) ([]byte, error) {
	s.Lock()
	defer s.Unlock()
	resp, err := s.client.Retr(name)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	return ioutil.ReadAll(resp)
}

func (s *FtpUploader) Rename(oldname, newname string) error {
	s.Lock()
	defer s.Unlock()

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
