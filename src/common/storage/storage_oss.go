package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

var _ Uploader = (*OssUploader)(nil)

func NewOssUploader(config map[string]string) *OssUploader {

	// TODO OSS_ENDPOINT 和 OSS_HOST 和 OSS_BUCKET 是否可以同cos一样利用parseurl获取

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
		urlPrefix: strings.Trim(config["UPLOAD_URL_PREFIX"], "/"),
	}
}

// Upload 上传图片
// storageName 云端路径名. 最终上传相对urlPrefix生成地址
// LocalFile 要上传的文件名
func (s *OssUploader) Upload(storageName string, LocalFile io.Reader) (string, error) {
	if s.client == nil {
		return "", errors.New("ossClient is error")
	}
	storageName = getAvailableUrl(s.urlPrefix + "/" + storageName)
	contentType := mime.TypeByExtension(filepath.Ext(storageName))
	if err := s.bucket.PutObject(storageName, LocalFile, oss.ContentType(contentType)); err != nil { //上传图片对象
		pine.Logger().Error("upoadFile failed", storageName, LocalFile == nil)
		return "", err
	}
	return s.host + "/" + storageName, nil
}

func (s *OssUploader) List(dir string) ([]File, error) {
	list, err := s.bucket.ListObjectsV2(
		oss.Prefix(strings.Trim(s.getObjectName(dir), "/")+"/"),
		oss.Delimiter("/"),
		oss.MaxKeys(200))

	if err != nil {
		return nil, err
	}
	var files = []File{}
	for _, prefix := range list.CommonPrefixes {
		files = append(files, File{Id: prefix, FullPath: "/" + prefix, Name: filepath.Base(prefix), IsDir: true})
	}

	for _, object := range list.Objects {
		files = append(files, File{
			Id:       object.Key,
			FullPath: "/" + object.Key,
			Name:     filepath.Base(object.Key),
			Size:     object.Size,
			Ctime:    object.LastModified,
			IsDir:    strings.HasSuffix(object.Key, "/"),
		})
	}

	return files, nil
}

func (s *OssUploader) GetEngineName() string {
	return "oss存储"
}

func (s *OssUploader) Content(name string) ([]byte, error) {
	f, err := s.bucket.GetObject(s.getObjectName(name))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (s *OssUploader) Rename(oldname, newname string) error {
	_, err := s.bucket.CopyObject(s.getObjectName(oldname), s.getObjectName(newname))
	if err == nil {
		s.bucket.DeleteObject(s.getObjectName(oldname))
	}
	return err
}

func (s *OssUploader) Mkdir(dir string) error {
	var byteBuffer bytes.Buffer

	return s.bucket.PutObject(s.getObjectName(dir) + "/", &byteBuffer, nil)
}

func (s *OssUploader) Rmdir(dir string) error {
	return s.bucket.DeleteObject(s.getObjectName(dir) + "/")
}

func (s *OssUploader) getObjectName(name string) string {
	return s.urlPrefix + "/" + name
}

func (s *OssUploader) Remove(name string) error {
	name = s.getObjectName(name)
	if exist, err := s.bucket.IsObjectExist(name); err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("object %s not exists", name)
	}
	return s.bucket.DeleteObject(name)
}

func (s *OssUploader) GetFullUrl(name string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(s.host, "/"), s.getObjectName(name))
}

func (s *OssUploader) Exists(name string) (bool, error) {
	return s.bucket.IsObjectExist(s.getObjectName(name))
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
