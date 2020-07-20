package backend

import (
	"bytes"
	"encoding/json"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/cache/providers/bbolt"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
	bolt "go.etcd.io/bbolt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type ConfigItem map[string]interface{}

type ConfigStruct map[string]ConfigItem

type SettingController struct {
	pine.Controller
}

func (c *SettingController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/setting/site", "Site")
	b.GET("/setting/cache", "Cache")
	b.POST("/setting/del-cache", "DelCache")
}

func (c *SettingController) Site(cache cache.AbstractCache, orm *xorm.Engine) {
	if c.Ctx().IsPost() {
		fields := strings.Split(c.Ctx().GetString("fields"), ",")
		if len(fields) == 0 {
			helper.Ajax("参数错误, 请设置fields", 1, c.Ctx())
			return
		}
		table := &tables.Setting{}
		for _, field := range fields {
			table.Value = c.Ctx().PostString(field, "")
			orm.Table(table).Where("`key`=?", field).MustCols("value").Update(table)
		}
		cache.Delete(controllers.CacheSetting)
		helper.Ajax("更新配置成功", 0, c.Ctx())
		return
	}
	var setting []tables.Setting
	if err := orm.Asc("listorder").Find(&setting); err != nil {
		helper.Ajax("查找配置失败: "+err.Error(), 1, c.Ctx())
		return
	}
	var tabs = map[string]*TabsSchema{}
	var tabsKeys = map[string][]string{}
	for _, v := range setting {
		if _, ok := tabs[v.Group]; !ok {
			tabs[v.Group] = &TabsSchema{Title: v.Group, Hash: v.Group, Body: FormController{
				Type:     "form",
				Title:    v.Group + " 修改",
				Mode:     "horizontal",
				Controls: []FormControl{},
			}}
			tabsKeys[v.Group] = []string{v.Key}
		} else {
			tabsKeys[v.Group] = append(tabsKeys[v.Group], v.Key)
		}
		body, opts := tabs[v.Group].Body, strings.SplitN(v.Editor, ":", 2)
		fc := FormControl{
			Type:  opts[0],
			Name:  v.Key,
			Label: v.FormName,
			Value: v.Value,
		}
		if fc.Value == "" {
			fc.Value = v.Default
		}
		if len(opts) > 1 {
			switch opts[0] {
			case "select":
				var kvs = []KV{}
				if err := json.Unmarshal([]byte(opts[1]), &kvs); err != nil {
					c.Logger().Error(err, ":", opts[1])
				}
				fc.Options = kvs
			}
		}
		body.Controls = append(body.Controls, fc)
		tabs[v.Group].Body = body
	}
	var tabArr []*TabsSchema
	for k, v := range tabs {
		v.Body.Api = "POST setting/site?fields=" + strings.Join(tabsKeys[k], ",")
		tabArr = append(tabArr, v)
	}
	helper.Ajax(pine.H{
		"type": "tabs",
		"tabs": tabArr,
	}, 0, c.Ctx())
}

func (c *SettingController) Cache() {
	var list = []map[string]string{
		{
			"key":         "index",
			"name":        "清理主页HTML",
			"description": "首页静态页面文件,清理后首次访问自动生成",
		},
		{
			"key":         "list",
			"name":        "清理栏目HTML",
			"description": "分类页面各个分页静态缓存文件,清理后首次访问自动生成",
		},
		{

			"key":         "news",
			"name":        "清理文档HTML",
			"description": "内容详情页面静态缓存文件,清理后首次访问自动生成",
		},
		{
			"key":         "page",
			"name":        "清理单页HTML",
			"description": "分类信息单页缓存文件,清理后首次访问自动生成",
		},
	}
	list = append(list, map[string]string{
		"key":         "data",
		"name":        "数据缓存",
		"description": "系统内的数据缓存, 统一清理, 包括: 分类缓存,模型缓存,文档缓存",
	})
	helper.Ajax(pine.H{"rows": list, "total": len(list)}, 0, c.Ctx())
}

func (c *SettingController) DelCache(iCache cache.AbstractCache) {
	keys := strings.Split(c.Ctx().PostString("key"), ",")
	setting, _ := config.SiteConfig()
	for _, key := range keys {
		switch key {
		case "data":
			// todo 如果切换驱动这里切换到对应的缓存逻辑代码. 开发期间更会会有直接的错误提醒
			cacheHandler := iCache.(*bbolt.PineBolt)
			cacheHandler.BoltDB().Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(cacheHandler.BucketName))
				c := b.Cursor()
				prefix := []byte("pinecms.")
				for k, _ := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, _ = c.Next() {
					if err := b.Delete(k); err != nil {
						pine.Logger().Error(err)
						return err
					}
				}
				return nil
			})
		case "index":
			basePath := filepath.Join(setting["SITE_STATIC_PAGE_DIR"], "index.html")
			_ = os.Remove(basePath)
		case "list":
			fallthrough
		case "page":
			fallthrough
		case "news":
			cats := models.NewCategoryModel().GetNextCategory(0)
			delCacheByCategories(cats, setting["SITE_STATIC_PAGE_DIR"])
		}
	}
	helper.Ajax("清理缓存数据成功", 0, c.Ctx())
}

func delCacheByCategories(cats []tables.Category, cacheBaseDir string) {
	for _, cat := range cats {
		if cat.Type == 2 {
			continue
		}
		if cat.Dir == "" {
			cat.Dir = models.NewCategoryModel().GetUrlPrefix(cat.Catid)
		}
		basePath := filepath.Join(cacheBaseDir, cat.Dir)
		pine.Logger().Debugf("remove cache path: %s", basePath)
		err := os.RemoveAll(basePath)
		if err != nil {
			pine.Logger().Error(err)
		}
	}
}
