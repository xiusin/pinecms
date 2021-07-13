package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
	"path/filepath"
	"strings"
)

type DatabaseBackupController struct {
	BaseController
}

// 定时备份任务功能

func (c *DatabaseBackupController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/backup/list", "BackupList")
	b.POST("/backup/delete", "BackupDelete")
	b.POST("/backup/download", "BackupDownload")
}

func (c *DatabaseBackupController) BackupList() {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	uploader := getStorageEngine(settingData)
	list, prefix, err := uploader.List(baseBackupDir)
	var files []map[string]string
	if err != nil {
		c.Logger().Error(err)
	}
	for _, v := range list {
		v = strings.TrimLeft(v, filepath.Join(prefix, baseBackupDir))
		files = append(files, map[string]string{"name": v})
	}
	helper.Ajax(files, 0, c.Ctx())
}

func (c *DatabaseBackupController) BackupDownload() {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	name := c.Ctx().FormValue("name")
	if name == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	relName := filepath.Join(baseBackupDir, name)
	// 判断文件是否存在
	uploader := getStorageEngine(settingData)
	exists, err := uploader.Exists(relName)
	if err != nil {
		helper.Ajax("获取文件信息错误: "+err.Error(), 1, c.Ctx())
		return
	}
	if !exists {
		helper.Ajax("文件不存在或已经被删除", 1, c.Ctx())
		return
	}
	// 返回下载地址
	helper.Ajax(uploader.GetFullUrl(relName), 0, c.Ctx())
}

func (c *DatabaseBackupController) BackupDelete(orm *xorm.Engine) {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)

	name := c.Ctx().FormValue("name")
	if name == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}

	relName := filepath.Join(baseBackupDir, name)
	uploader := getStorageEngine(settingData)
	if err := uploader.Remove(relName); err != nil {
		helper.Ajax("删除文件错误: "+err.Error(), 1, c.Ctx())
		return
	}
	helper.Ajax("删除文件成功", 0, c.Ctx())
}
