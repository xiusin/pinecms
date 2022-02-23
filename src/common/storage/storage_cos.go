package storage

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
)

type CosUploader struct {
	*cos.Client
	host      string
	urlPrefix string
}

func NewCosUploader(config map[string]string) *CosUploader {

	u, _ := url.Parse(config["COS_BASE_HOST"]) // "https://pinecms-125126195911124.cos.ap-beijing.myqcloud.com"
	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{Transport: &cos.AuthorizationTransport{SecretID: config["COS_SECRET_ID"], SecretKey: config["COS_SECRET_KEY"]}})

	return &CosUploader{
		Client:    client,
		host:      config["OSS_HOST"],
		urlPrefix: strings.Trim(config["UPLOAD_URL_PREFIX"], "/"),
	}
}

func (s *CosUploader) Remove(name string) error {
	name = s.getObjectName(name)
	if exist, err := s.Object.IsExist(context.Background(), name); err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("object %s not exists", name)
	}
	_, err := s.Object.Delete(context.Background(), name)
	return err
}

func (s *CosUploader) GetFullUrl(name string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(s.host, "/"), s.getObjectName(name))
}

func (s *CosUploader) Exists(name string) (bool, error) {
	return s.Object.IsExist(context.Background(), s.getObjectName(name))
}

// Upload 上传图片
// storageName 云端路径名. 最终上传相对urlPrefix生成地址
// LocalFile 要上传的文件名
func (s *CosUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	storageName = getAvailableUrl(s.urlPrefix + "/" + storageName)
	//contentType := mime.TypeByExtension(filepath.Ext(storageName))
	if _, err := s.Object.Put(context.Background(), storageName, LocalFile, nil); err != nil { //上传图片对象
		return "", err
	}
	return s.host + "/" + storageName, nil
}

func (s *CosUploader) List(dir string) ([]File, error) {

	list, _, err := s.Bucket.Get(context.Background(), &cos.BucketGetOptions{
		Prefix:    strings.Trim(s.getObjectName(dir), "/") + "/",
		Delimiter: "/",
		MaxKeys:   200,
	})

	if err != nil {
		return nil, err
	}
	var files = []File{}
	for _, prefix := range list.CommonPrefixes {
		files = append(files, File{Id: prefix, FullPath: "/" + prefix, Name: filepath.Base(prefix), IsDir: true})
	}

	for _, object := range list.Contents {
		t, _ := time.Parse("2006-01-02 15:04:05", object.LastModified)

		files = append(files, File{
			Id:       object.Key,
			FullPath: "/" + object.Key,
			Name:     filepath.Base(object.Key),
			Size:     object.Size,
			Ctime:    t,
			IsDir:    strings.HasSuffix(object.Key, "/"),
		})
	}

	return files, nil
}

func (s *CosUploader) GetEngineName() string {
	return "oss存储"
}

func (s *CosUploader) Content(name string) ([]byte, error) {
	f, err := s.Object.Get(context.Background(), s.getObjectName(name), nil)
	if err != nil {
		return nil, err
	}
	defer f.Body.Close()

	return ioutil.ReadAll(f.Body)
}

func (s *CosUploader) Rename(oldname, newname string) error {
	_, _, err := s.Object.Copy(context.Background(), s.getObjectName(newname), s.getObjectName(oldname), nil)
	if err == nil {
		s.Object.Delete(context.Background(), s.getObjectName(oldname))
	}
	return err
}

func (s *CosUploader) Mkdir(dir string) error {
	return nil
}

func (s *CosUploader) Rmdir(dir string) error {
	return nil
}

func (s *CosUploader) getObjectName(name string) string {
	return s.urlPrefix + "/" + name
}

func init() {
	di.Set(fmt.Sprintf(controllers.ServiceUploaderEngine, (&CosUploader{}).GetEngineName()), func(builder di.AbstractBuilder) (interface{}, error) {
		var engine Uploader
		var err error
		defer func() {
			if errPanic := recover(); errPanic != nil {
				engine = nil
				err = fmt.Errorf("%s", err)
			}
		}()
		cfg, err := config.SiteConfig()
		if err != nil {
			return nil, err
		}
		engine = NewCosUploader(cfg)
		return engine, nil
	}, false)
}
