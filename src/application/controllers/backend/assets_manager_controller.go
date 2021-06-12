package backend

import (
	"encoding/json"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
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
	BaseController
}

func (c *AssetsManagerController) Construct() {
	c.KeywordsSearch = []KeywordWhere{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.Link{}
	c.Entries = &[]*tables.Link{}

}

func (c *AssetsManagerController) PostList() {
	conf := di.MustGet("pinecms.config").(*config.Config)
	themeDir := filepath.Join(conf.View.FeDirname, conf.View.Theme)
	fs, err := ioutil.ReadDir(themeDir)

	var p listParam
	parseParam(c.Ctx(), &p)

	var files []map[string]interface{}
	if err == nil {
		for _, f := range fs {
			if !f.IsDir() {
				if strings.HasSuffix(f.Name(), ".jet") {
					files = append(files, map[string]interface{}{
						"id":      f.Name(),
						"name":    f.Name(),
						"size":    f.Size(),
						"updated": f.ModTime().In(helper.GetLocation()).Format(helper.TimeFormat),
					})
				}
			}
		}
	} else {
		c.Logger().Error("读取模板列表错误: "+err.Error(), 1, c.Ctx())
	}
	helper.Ajax(pine.H{
		"list": files,
		"pagination": pine.H{
			"page":  p.Page,
			"size":  p.Size,
			"total": len(files),
		},
	}, 0, c.Ctx())

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
	conf := di.MustGet("pinecms.config").(*config.Config)
	if len(p.Id) == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	fullPath := filepath.Join(conf.View.FeDirname, conf.View.Theme, p.Id)
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
			contentByts, err := ioutil.ReadFile(filepath.Join(conf.View.FeDirname, f.Name(), "config.json"))
			if err != nil {
				continue
			}
			var configMap ThemeConfig
			err = json.Unmarshal(contentByts, &configMap)
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
	helper.Ajax(dirs, 0, c.Ctx())
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

func (c *AssetsManagerController) GetInfo() {
	conf := di.MustGet("pinecms.config").(*config.Config)
	name := c.Ctx().GetString("id")
	if name == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	fullPath := filepath.Join(conf.View.FeDirname, conf.View.Theme, name)
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
	helper.Ajax(map[string]interface{}{
		"id":      name,
		"name":    name,
		"size":    stat.Size(),
		"updated": stat.ModTime().In(helper.GetLocation()).Format(helper.TimeFormat),
		"content": string(content),
	}, 0, c.Ctx())
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
