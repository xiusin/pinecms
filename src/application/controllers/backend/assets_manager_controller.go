package backend

import (
	"github.com/xiusin/pine/cache"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

type AssetsManagerController struct {
	pine.Controller
}

// 定时备份任务功能

func (c *AssetsManagerController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/assets-manager/list", "Manager")
	b.ANY("/assets-manager/edit", "Edit")
	b.ANY("/assets-manager/add", "Add")
	b.ANY("/assets-manager/theme", "Theme")
	b.ANY("/assets-manager/theme-thumb/:theme", "ThemeThumb")
	b.POST("/assets-manager/set-theme", "SetTheme")
}

func (c *AssetsManagerController) Manager(orm *xorm.Engine) {
	if c.Ctx().URLParam("datagrid") == "true" {
		conf := di.MustGet("pinecms.config").(*config.Config)
		fs, err := ioutil.ReadDir(filepath.Join(conf.View.FeDirname, conf.View.Theme))
		if err != nil {
			helper.Ajax("读取模板列表错误: "+err.Error(), 1, c.Ctx())
			return
		}
		var files = []map[string]interface{}{}
		for _, f := range fs {
			if !f.IsDir() {
				if strings.HasSuffix(f.Name(), ".jet") {
					files = append(files, map[string]interface{}{
						"name":    f.Name(),
						"fname":   f.Name(),
						"size":    f.Size(),
						"updated": f.ModTime().In(helper.GetLocation()).Format(helper.TimeFormat),
					})
				}
			}
		}
		c.Ctx().Render().JSON(files)
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("assets_list_datagrid", "/b/assets-manager/list?datagrid=true", helper.EasyuiOptions{
		"title":      models.NewMenuModel().CurrentPos(menuid),
		"toolbar":    "assets_list_datagrid_toolbar",
		"pagination": "false",
	}, helper.EasyuiGridfields{
		"文件名":  {"field": "name", "width": "40", "index": "0"},
		"文件大小": {"field": "size", "width": "20", "index": "1", "formatter": "fileSizeFormatter"},
		"修改时间": {"field": "updated", "width": "20", "index": "2"},
		"操作":   {"field": "fname", "index": "3", "formatter": "assetListOpFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/assets_list.html")
}

func (c *AssetsManagerController) Theme() {
	conf := di.MustGet("pinecms.config").(*config.Config)
	fs, err := ioutil.ReadDir(conf.View.FeDirname)
	if err != nil {
		helper.Ajax("读取模板主题目录失败: "+err.Error(), 1, c.Ctx())
		return
	}
	var dirs = []map[string]interface{}{}
	for _, f := range fs {
		if f.IsDir() {
			dirs = append(dirs, map[string]interface{}{
				"name":     f.Name(),
				"thumb":  "/b/assets-manager/theme-thumb/" + f.Name(),
			})
		}
	}
	c.Ctx().Render().ViewData("dirs", dirs)
	c.Ctx().Render().ViewData("theme", conf.View.Theme)
	c.Ctx().Render().HTML("backend/assets_theme.html")
}


func (c *AssetsManagerController) SetTheme(cache cache.ICache) {
	conf := di.MustGet("pinecms.config").(*config.Config)
	name := c.Ctx().FormValue("theme")
	if name == "" {
		helper.Ajax("模板主题参数错误", 1, c.Ctx())
		return
	}
	fs, err := os.Stat(filepath.Join(conf.View.FeDirname, name))
	if err != nil || !fs.IsDir() {
		helper.Ajax("模板主题不存在", 1, c.Ctx())
		return
	}
	if cache.Set("pinecms_theme", []byte(name))	== nil {
		conf.View.Theme = name
		helper.Ajax("设置主题成功", 0, c.Ctx())
	} else {
		helper.Ajax("设置主题失败", 1, c.Ctx())
	}
}

func (c *AssetsManagerController) Edit(orm *xorm.Engine) {
	conf := di.MustGet("pinecms.config").(*config.Config)
	if c.Ctx().IsPost() {
		//origin := c.Ctx().URLParam("origin")
		name := c.Ctx().PostValue("name")
		content := c.Ctx().PostValue("content")
		f := filepath.Join(conf.View.FeDirname,conf.View.Theme, name)
		_, err := os.Stat(f)
		if err != nil {
			helper.Ajax("获取模板状态失败："+err.Error(), 1, c.Ctx())
			return
		}
		if err := ioutil.WriteFile(f, []byte(content), os.ModePerm); err != nil {
			helper.Ajax("写入模板内容失败: "+err.Error(), 1, c.Ctx())
			return
		}
		helper.Ajax("修改成功", 0, c.Ctx())
		return
	}
	name := c.Ctx().URLParam("name")
	typeName := c.Ctx().URLParam("type")
	if name == "" || typeName == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	content, err := ioutil.ReadFile(filepath.Join(conf.View.FeDirname, conf.View.Theme, name))
	if err != nil {
		helper.Ajax("读取模板错误: "+err.Error(), 1, c.Ctx())
		return
	}
	c.Ctx().Render().ViewData("name", name)
	c.Ctx().Render().ViewData("typeName", typeName)
	c.Ctx().Render().ViewData("content", template.HTML(content))
	c.Ctx().Render().HTML("backend/assets_edit.html")
}

func (c *AssetsManagerController) Add(orm *xorm.Engine) {
	conf := di.MustGet("pinecms.config").(*config.Config)
	if c.Ctx().IsPost() {
		name := c.Ctx().PostValue("name")
		if !strings.HasSuffix(name, ".jet") {
			helper.Ajax("模板文件请以'.jet'为后缀", 1, c.Ctx())
			return
		}
		content := c.Ctx().PostValue("content")
		f := filepath.Join(conf.View.FeDirname, conf.View.Theme, name)
		_, err := os.Stat(f)
		if err == nil {
			helper.Ajax("模板已经存在", 1, c.Ctx())
			return
		}
		if err := ioutil.WriteFile(f, []byte(content), os.ModePerm); err != nil {
			helper.Ajax("写入模板内容失败: "+err.Error(), 1, c.Ctx())
			return
		}
		helper.Ajax("添加模板成功", 0, c.Ctx())
		return
	}
	c.Ctx().Render().HTML("backend/assets_add.html")
}

func (c *AssetsManagerController) ThemeThumb() {
	conf := di.MustGet("pinecms.config").(*config.Config)
	themeName := c.Param().Get("theme")
	dirName := filepath.Join(conf.View.FeDirname, themeName, "thumb.png")
	c.Ctx().Writer().Header().Set("Content-type", "img/png")
	//todo 打开连接直接显示而不下载
	c.Ctx().SendFile(dirName)
}
