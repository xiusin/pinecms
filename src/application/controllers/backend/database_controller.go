package backend

import (
	"bytes"
	"fmt"
	"io"
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
	b.ANY("/database/list", "Manager")
	b.POST("/database/repair", "Repair")
	b.POST("/database/optimize", "Optimize")
	b.POST("/database/backup", "Backup")
	b.POST("/database/exec", "Exec")
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

func (c *DatabaseController) Exec(orm *xorm.Engine) {
	var sql = struct {
		Sql string `json:"sql"`
	}{}
	c.Ctx().BindJSON(&sql)
	sql.Sql = strings.Trim(sql.Sql, "\n ")
	if sql.Sql == "" {
		helper.Ajax("请填入要操作的SQL语句", 1, c.Ctx())
		return
	}
	if len(sql.Sql) < 6 {
		helper.Ajax("请填入完整的SQL", 1, c.Ctx())
		return
	}
	act := strings.ToLower(sql.Sql[:6])
	if act == "select" {
		helper.Ajax("暂不支持查询操作", 1, c.Ctx())
		return
	}
	if (act == "delete" || act == "update") && !strings.Contains(strings.ToLower(sql.Sql), "where") {
		helper.Ajax("更新或删除操作必须带有WHERE语句", 1, c.Ctx())
		return
	}
	_, err := orm.Exec(sql.Sql)
	if err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	helper.Ajax("操作成功", 0, c.Ctx())
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
	msg, code := c.backup(settingData)
	helper.Ajax(msg, int64(code), c.Ctx())
}

func (c *DatabaseController) backup(settingData map[string]string) (msg string, erode int) {
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
	defer zipw.Close()
	w, err := zipw.Encrypt(fNameBaseName+".sql", settingData["UPLOAD_DATABASE_PASS"])
	if err != nil {
		return "打包zip失败: " + err.Error(), 1
	}
	_, err = io.Copy(w, buf)
	if err != nil {
		return "打包zip失败: " + err.Error(), 1
	}
	zipw.Flush()
	f, err := uploader.Upload(uploadFile, zipsc)
	if err != nil {
		return "备份表数据失败: " + err.Error(), 1
	}
	return "备份数据库成功: " + f, 0
}
