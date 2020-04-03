package backend

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/config"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
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
		var setting []*tables.Setting
		act := c.Ctx().URLParam("dosubmit")
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
	menuid, err := c.Ctx().URLParamInt64("menuid")
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
				"name":        "首页缓存",
				"description": "首页静态页面文件,清理后首次访问自动生成",
			},
			{
				"key1":        "list",
				"key":         "list",
				"name":        "分类缓存",
				"description": "分类页面各个分页静态缓存文件,清理后首次访问自动生成",
			},
			{

				"key1":        "news",
				"key":         "news",
				"name":        "内容缓存",
				"description": "内容详情页面静态缓存文件,清理后首次访问自动生成",
			},
			{
				"key1":        "page",
				"key":         "page",
				"name":        "单页缓存",
				"description": "分类信息单页缓存数据,清理后首次访问自动生成",
			},
			//{
			//	"key1":        "log",
			//	"key":         "log",
			//	"name":        "日志数据",
			//	"description": "网站运行过程中会记录各种错误日志，以文件的方式保存, 可删除",
			//},
		}

		models,_ := models.NewDocumentModel().GetList(1, 1000)

		for _, v := range models {
			list = append(list, map[string]string{
				"key1":        fmt.Sprintf("model_%d", v.Id),
				"key":         fmt.Sprintf("model_%d", v.Id),
				"name":        fmt.Sprintf("模型 %s 缓存", v.Name),
				"description": fmt.Sprintf("可直接清理模型 %s 下所有缓存. ", v.Name),
			})
		}

		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": len(list)})

		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
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

func (c *SettingController) DelCache() {
	keys := strings.Split(c.Ctx().PostString("key"), ",")
	setting, _ := config.SiteConfig()
	for _, key := range keys {
		switch key {
		case "index":
			basePath := filepath.Join(setting["SITE_STATIC_PAGE_DIR"], "index.html")
			os.Remove(basePath)
		case "list":
			fallthrough
		case "page":
			fallthrough
		case "news":
			cats := models.NewCategoryModel().GetNextCategory(0)
			delCacheByCategories(cats, setting["SITE_STATIC_PAGE_DIR"])
		default:
			if strings.HasPrefix(key, "model_") {
				modelid,_ := strconv.Atoi(strings.TrimPrefix(key, "model_"))
				if modelid == 0 {
					continue
				}
				cats,_ := models.NewCategoryModel().GetCategoryByModelID(int64(modelid))
				delCacheByCategories(cats, setting["SITE_STATIC_PAGE_DIR"])
			} else {
				helper.Ajax("错误请求", 1, c.Ctx())
				return
			}
		}
	}

	helper.Ajax("清理缓存数据成功", 0, c.Ctx())
}

func delCacheByCategories(cats []tables.Category, cacheBaseDir string)  {
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
