package frontend

import (
	"fmt"
	jet2 "github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type IndexController struct {
	pine.Controller
}

var debug = true

func (c *IndexController) RegisterRoute(b pine.IRouterWrapper) {
	// 必须放到最后 否则搜索路由时会优先被此路由拦截到
	b.GET("/*pagename", "Bootstrap")
}

func viewDataToJetMap(binding map[string]interface{}) jet2.VarMap {
	vars := jet2.VarMap{}
	for k, v := range binding {
		vars[k] = reflect.ValueOf(v)
	}
	return vars
}

var catTable = &tables.Category{}

func (c *IndexController) Bootstrap(orm *xorm.Engine) {
	//c.ViewData("__typeid", 2)
	//c.View(template("testlib.jet"))
	//return
	// todo 拦截存在静态文件的问题, 不过最好交给nginx等服务器转发
	// todo 开启前端资源缓存 304
	pageName := strings.Trim(strings.ReplaceAll(c.Ctx().Params().Get("pagename"), "//", "/"), "/") // 必须包含.html
	switch pageName {
	case "index.html", "":
		c.Index()
	default:
		urlPartials := strings.Split(pageName, "/")
		var last string
		var fileName string
		var isDetail bool
		// 查找最后一段路由对应的地址
		if strings.HasSuffix(pageName, ".html") {
			last = urlPartials[len(urlPartials)-2]
			fileName = urlPartials[len(urlPartials)-1]
			// 分析页码
			if strings.HasPrefix(fileName, "index_") {
				fileInfo := strings.Split(fileName, "_") // index_2.html => 某个分类的第二页
				c.Ctx().Params().Set("page", strings.TrimSuffix(fileInfo[1], ".html"))
			} else if fileName != "index.html" {
				isDetail = true
				c.Ctx().Params().Set("aid", strings.TrimSuffix(fileName, ".html")) // 设置文档ID
			}
		} else {
			last = urlPartials[len(urlPartials)-1] // 目录名
			fileName = "index.html"
			pageName = filepath.Join(pageName, fileName)
			c.Ctx().Params().Set("page", "1")
		}
		var cat tables.Category
		exist, _ := orm.Table(catTable).Where("dir = ?", last).Get(&cat)
		if !exist {
			// 拆分出来想要的数据 page_{tid}
			if strings.HasPrefix(last, "page_") {
				c.Ctx().Params().Set("tid", strings.TrimPrefix(last, "page_"))
				c.Page()
				return
			} else {
				infos := strings.Split(last, "_") // 根据模型{model_table}_{tid}拆分信息
				modelTable := infos[0]
				model := &tables.DocumentModel{}
				exist, _ = orm.Table(model).Where("`table` = ?", modelTable).Get(model)
				if exist {
					c.Ctx().Params().Set("tid", infos[1])
					c.List(pageName)
					return
				}
			}
			c.Ctx().Abort(http.StatusNotFound)
			c.Logger().Error("路由地址无法匹配完整内容", c.Ctx().Request().URL.Path)
			return
		}
		// 匹配所有内容
		prefix := models.NewCategoryModel().GetUrlPrefix(cat.Catid)
		if !strings.HasPrefix(pageName, prefix) {
			c.Logger().Error("路由地址无法匹配完整内容", c.Ctx().Request().URL.Path)
			c.Ctx().Abort(http.StatusNotFound)
			return
		}
		c.Ctx().Params().Set("tid", strconv.Itoa(int(cat.Catid)))
		if isDetail {
			c.Detail(pageName)
		} else if cat.Type == 0 {
			c.List(pageName)
		} else {
			c.Page()
		}
	}
}

func (c *IndexController) Index() {
	c.setTemplateData()
	indexPage := "index.html"
	pageFilePath := GetStaticFile(indexPage)
	os.MkdirAll(filepath.Dir(pageFilePath), os.ModePerm)
	f, err := os.OpenFile(pageFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		c.Logger().Error(err)
		return
	}
	defer f.Close()
	jet := pine.Make(controllers.ServiceJetEngine).(*jet.PineJet)
	temp, err := jet.GetTemplate(template("index.jet"))
	if err != nil {
		c.Logger().Error(err)
		return
	}
	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), nil)
	if err != nil {
		c.Logger().Error(err)
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().Writer().Write(data)
}

func (c *IndexController) List(pageFilePath string) {
	c.setTemplateData()
	pageFilePath = GetStaticFile(pageFilePath)
	queryTid, _ := c.Ctx().GetInt64("tid")
	tid, _ := c.Ctx().Params().GetInt64("tid", queryTid)
	if tid < 1 {
		c.Ctx().Abort(404)
		return
	}
	category, err := models.NewCategoryModel().GetCategoryFullWithCache(tid)
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(404)
		return
	}
	setting, _ := config.SiteConfig()
	var globalSize, _ = strconv.Atoi(setting["SITE_PAGE_SIZE"])
	if globalSize == 0 {
		globalSize = 25
	}
	pageNum, _ := c.Ctx().Params().GetInt("page", 1)
	if pageNum < 1 {
		pageNum = 1
	}
	c.ViewData("typeid", tid)
	c.ViewData("__typeid", tid)
	c.ViewData("pagelist", func(listsize ...int) string {
		if len(listsize) == 0 {
			listsize = append(listsize, globalSize)
		}
		total, _ := getOrmSess(category.Model).Where("catid = ?", tid).Count()
		totalPage := int(math.Ceil(float64(total) / float64(listsize[0])))
		pagenum, _ := c.Ctx().GetInt("page")
		if pagenum < 1 {
			pagenum = 1
		} else if pagenum > totalPage {
			pagenum = totalPage
		}
		return helper.NewPage("/"+models.NewCategoryModel().GetUrlPrefix(tid), pageNum, listsize[0], int(total)).String()
	})

	if category.Model.FeTplList == "" {
		category.Model.FeTplList = "list_article.jet"
	}

	os.MkdirAll(filepath.Dir(pageFilePath), os.ModePerm)
	f, err := os.OpenFile(pageFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	defer f.Close()
	jet := pine.Make(controllers.ServiceJetEngine).(*jet.PineJet)
	temp, err := jet.GetTemplate(template(category.Model.FeTplList))
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), struct{ Field *tables.Category }{Field: category})
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().Writer().Write(data)
}

func (c *IndexController) Detail(pagename string) {
	c.setTemplateData()
	pageFilePath := GetStaticFile(pagename)
	aid, _ := c.Param().GetInt64("aid")
	tid, _ := c.Param().GetInt64("tid")
	var err error
	if tid < 1 || aid < 1 {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	category, err := models.NewCategoryModel().GetCategoryFullWithCache(tid)
	if err != nil {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	sess := getOrmSess(category.Model).Where("id = ?", aid).Where("catid = ?", tid).Limit(1)
	rest, err := sess.QueryString()
	if err != nil || len(rest) == 0 {
		pine.Logger().Errorf("查找tableName:%s错误: %s", category.Model.Table, err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	article := rest[0]
	article["catname"] = category.Catname
	article["position"] = getCategoryPos(tid)
	article["caturl"] = helper.ListUrl(int(tid))
	c.ViewData("__typeid", tid)
	detailUrlFunc := c.Ctx().Value("detail_url").(func(string, ...string) string)
	c.ViewData("prenext", func() string {
		var str string
		// aid 根据aid 读取上一篇和下一篇
		pre, _ := getOrmSess(category.Model).Where("id < ?", aid).Desc("id").Limit(1).QueryString()
		if len(pre) > 0 {
			str += "<p><a href='" + detailUrlFunc(pre[0]["id"], pre[0]["catid"]) + "'>上一篇：" + pre[0]["title"] + "</a></p>"
		}
		next, _ := getOrmSess(category.Model).Where("id > ?", aid).Asc("id").Limit(1).QueryString()
		if len(next) > 0 {
			str += "<p><a href='" + detailUrlFunc(next[0]["id"], next[0]["catid"]) + "'>下一篇: " + next[0]["title"] + "</a></p>"
		}
		return str
	})
	c.ViewData("likearticle", func(row int, kws ...string) []map[string]string {
		var keywords []interface{}
		if kws == nil || len(kws) == 0 {
			kws = append(kws, article["keywords"], article["tags"])
		}
		if row < 1 {
			row = 10
		}
		var conds = []string{}
		sess := getOrmSess(category.Model).Where("id <> ?", article["id"])
		for _, kw := range kws {
			splitKeywords := strings.Split(kw, ",")
			for _, keyword := range splitKeywords {
				conds = append(conds, "CONCAT(keywords,' ',title, ' ', tags) LIKE ?")
				keywords = append(keywords, "%"+strings.Trim(keyword, "")+"%")
			}
		}
		sess.And(strings.Join(conds, " OR "), keywords...)
		articles, _ := sess.Limit(row).Desc("id").QueryString()
		for k := range articles {
			articles[k]["url"] = detailUrlFunc(articles[k]["id"], articles[k]["catid"])
		}
		return articles
	})
	if category.Model.FeTplDetail == "" {
		category.Model.FeTplDetail = "detail.jet"
	}
	os.MkdirAll(filepath.Dir(pageFilePath), os.ModePerm)
	f, err := os.OpenFile(pageFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		c.Ctx().WriteString(err.Error())
		return
	}
	defer f.Close()
	jet := pine.Make(controllers.ServiceJetEngine).(*jet.PineJet)
	temp, err := jet.GetTemplate(template(category.Model.FeTplDetail))
	if err != nil {
		c.Ctx().WriteString(err.Error())
		return
	}
	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), struct{ Field map[string]string }{Field: article})
	if err != nil {
		c.Ctx().WriteString(err.Error())
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().Writer().Write(data)
}

func (c *IndexController) Page() {
	c.setTemplateData()
	tid, _ := c.Ctx().Params().GetInt64("tid")
	if tid < 1 {
		c.Ctx().Abort(404)
		return
	}
	page := models.NewPageModel().GetPage(tid)
	if page == nil {
		c.Ctx().Abort(404)
		return
	}

	page.Position = getCategoryPos(tid)
	c.ViewData("field", page)
	//data, _ := ioutil.ReadFile(pageFilePath)
	//c.Ctx().Writer().Write(data)
	c.View(template("page.jet"))
}

func (c *IndexController) Click() {
	//aid, _ := c.Ctx().GetInt64("aid")
	//tid, _ := c.Ctx().GetInt64("tid")
	//if aid < 1 || tid < 1 {
	//	c.Ctx().Abort(http.StatusNotFound)
	//	return
	//}
	//clickCache := fmt.Sprintf("click_%d_%d", tid, aid)
	//info := c.Ctx().GetCookie(clickCache)
	//if len(info) == 0 {
	//	res, err := di.MustGet("orm").(*xorm.Engine).Table(models.NewCategoryModel().GetTable(tid)).ID(aid).Incr("visit_count").Exec()
	//	if err != nil {
	//		logger.Error("无法更新点击数据", err)
	//		return
	//	}
	//	if affe, _ := res.RowsAffected(); affe > 0 {
	//		c.Ctx().SetCookie(clickCache, "1", 0)
	//	}
	//}
}

func (c *IndexController) GetClick() {
	//aid, _ := c.Ctx().GetInt64("aid")
	//tid, _ := c.Ctx().GetInt64("tid")
	//if aid < 1 || tid < 1 {
	//	c.Ctx().Abort(http.StatusNotFound)
	//	return
	//}
	//fmt.Println(di.MustGet("orm").(*xorm.Engine).Table(models.NewCategoryModel().GetTable(tid)).ID(aid).Select("visit_count").QueryString())
}

func getOrmSess(model *tables.DocumentModel) *xorm.Session {
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(controllers.GetTableName(model.Table)).Where("status = 1").Where("deleted_time IS NULL")
}

func (c *IndexController) setTemplateData() {
	host := c.Ctx().Request().Host
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
		return fmt.Sprintf("//%s/%s/%d.html", host, urlPrefix, iaid)
	}
	c.Ctx().Set("detail_url", detailUrl)
	if tid, _ := c.Param().GetInt64("tid"); tid <= 0 {
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
		p,_ := c.Ctx().Params().GetFloat64("page", 1)
		c.Ctx().Render().ViewData("page", p)
	}
	c.Ctx().Render().ViewData("detail_url", detailUrl)
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

func getCategoryPos(tid int64) string {
	var position []string
	var res string
	var arr []tables.Category
	key := fmt.Sprintf(controllers.CacheCategoryPosPrefix, tid)
	icache := di.MustGet(controllers.ServiceICache).(cache.AbstractCache)
	err := icache.GetWithUnmarshal(key, &res)
	if err != nil {
		m := models.NewCategoryModel()
		arr = m.GetPosArr(tid)
		for _, cat := range arr {
			if cat.Type != 2 {
				position = append(position, "<a href='"+m.GetUrlPrefix(cat.Catid)+"'>"+cat.Catname+"</a>")
			} else {
				position = append(position, "<a href='"+cat.Url+"'>"+cat.Catname+"</a>")
			}
		}
		if len(arr) > 0 {
			res = strings.Join(position, " > ")
			icache.SetWithMarshal(key, res)
		}
	}
	return res
}

func GetStaticFile(filename string) string {
	setting, _ := config.SiteConfig()
	if setting["SITE_STATIC_PAGE_DIR"] == "" {
		setting["SITE_STATIC_PAGE_DIR"] = "resources/html"
	}
	return filepath.Join(setting["SITE_STATIC_PAGE_DIR"], filename)
}
