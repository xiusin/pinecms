package backend

import (
	"encoding/json"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

type AssetsManagerController struct {
	pine.Controller
}

// 定时备份任务功能

func (c *AssetsManagerController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/assets-manager/list", "Manager")
	b.ANY("/assets-manager/attachments-list", "AttachmentsList")
	b.ANY("/assets-manager/attachments-delete", "AttachmentsDelete")
	b.ANY("/assets-manager/edit", "Edit")
	b.ANY("/assets-manager/add", "Add")
	b.ANY("/assets-manager/theme", "Theme")
	b.GET("/assets-manager/theme-thumb/:theme", "ThemeThumb")
	b.POST("/assets-manager/set-theme", "SetTheme")
}

func (c *AssetsManagerController) Manager(orm *xorm.Engine) {
	conf := di.MustGet("pinecms.config").(*config.Config)
	themeDir := filepath.Join(conf.View.FeDirname, conf.View.Theme)
	fs, err := ioutil.ReadDir(themeDir)
	var files = []map[string]interface{}{}
	if err == nil {
		for _, f := range fs {
			if !f.IsDir() {
				if strings.HasSuffix(f.Name(), ".jet") {
					content, _ := ioutil.ReadFile(filepath.Join(themeDir, f.Name()))
					files = append(files, map[string]interface{}{
						"name":    f.Name(),
						"size":    f.Size(),
						"updated": f.ModTime().In(helper.GetLocation()).Format(helper.TimeFormat),
						"content": string(content),
					})
				}
			}
		}
	} else {
		c.Logger().Error("读取模板列表错误: "+err.Error(), 1, c.Ctx())
	}
	helper.Ajax(pine.H{
		"rows":  files,
		"total": len(files),
	}, 0, c.Ctx())
}

func (c *AssetsManagerController) Theme() {
	conf := di.MustGet("pinecms.config").(*config.Config)
	fs, err := ioutil.ReadDir(conf.View.FeDirname)
	if err != nil {
		helper.Ajax("读取模板主题目录失败: "+err.Error(), 1, c.Ctx())
		return
	}
	var dirs = []*ThemeConfig{}
	for _, f := range fs {
		if f.IsDir() {
			// 读取一个json配置文件
			contentByts,err := ioutil.ReadFile(filepath.Join(conf.View.FeDirname, f.Name(), "config.json"))
			if err != nil {
				continue
			}
			var configMap ThemeConfig
			err = json.Unmarshal(contentByts,&configMap)
			if err != nil {
				continue
			}
			configMap.Dir = f.Name()
			if configMap.Name == conf.View.Theme {
				configMap.IsDefault = true
			}
			dirs = append(dirs, &configMap)
		}
	}
	helper.Ajax(dirs, 0 , c.Ctx())
}

func (c *AssetsManagerController) SetTheme(cache cache.AbstractCache) {
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
	if cache.Set(controllers.CacheTheme, []byte(name)) == nil {
		conf.View.Theme = name
		helper.Ajax("设置主题成功", 0, c.Ctx())
	} else {
		helper.Ajax("设置主题失败", 1, c.Ctx())
	}
}

func (c *AssetsManagerController) Edit() {
	conf := di.MustGet("pinecms.config").(*config.Config)
	if c.Ctx().IsPost() {
		//origin := c.Ctx().URLParam("origin")
		name := c.Ctx().PostValue("name")
		content := c.Ctx().PostValue("content")
		f := filepath.Join(conf.View.FeDirname, conf.View.Theme, name)
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
	name := c.Ctx().GetString("name")
	typeName := c.Ctx().GetString("type")
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
	themeName := c.Ctx().Params().Get("theme")
	dirName := filepath.Join(conf.View.FeDirname, themeName, "thumb.png")
	c.Ctx().SetContentType("img/png")
	//todo 打开连接直接显示而不下载
	c.Ctx().SendFile(dirName)
}

func (c *AssetsManagerController) AttachmentsList() {
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")

	keywords := c.Ctx().GetString("keywords")
	list, total := models.NewAttachmentsModel().GetList(keywords, page, rows)
	helper.Ajax(pine.H{
		"rows": list,
		"total": total,
	}, 0, c.Ctx())
}

func (c *AssetsManagerController) AttachmentsDelete() {
	id, _ := c.Ctx().PostInt64("id")
	if id < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	if models.NewAttachmentsModel().Delete(id) {
		helper.Ajax("删除附件成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除附件失败", 1, c.Ctx())
	}
}
