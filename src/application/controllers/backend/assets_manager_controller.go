package backend

import (
	"encoding/json"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"xorm.io/xorm"
)

type AssetsManagerController struct {
	BaseController
	conf *config.Config
}

func (c *AssetsManagerController) Construct() {
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Orm = helper.GetORM()
	c.conf = di.MustGet("pinecms.config").(*config.Config)
}

func (c *AssetsManagerController) GetSelect() {
	themeDir := filepath.Join(c.conf.View.FeDirname, c.conf.View.Theme)
	files := c.getTempList(themeDir)
	var kv []tables.KV
	for _, file := range files {
		kv = append(kv, tables.KV{
			Label: file["id"].(string),
			Value: file["id"].(string),
		})
	}
	helper.Ajax(kv, 0, c.Ctx())
}

func (c *AssetsManagerController) getTempList(dir string) []map[string]interface{} {
	var files []map[string]interface{}
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		f, _ := d.Info()
		if strings.HasSuffix(f.Name(), ".jet") {
			files = append(files, map[string]interface{}{
				"id":      strings.TrimLeft(strings.TrimPrefix(path, dir), "/"),
				"name":    strings.TrimLeft(strings.TrimPrefix(path, dir), "/"),
				"size":    f.Size(),
				"updated": f.ModTime().In(helper.GetLocation()).Format(helper.TimeFormat),
			})
		}
		return nil
	})
	return files
}

func (c *AssetsManagerController) PostList() {
	themeDir := filepath.Join(c.conf.View.FeDirname, c.conf.View.Theme)
	helper.Ajax(helper.DirTree(themeDir), 0, c.Ctx())
}

func (c *AssetsManagerController) PostEdit() {
	p := struct {
		Content string `json:"content"`
		Name    string `json:"name"`
		Id      string `json:"id"`
	}{}
	if err := c.Ctx().BindJSON(&p); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	if len(p.Id) == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	fullPath := filepath.Join(c.conf.View.FeDirname, c.conf.View.Theme, p.Id)
	f, err := os.OpenFile(fullPath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	defer f.Close()
	if _, err := f.WriteString(p.Content); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	helper.Ajax("修改成功", 0, c.Ctx())
}

func (c *AssetsManagerController) GetThemes() {
	files, err := ioutil.ReadDir(c.conf.View.FeDirname)
	if err != nil {
		helper.Ajax("读取主题目录失败: "+err.Error(), 1, c.Ctx())
		return
	}
	var dirs []*ThemeConfig
	for _, f := range files {
		if f.IsDir() {
			contentByts, err := ioutil.ReadFile(filepath.Join(c.conf.View.FeDirname, f.Name(), "config.json"))
			if err != nil {
				continue
			}
			var tConf ThemeConfig
			err = json.Unmarshal(contentByts, &tConf)
			if err != nil {
				continue
			}
			tConf.Dir = f.Name()
			if tConf.Dir == c.conf.View.Theme {
				tConf.IsDefault = true
			}
			dirs = append(dirs, &tConf)
		}
	}
	helper.Ajax(dirs, 0, c.Ctx())
}

func (c *AssetsManagerController) PostTheme(cache cache.AbstractCache) {
	var p = map[string]interface{}{}
	_ = c.Ctx().BindJSON(&p)
	name := p["theme"].(string)

	if name == "" {
		helper.Ajax("模板主题参数错误", 1, c.Ctx())
		return
	}
	fileSys, err := os.Stat(filepath.Join(c.conf.View.FeDirname, name))
	if err != nil || !fileSys.IsDir() {
		helper.Ajax("模板主题不存在", 1, c.Ctx())
		return
	}
	if cache.Set(controllers.CacheTheme, []byte(name)) == nil {
		c.conf.View.Theme = name
		helper.Ajax("设置主题成功", 0, c.Ctx())
	} else {
		helper.Ajax("设置主题失败", 1, c.Ctx())
	}
}

func (c *AssetsManagerController) GetInfo() {
	fullPath, _ := c.Ctx().GetString("path")
	if fullPath == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	f, err := os.Open(fullPath)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	defer f.Close()
	stat, _ := f.Stat()
	var content = make([]byte, stat.Size())
	_, err = f.Read(content)
	if err != nil {
		helper.Ajax("读取模板错误: "+err.Error(), 1, c.Ctx())
		return
	}
	helper.Ajax(string(content), 0, c.Ctx())
}

func (c *AssetsManagerController) PostAdd(orm *xorm.Engine) {
	if c.Ctx().IsPost() {
		name, _ := c.Ctx().GetString("name")
		if !strings.HasSuffix(name, ".jet") {
			helper.Ajax("模板文件请以'.jet'为后缀", 1, c.Ctx())
			return
		}
		content, _ := c.Ctx().GetString("content")
		f := filepath.Join(c.conf.View.FeDirname, c.conf.View.Theme, name)
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

func (c *AssetsManagerController) GetThumb() {
	themeName, _ := c.Ctx().GetString("id")
	dirName := filepath.Join(c.conf.View.FeDirname, themeName, "thumb.png")
	c.Ctx().SetContentType("img/png")
	c.Ctx().SendFile(dirName)
}
