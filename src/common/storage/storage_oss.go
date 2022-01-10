package storage

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
)

type OssUploader struct {
	client    *oss.Client
	bucket    *oss.Bucket
	host      string
	urlPrefix string
}

func (s *OssUploader) Remove(name string) error {
	return s.bucket.DeleteObject(strings.TrimLeft(filepath.Join(s.urlPrefix, name), "/"))
}

func (s *OssUploader) GetFullUrl(name string) string {
	return fmt.Sprintf("%s/%s",
		strings.TrimRight(s.host, "/"),
		strings.TrimLeft(filepath.Join(s.urlPrefix, name), "/"),
	)
}

func (s *OssUploader) Exists(name string) (bool, error) {
	name = strings.TrimLeft(getAvailableUrl(filepath.Join(s.urlPrefix, name)), "/")
	return s.bucket.IsObjectExist(name)
}

func checkIsValidConf(config map[string]string) {
	if config["OSS_ENDPOINT"] == "" || config["OSS_KEYID"] == "" || config["OSS_KEYSECRET"] == "" || config["OSS_BUCKET"] == "" {
		panic("请配置OSS信息")
	}
}

func NewOssUploader(config map[string]string) *OssUploader {
	checkIsValidConf(config)
	client, err := oss.New(config["OSS_ENDPOINT"], config["OSS_KEYID"], config["OSS_KEYSECRET"])
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(config["OSS_BUCKET"])
	if err != nil {
		return nil
	}
	return &OssUploader{
		client:    client,
		bucket:    bucket,
		host:      config["OSS_HOST"],
		urlPrefix: config["UPLOAD_URL_PREFIX"],
	}
}

// Upload 上传图片
// storageName 云端路径名. 最终上传相对urlPrefix生成地址
// LocalFile 要上传的文件名
func (s *OssUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	storageName = getAvailableUrl(filepath.Join(s.urlPrefix, storageName))
	if s.client == nil {
		return "", errors.New("ossClient is error")
	}
	storageName = strings.TrimLeft(storageName, "/")
	contentType := mime.TypeByExtension(filepath.Ext(storageName))
	if err := s.bucket.PutObject(storageName, LocalFile, oss.ContentType(contentType)); err != nil { //上传图片对象
		pine.Logger().Error("upoadFile failed", storageName, LocalFile == nil)
		return "", err
	}
	return s.host + "/" + storageName, nil
}

func (s *OssUploader) List(dir string) ([]File, string, error) {
	list, err := s.bucket.ListObjects(oss.Prefix(strings.TrimLeft(getAvailableUrl(filepath.Join(s.urlPrefix, dir)), "/")))
	if err != nil {
		return nil, "", err
	}
	var files []File
	for _, object := range list.Objects {
		files = append(files, File{
			Id:       object.Key,
			FullPath: s.host + "/" + object.Key,
			Name:     filepath.Base(object.Key),
			Size:     object.Size,
			Ctime:    object.LastModified,
		})
	}
	return files, s.host, nil
}

func (s *OssUploader) GetEngineName() string {
	return "oss存储"
}

func init() {
	di.Set(fmt.Sprintf(controllers.ServiceUploaderEngine, (&OssUploader{}).GetEngineName()), func(builder di.AbstractBuilder) (interface{}, error) {
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
		engine = NewOssUploader(cfg)
		return engine, nil
	}, false)
}
