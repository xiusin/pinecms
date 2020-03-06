package backend

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alexmullins/zip"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
)

type DatabaseController struct {
	pine.Controller
}

var baseBackupDir = fmt.Sprintf("%s/%s", "database", "backup")

// 定时备份任务功能

func (c *DatabaseController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/database/manager", "Manager")
	b.POST("/database/repair", "Repair")
	b.POST("/database/optimize", "Optimize")
	b.POST("/database/backup", "Backup")
	b.ANY("/database/backup-list", "BackupList")
	b.POST("/database/backup-delete", "BackupDelete")
	b.POST("/database/backup-download", "BackupDownload")

	// 检查项
	go func() {
		defer func() {
			if err := recover(); err != nil {
				pine.Logger().Errorf("start backup database failed", err)
			}
		}()
		setting, err := controllers.GetSetting(
			di.MustGet("*xorm.Engine").(*xorm.Engine),
			di.MustGet("cache.ICache").(cache.ICache))
		if err != nil {
			panic(err)
		}
		pine.Logger().Debug("启动自动备份线程")

		ticker := time.NewTicker(time.Hour) // 每小时调用一次

		for range ticker.C {
			pine.Logger().Debug("备份数据库")
			autoBackupHour, err := strconv.Atoi(setting["DATABASE_AUTO_BACKUP_TIME"])
			if err == nil &&
				autoBackupHour >= 0 &&
				autoBackupHour <= 23 &&
				time.Now().In(helper.GetLocation()).Hour() == autoBackupHour {
				msg, code := c.backup(setting, true)
				if code == 1 {
					pine.Logger().Error("自动备份数据库失败", msg)
				} else {
					pine.Logger().Print("自动备份数据库成功", msg)
				}
			}

		}
	}()

}

func (c *DatabaseController) Manager(orm *xorm.Engine) {
	if c.Ctx().URLParam("datagrid") != "" {
		mataDatas, _ := orm.DBMetas()

		var data = []map[string]interface{}{}
		for _, mataData := range mataDatas {
			total, _ := orm.Table(mataData.Name).Count()
			data = append(data, map[string]interface{}{
				"table_id": mataData.Name,
				"total":    total,
				"engine":   mataData.StoreEngine,
				"name":     mataData.Name,
				//"charset": mataData.Charset,
				"comment": mataData.Comment,
			})
		}

		c.Ctx().Render().JSON(data)
		return

	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("database_list_datagrid", "/b/database/manager?datagrid=true", helper.EasyuiOptions{
		"title":        models.NewMenuModel().CurrentPos(menuid),
		"toolbar":      "database_list_datagrid_toolbar",
		"singleSelect": "false",
		"pagination":   "false",
	}, helper.EasyuiGridfields{
		"":    {"field": "table_id", "checkbox": "true", "index": "0"},
		"数据表": {"field": "name", "width": "20", "index": "1"},
		"记录数": {"field": "total", "width": "30", "index": "2"},
		//"空间": {"field": "size", "width": "60", "index": "2"},
		"引擎": {"field": "engine", "width": "10", "index": "3"},
		//"编码": {"field": "charset", "width": "25", "index": "2"},
		//"更新时间": {"field": "updated", "width": "25", "index": "3"},
		"备注": {"field": "comment", "width": "70", "index": "4"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/database_list.html")
}

func (c *DatabaseController) Repair(orm *xorm.Engine) {
	tables := strings.Split(c.Ctx().FormValue("tables"), ",")
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
	tables := strings.Split(c.Ctx().FormValue("tables"), ",")
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
	if c.Ctx().URLParam("datagrid") == "true" {
		uploader := getStorageEngine(settingData)
		list, prefix, err := uploader.List(baseBackupDir)
		if err != nil {
			helper.Ajax("获取列表失败: "+err.Error(), 0, c.Ctx())
			return
		}

		var files = []map[string]string{}
		for _, v := range list {
			v = strings.TrimLeft(v, filepath.Join(prefix, baseBackupDir))
			files = append(files, map[string]string{
				"name":        v,
				"name_format": v,
			})
		}
		c.Ctx().Render().JSON(files)
		return
	}

	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("database_backup_list_datagrid", "/b/database/backup-list?datagrid=true", helper.EasyuiOptions{
		"title":      models.NewMenuModel().CurrentPos(menuid),
		"toolbar":    "database_backup_list_datagrid_toolbar",
		"pagination": "false",
	}, helper.EasyuiGridfields{
		"文件名": {"field": "name", "width": "20", "index": "0"},
		"操作":  {"field": "name_format", "index": "1", "formatter": "backupListOpFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/database_backup_list.html")
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
	orm := di.MustGet("*xorm.Engine").(*xorm.Engine)
	if settingData["UPLOAD_DATABASE_PASS"] == "" {
		return "请先设置备份数据库打包zip的密码", 1
	}
	fNameBaseName := strings.Replace(time.Now().In(helper.GetLocation()).Format(helper.TimeFormat), " ", "-", 1)
	fNameBaseName = strings.Replace(fNameBaseName, ":", "", 3)
	if auto {
		fNameBaseName = fmt.Sprintf("auto-%s", fNameBaseName)
	}
	uploader := getStorageEngine(settingData)
	uploadFile := fmt.Sprintf("%s/%s", baseBackupDir, fNameBaseName+".zip")
	buf := bytes.NewBuffer([]byte{})
	if err := orm.DumpAll(buf); err != nil {
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
