package filemanager

import (
	"sync"

	"github.com/xiusin/pinecms/src/config"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend/filemanager/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/common/storage"
)

type ResResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var once sync.Once

const Logined = "true"

func InitInstall(app *pine.Application, urlPrefix, dir string) {
	once.Do(func() {
		app.Static(urlPrefix, dir, 1)

		orm := helper.GetORM()
		defer func() {
			if err := recover(); err != nil {
				pine.Logger().Warning("初始化安装失败", err)
			}
		}()

		if err := orm.Sync2(&tables.FileManagerAccount{}); err != nil {
			pine.Logger().Warning(err)
		}
		if count, _ := orm.Count(&tables.FileManagerAccount{}); count == 0 {
			user := &tables.FileManagerAccount{Username: "admin", Nickname: "Administer", Engine: "本地存储"}
			user.Init()
			user.Password = user.GetMd5Pwd("admin888")
			if _, err := orm.InsertOne(user); err != nil {
				pine.Logger().Warning("新增用户失败", err)
			}
		}
		pine.Logger().Debug("初始化安装成功")
	})
}

func ResponseError(c *pine.Context, msg string) {
	c.Render().JSON(pine.H{
		"result": ResResult{
			Status:  "danger",
			Message: msg,
		},
	})
}

type EngineFn func(map[string]string) storage.Uploader

type Engine struct {
	Name   string
	Engine EngineFn
}

func EngineList() []Engine {
	return []Engine{
		{"本地存储", func(opt map[string]string) storage.Uploader {
			cnf, _ := config.SiteConfig()
			return storage.NewFileUploader(cnf)
		}},
		{"Oss对象存储", func(opt map[string]string) storage.Uploader {
			return storage.NewOssUploader(opt)
		}},
	}
}

func GetUserUploader(u *tables.FileManagerAccount) storage.Uploader {
	if u == nil {
		return nil
	}
	for _, v := range EngineList() {
		if v.Name == u.Engine {
			cnf, _ := config.SiteConfig()
			return v.Engine(cnf)
		}
	}
	return nil
}

type FMFileProps struct {
	HasSubdirectories    bool `json:"hasSubdirectories"`
	SubdirectoriesLoaded bool `json:"subdirectoriesLoaded"`
	ShowSubdirectories   bool `json:"showSubdirectories"`
}

type FMFile struct {
	ID        interface{} `json:"id"`
	Basename  string      `json:"basename"`
	Filename  string      `json:"filename"`
	Dirname   string      `json:"dirname"`
	Path      string      `json:"path"`
	ParentID  string      `json:"parentId"`
	Timestamp int64       `json:"timestamp"`
	ACL       int         `json:"acl"`
	Size      int         `json:"size"`
	Type      string      `json:"type"`
	Extension string      `json:"extension"`
	Props     FMFileProps `json:"props"`
	Author    string      `json:"author"`
}

type DelItem struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

