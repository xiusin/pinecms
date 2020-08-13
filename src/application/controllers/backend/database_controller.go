package backend

import (
	"bytes"
	"fmt"
	"github.com/xiusin/pinecms/src/config"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexmullins/zip"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
)

type DatabaseController struct {
	pine.Controller
}

var baseBackupDir = fmt.Sprintf("%s/%s", "database", "backup")

// 定时备份任务功能

func (c *DatabaseController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/database/manager", "Manager")
	if config.DBConfig().Db.DbDriver == "sqlite3" {
		b.POST("/database/repair", "NoSupport")
		b.POST("/database/optimize", "NoSupport")
		b.POST("/database/backup", "NoSupport")
	} else {
		b.POST("/database/repair", "Repair")
		b.POST("/database/optimize", "Optimize")
		b.POST("/database/backup", "Backup")
	}
	b.ANY("/database/backup-list", "BackupList")
	b.POST("/database/backup-delete", "BackupDelete")
	b.POST("/database/backup-download", "BackupDownload")
}

func (c *DatabaseController) NoSupport() {
	helper.Ajax("SQLITE不支持此功能", 1, c.Ctx())
}

func (c *DatabaseController) Manager(orm *xorm.Engine) {
	mataDatas, err := orm.DBMetas()
	var data = []map[string]interface{}{}
	if err != nil {
		pine.Logger().Error("读取数据库元信息失败", err)
	} else {
		for _, mataData := range mataDatas {
			total, _ := orm.Table(mataData.Name).Count()
			data = append(data, map[string]interface{}{
				"id":      mataData.Name,
				"total":   total,
				"engine":  mataData.StoreEngine,
				"comment": mataData.Comment,
			})
		}
	}
	helper.Ajax(data, 0, c.Ctx())
}

func (c *DatabaseController) Repair(orm *xorm.Engine) {
	tables := strings.Split(c.Ctx().FormValue("ids"), ",")
	if len(tables) == 0 {
		helper.Ajax("请选择要修复的表", 1, c.Ctx())
		return
	}
	for _, table := range tables {
		_, err := orm.Exec("REPAIR TABLE `" + table + "`")
		if err != nil {
			helper.Ajax("表修复错误："+table+": "+err.Error(), 1, c.Ctx())
			return
		}
	}
	helper.Ajax("表 "+c.Ctx().FormValue("tables")+" 修复成功", 0, c.Ctx())
}

func (c *DatabaseController) Optimize(orm *xorm.Engine) {
	tables := strings.Split(c.Ctx().FormValue("ids"), ",")
	if len(tables) == 0 {
		helper.Ajax("请选择要优化的表", 1, c.Ctx())
		return
	}

	for _, table := range tables {
		_, err := orm.Exec("OPTIMIZE TABLE `" + table + "`")
		if err != nil {
			helper.Ajax("表优化错误："+table+": "+err.Error(), 1, c.Ctx())
			return
		}
	}
	helper.Ajax("表 "+c.Ctx().FormValue("tables")+" 优化成功", 0, c.Ctx())
}

func (c *DatabaseController) Backup() {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	msg, code := c.backup(settingData, false)
	helper.Ajax(msg, int64(code), c.Ctx())
}

func (c *DatabaseController) BackupList() {
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	uploader := getStorageEngine(settingData)
	list, prefix, err := uploader.List(baseBackupDir)
	var files = []map[string]string{}
	if err != nil {
		c.Logger().Error(err)
	}
	for _, v := range list {
		v = strings.TrimLeft(v, filepath.Join(prefix, baseBackupDir))
		files = append(files, map[string]string{
			"name": v,
		})
	}
	helper.Ajax(files, 0, c.Ctx())
}

func (c *DatabaseController) BackupDownload() {
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

func (c *DatabaseController) BackupDelete(orm *xorm.Engine) {
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

func (c *DatabaseController) backup(settingData map[string]string, auto bool) (msg string, errcode int) {
	orm := pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	if settingData["UPLOAD_DATABASE_PASS"] == "" {
		return "请先设置备份数据库打包zip的密码", 1
	}
	fNameBaseName := strings.Replace(
		strings.Replace(time.Now().In(helper.GetLocation()).Format(helper.TimeFormat), " ", "-", 1),
		":", "", 3)
	uploader := getStorageEngine(settingData)
	uploadFile := fmt.Sprintf("%s/%s", baseBackupDir, fNameBaseName+".zip")
	buf := bytes.NewBuffer([]byte{})
	if err := orm.DumpAll(buf); err != nil {
		pine.Logger().Error("备份数据表失败", err)
		return "备份表数据失败: " + err.Error(), 1
	}
	zipsc := bytes.NewBuffer([]byte{})
	zipw := zip.NewWriter(zipsc)
	w, err := zipw.Encrypt(fNameBaseName+".sql", settingData["UPLOAD_DATABASE_PASS"])
	if err != nil {
		return "打包zip失败: " + err.Error(), 1
	}
	_, err = io.Copy(w, buf)
	if err != nil {
		zipw.Close()
		return "打包zip失败: " + err.Error(), 1
	}
	zipw.Flush()
	zipw.Close()
	f, err := uploader.Upload(uploadFile, zipsc)
	if err != nil {
		return "备份表数据失败: " + err.Error(), 1
	}
	return "备份数据库成功: " + f, 0
}
