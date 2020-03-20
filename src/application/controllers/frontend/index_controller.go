package frontend

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/config"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type IndexController struct {
	pine.Controller
}

func (c *IndexController) RegisterRoute(b pine.IRouterWrapper) {
	b.GET("/", "Index")
	b.GET("/list", "List")
	b.GET("/detail", "Detail")
	b.GET("/page", "Page")
	b.GET("/click", "Click")
}

func (c *IndexController) Index() {
	// todo 渲染到静态文件html ,path 调整为静态地址
	// todo writer 为文件对象处理
	c.Ctx().Render().HTML(template("index.jet"))
}

func (c *IndexController) List() {
	tid, err := c.Ctx().GetInt64("tid")
	if err != nil || tid < 1 {
		c.Ctx().Abort(404)
		return
	}
	category, err := models.NewCategoryModel().GetCategory(tid)
	if err != nil {
		c.Ctx().Abort(404)
		return
	}
	var globalSize = 10
	c.ViewData("field", category)
	c.ViewData("typeid", tid)
	c.ViewData("__typeid", tid)
	c.ViewData("list", func(pagesize int, orderby string) []map[string]string {
		if orderby == "" {
			orderby = "listorder desc"
		}
		if pagesize > 0 {
			globalSize = pagesize
		} else {
			pagesize = globalSize
		}
		list, err := getOrmSess().Limit(pagesize).Where("catid = ?", tid).
			Where("deleted_time IS NULL").Where("status = 1").OrderBy(orderby).QueryString()
		if err != nil {
			panic(err)
		}
		if list == nil {
			list = []map[string]string{}
		}
		return list
	})

	c.ViewData("pagelist", func(listsize int) string {
		total, _ := getOrmSess().Where("catid = ?", tid).
			Where("deleted_time IS NULL").Where("status = 1").Count()
		// 计算页码
		totalPage := int(math.Ceil(float64(total) / float64(globalSize)))
		pagenum, _ := c.Ctx().GetInt("page")
		if pagenum < 1 {
			pagenum = 1
		} else if pagenum > totalPage {
			pagenum = totalPage
		}
		return fmt.Sprintf("总记录数: %d, 总页数: %d, 当前页数: %d, 分页条目数为: %d", total, totalPage, pagenum, globalSize)
	})

	if category.ListTpl == "" {
		category.ListTpl = "list_article.jet"
	}
	c.Render().HTML(template(category.ListTpl))
}

func (c *IndexController) Detail(icache cache.ICache) {
	var err error
	var category tables.Category
	aid, _ := c.Ctx().GetInt64("aid")
	tid, _ := c.Ctx().GetInt64("tid")
	if tid < 1 || aid < 1 {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	categoryCacheKey := fmt.Sprintf(controllers.CacheCategoryInfoPrefix, tid)
	// 获取模型
	data, _ := icache.Get(categoryCacheKey)

	if len(data) > 0 {
		json.Unmarshal(data, &category)
	}
	catmodel := models.NewCategoryModel()
	if category.Catid == 0 {
		category, err = catmodel.GetCategory(tid)
		if err != nil || category.ModelId == 0 {
			c.Ctx().Abort(http.StatusNotFound)
			return
		}
		data, _ := json.Marshal(&category)
		icache.Set(categoryCacheKey, data)
	}
	tableName := models.NewDocumentModel().GetTableName(category.ModelId)
	rest, err := getOrmSess(tableName).Where("id = ?", aid).Where("catid = ?", tid).Where("status = 1").Where("deleted_time IS NULL").Limit(1).QueryString()
	if err != nil {
		pine.Logger().Errorf("查找tableName:%s错误: %s",tableName, err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	article := rest[0]
	article["catname"] = category.Catname
	article["position"] = getCategoryPos(tid)

	c.ViewData("field", article)
	c.ViewData("__typeid", tid)
	c.ViewData("prenext", func() string {
		var str string
		// aid 根据aid 读取上一篇和下一篇
		pre, _ := getOrmSess(tableName).Where("id < ?", aid).Where("status = 1").Where("deleted_time IS NULL").Desc("id").Limit(1).QueryString()
		if len(pre) > 0 {
			str += "<a href='/detail?aid=" + pre[0]["id"] + "'>上一篇：" + pre[0]["title"] + "</a>"
		}
		next, _ := getOrmSess(tableName).Where("id > ?", aid).Where("status = 1").Where("deleted_time IS NULL").Asc("id").Limit(1).QueryString()
		if len(next) > 0 {
			str += "<a href='/detail?aid=" + next[0]["id"] + "'>下一篇: " + next[0]["title"] + "</a>"
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
		sess := getOrmSess(tableName).Where("id <> ?", article["id"]).Where("status = 1").Where("deleted_time IS NULL")
		for _, kw := range kws {
			splitKeywords := strings.Split(kw, ",")
			for _, keyword := range splitKeywords {
				conds = append(conds, "CONCAT(keywords,' ',title, ' ', tags) LIKE ?")
				keywords = append(keywords, "%"+strings.Trim(keyword, "")+"%")
			}
		}
		sess.And(strings.Join(conds, " OR "), keywords...)
		articles, _ := sess.Limit(row).Desc("id").QueryString()
		return articles
	})
	c.Ctx().Render().HTML(template("detail.jet"))
}

func (c *IndexController) Page() {
	tid, _ := c.Ctx().GetInt64("tid")
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
	c.View(template("page.jet"))
}

func (c *IndexController) Click() {
	aid, _ := c.Ctx().GetInt64("aid")
	tid, _ := c.Ctx().GetInt64("tid")
	if aid < 1 || tid < 1 {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	clickCache := fmt.Sprintf("click_%d_%d", tid, aid)
	info := c.Ctx().GetCookie(clickCache)
	if len(info) == 0 {
		res, err := di.MustGet("orm").(*xorm.Engine).Table(models.NewCategoryModel().GetTable(tid)).ID(aid).Incr("visit_count").Exec()
		if err != nil {
			logger.Error("无法更新点击数据", err)
			return
		}
		if affe, _ := res.RowsAffected(); affe > 0 {
			c.Ctx().SetCookie(clickCache, "1", 0)
		}
	}
}

func (c *IndexController) GetClick() {
	aid, _ := c.Ctx().GetInt64("aid")
	tid, _ := c.Ctx().GetInt64("tid")
	if aid < 1 || tid < 1 {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	fmt.Println(di.MustGet("orm").(*xorm.Engine).Table(models.NewCategoryModel().GetTable(tid)).ID(aid).Select("visit_count").QueryString())
}

func getOrmSess(tableName ...string) *xorm.Session {
	if len(tableName) == 0 {
		tableName = append(tableName, "articles")
	}
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(controllers.GetTableName(tableName[0]))
}

func template(tpl string) string {
	conf := di.MustGet(controllers.ServiceConfig).(*config.Config)
	return filepath.Join(conf.View.Theme, tpl)
}

func getCategoryPos(tid int64) string {
	var position []string
	var res string
	var arr []tables.Category
	key := fmt.Sprintf(controllers.CacheCategoryPosPrefix, tid)
	icache := di.MustGet(controllers.ServiceICache).(cache.ICache)
	data, err := icache.Get(key)
	if len(data) > 0 {
		err = json.Unmarshal(data, &res)
	}
	if err != nil {
		arr = models.NewCategoryModel().GetPosArr(tid)
		for _, cat := range arr {
			if cat.Type == 0 {
				position = append(position, "<a href='/list?tid="+strconv.Itoa(int(cat.Catid))+"'>"+cat.Catname+"</a>")
			} else if cat.Type == 1 {
				position = append(position, "<a href='/page?tid="+strconv.Itoa(int(cat.Catid))+"'>"+cat.Catname+"</a>")
			} else {
				position = append(position, "<a href='"+cat.Url+"'>"+cat.Catname+"</a>")
			}
		}
		if len(arr) > 0 {
			res = strings.Join(position, " > ")
			icache.Set(key, []byte(res))
		}
	}
	return res
}
