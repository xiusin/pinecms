package backend

import (
	"bytes"
	"encoding/json"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/cache/providers/bbolt"
	"github.com/xiusin/pinecms/src/config"
	bolt "go.etcd.io/bbolt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/xiusin/pinecms/src/application/controllers"
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
	b.ANY("/setting/cache", "Cache")
	b.POST("/setting/del-cache", "DelCache")
}

func (c *SettingController) Site(iCache cache.AbstractCache, orm *xorm.Engine) {
	if c.Ctx().IsPost() {
		var setting []tables.Setting
		act := c.Ctx().GetString("dosubmit")
		var setval []ConfigItem
		if act == "" {
			if err := orm.Asc("listorder").Find(&setting); err != nil {
				return
			}
			//没有值的配置项自动合并
			for _, v := range setting {
				v.EditorOpt = v.Editor
				if strings.HasPrefix(v.Editor, "{") {
					var options = map[string]interface{}{}
					if err := json.Unmarshal([]byte(v.Editor), &options); err == nil {
						v.EditorOpt = options
					}
				}
				setval = append(setval, ConfigItem{
					"key":     v.Key,
					"name":    v.FormName,
					"group":   v.Group,
					"editor":  v.EditorOpt,
					"default": v.Default,
					"value":   v.Value,
				})
			}
			result := map[string]interface{}{
				"rows":  setval,
				"total": len(setval),
			}
			c.Ctx().Render().JSON(result)
			return
		}
		post := c.Ctx().PostData()
		for k, v := range post {
			if k == "dosubmit" || len(v) == 0 {
				continue
			}
			c.Ctx().Value("orm").(*xorm.Engine).Table(new(tables.Setting)).Where("`key`=?", k).MustCols("value").Update(&tables.Setting{Value: v[0]})
		}
		iCache.Delete(controllers.CacheSetting)
		helper.Ajax("更新配置信息成功", 0, c.Ctx())
		return
	}
	menuid, err := c.Ctx().GetInt64("menuid")
	if err != nil {
		menuid = 0
	}
	currentpos := models.NewMenuModel().CurrentPos(menuid)
	grid := helper.Propertygrid("setting_site_propertygrid", helper.EasyuiOptions{
		"title":   currentpos,
		"url":     "/b/setting/site?grid=propertygrid",
		"toolbar": "setting_site_propertygrid_toolbar",
	})
	c.Ctx().Render().ViewData("grid", template.HTML(grid))
	c.Ctx().Render().HTML("backend/setting_site.html")
}

func (c *SettingController) Cache(iCache cache.AbstractCache) {
	if c.Ctx().GetString("getlist") != "" {
		var list = []map[string]string{
			{
				"key1":        "index",
				"key":         "index",
				"name":        "清理主页HTML",
				"description": "首页静态页面文件,清理后首次访问自动生成",
			},
			{
				"key1":        "list",
				"key":         "list",
				"name":        "清理栏目HTML",
				"description": "分类页面各个分页静态缓存文件,清理后首次访问自动生成",
			},
			{

				"key1":        "news",
				"key":         "news",
				"name":        "清理文档HTML",
				"description": "内容详情页面静态缓存文件,清理后首次访问自动生成",
			},
			{
				"key1":        "page",
				"key":         "page",
				"name":        "清理单页HTML",
				"description": "分类信息单页缓存文件,清理后首次访问自动生成",
			},
		}

		list = append(list, map[string]string{
			"key1":        "data",
			"key":         "data",
			"name":        "数据缓存",
			"description": "系统内的数据缓存, 统一清理, 包括: 分类缓存,模型缓存,文档缓存",
		})

		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": len(list)})

		return
	}
	menuid, _ := c.Ctx().GetInt64("menuid")
	table := helper.Datagrid("setting_cache_datagrid", "/b/setting/cache?getlist=true", helper.EasyuiOptions{
		"title":        models.NewMenuModel().CurrentPos(menuid),
		"toolbar":      "setting_cache_datagrid_toolbar",
		"singleSelect": "false",
		"pagination":   "false",
	}, helper.EasyuiGridfields{
		"选择": {"field": "key1", "checkbox": "true", "index": "0"},
		"缓存": {"field": "name", "width": "20", "index": "1"},
		"描述": {"field": "description", "width": "30", "index": "2"},
		"操作": {"field": "key", "width": "50", "index": "3", "formatter": "settingCacheOperateFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/setting_cache.html")
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
