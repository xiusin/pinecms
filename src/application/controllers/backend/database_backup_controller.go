package backend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
	"os"
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
	var files = []map[string]interface{}{}
	if err != nil {
		c.Logger().Error(err)
	}
	for _, v := range list {
		finfo, err := os.Stat(filepath.Join(settingData["UPLOAD_DIR"], baseBackupDir, strings.TrimLeft(v, filepath.Join(prefix, baseBackupDir))))
		if err != nil {
			helper.Ajax(err, 1, c.Ctx())
			return
		}
		files = append(files, map[string]interface{}{
			"id":     finfo.Name(),
			"name":   finfo.Name(),
			"size":   finfo.Size(),
			"ctime":  finfo.ModTime(),
		})
	}
	helper.Ajax(files, 0, c.Ctx())
}

func (c *DatabaseBackupController) BackupDownload() {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	name := strings.Trim(c.Input().Get("name").String(), `"`)
	if len(name) == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	relName := filepath.Join(baseBackupDir, name)
	uploader := getStorageEngine(settingData)
	exists, _ := uploader.Exists(relName)
	if !exists {
		helper.Ajax("文件不存在或已经被删除", 1, c.Ctx())
		return
	}
	helper.Ajax(uploader.GetFullUrl(relName), 0, c.Ctx())
}

func (c *DatabaseBackupController) BackupDelete() {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	names := c.Input().GetArray("ids")
	if len(names) == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	relName := filepath.Join(baseBackupDir, strings.Trim(names[0].String(), `"`))
	uploader := getStorageEngine(settingData)
	if err := uploader.Remove(relName); err != nil {
		helper.Ajax("删除文件错误: "+err.Error(), 1, c.Ctx())
		return
	}
	helper.Ajax("删除文件成功", 0, c.Ctx())
}
