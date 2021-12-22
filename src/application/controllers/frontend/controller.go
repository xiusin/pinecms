package frontend

import (
	"fmt"
	jet2 "github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/config"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"xorm.io/xorm"
)

type IndexController struct {
	pine.Controller
}

func (c *IndexController) RegisterRoute(b pine.IRouterWrapper) {
	// 必须放到最后 否则搜索路由时会优先被此路由拦截到
	b.GET("/search.go", "Search")
	b.GET("/*pagename", "Bootstrap")
}

func viewDataToJetMap(binding map[string]interface{}) jet2.VarMap {
	vars := jet2.VarMap{}
	for k, v := range binding {
		vars[k] = reflect.ValueOf(v)
	}
	return vars
}

func (c *IndexController) setTemplateData() {
	var detailUrl = func(aid string, tid ...string) string {
		if len(tid) == 0 {
			tid = []string{c.Ctx().Params().Get("tid")}
		}
		iaid, _ := strconv.Atoi(aid)
		itid, _ := strconv.Atoi(tid[0])
		if itid == 0 {
			pine.Logger().Error("传入tid参数错误")
			return ""
		}
		urlPrefix := models.NewCategoryModel().GetUrlPrefix(int64(itid))
		return fmt.Sprintf("/%s/%d.html", urlPrefix, iaid)
	}
	c.Ctx().Set("detail_url", detailUrl)
	if tid, _ := c.Ctx().Params().GetInt64("tid"); tid <= 0 {
		c.ViewData("isActive", func(id int64) bool {
			treeCats := models.NewCategoryModel().GetPosArr(tid)
			for _, v := range treeCats {
				if v.Catid == id {
					return true
				}
			}
			return false
		})
	}
	if c.Ctx().Params().Get("page") != "" {
		p, _ := c.Ctx().Params().GetFloat64("page", 1)
		c.Ctx().Render().ViewData("page", p)
	}
	c.Ctx().Render().ViewData("detail_url", detailUrl)
}

func getOrmSess(model *tables.DocumentModel) *xorm.Session {
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(controllers.GetTableName(model.Table)).Where("status = 1").Where("deleted_time IS NULL")
}

func template(tpl string) string {
	//todo 支持mobile pc
	conf := di.MustGet(controllers.ServiceConfig).(*config.Config)
	path := filepath.Join(conf.View.Theme, tpl)
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}

func GetStaticFile(filename string) string {
	setting, _ := config.SiteConfig()
	if setting["SITE_STATIC_PAGE_DIR"] == "" {
		setting["SITE_STATIC_PAGE_DIR"] = "resources/html"
	}
	return filepath.Join(setting["SITE_STATIC_PAGE_DIR"], filename)
}
