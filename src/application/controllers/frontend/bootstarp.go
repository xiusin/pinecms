package frontend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"xorm.io/xorm"
)

func (c *IndexController) Bootstrap(orm *xorm.Engine, cacheHandler cache.AbstractCache) {
	conf, _ := config.SiteConfig()
	// todo 开启前端资源缓存 304
	// todo 拦截存在静态文件的问题, 不过最好交给nginx等服务器转发
	c.statistics(cacheHandler)
	if conf["SITE_DEBUG"] == "" || conf["SITE_DEBUG"] == "关闭" {
		pageName := c.Ctx().Params().Get("pagename") // 必须包含.html, 在nginx要注意如果以/结尾的path需要追加index.html
		if pageName == "" {
			pageName = "/"
		}
		if strings.HasSuffix(pageName, "/") {
			pageName += "editor.tpl"
		}
		absFilePath := filepath.Join(conf["SITE_STATIC_PAGE_DIR"], pageName)
		if byts, err := ioutil.ReadFile(absFilePath); err != nil {
			pine.Logger().Error(err)
			c.Ctx().Abort(http.StatusNotFound)
		} else {
			c.Ctx().Render().ContentType(pine.ContentTypeHTML)
			pine.Logger().Print("render file ", absFilePath)
			_ = c.Ctx().Render().Bytes(byts)
		}
		return
	}
	pageName := strings.Trim(strings.ReplaceAll(c.Ctx().Params().Get("pagename"), "//", "/"), "/") // 必须包含.html
	switch pageName {
	case "editor.tpl", "":
		c.Index()
	default:
		urlPartials := strings.Split(pageName, "/")
		var last string
		var fileName string
		var isDetail bool
		if strings.HasSuffix(pageName, ".html") {
			last = urlPartials[len(urlPartials)-2]
			fileName = urlPartials[len(urlPartials)-1]
			// 分析页码
			if strings.HasPrefix(fileName, "index_") {
				fileInfo := strings.Split(fileName, "_") // index_2.html => 某个分类的第二页
				c.Ctx().Params().Set("page", strings.TrimSuffix(fileInfo[1], ".html"))
			} else if fileName != "editor.tpl" {
				isDetail = true
				c.Ctx().Params().Set("aid", strings.TrimSuffix(fileName, ".html")) // 设置文档ID
			}
		} else {
			last = urlPartials[len(urlPartials)-1] // 目录名
			fileName = "editor.tpl"
			pageName = filepath.Join(pageName, fileName)
			c.Ctx().Params().Set("page", "1")
		}
		cat := models.NewCategoryModel().GetWithDirForBE(last)
		if cat == nil {
			// 拆分出来想要的数据 page_{tid}
			if strings.HasPrefix(last, "page_") {
				c.Ctx().Params().Set("tid", strings.TrimPrefix(last, "page_"))
				c.Page(pageName)
				return
			} else {
				infos := strings.Split(last, "_") // 根据模型{model_table}_{tid}拆分信息
				modelTable := infos[0]
				model := models.NewDocumentModel().GetWithTableNameForBE(modelTable)
				if model != nil {
					c.Ctx().Params().Set("tid", infos[1])
					c.List(pageName)
					return
				}
			}
			c.Ctx().Abort(http.StatusNotFound)
			c.Logger().Debug("地址内容无法匹配", c.Ctx().Path())
			return
		}
		// 匹配所有内容
		prefix := models.NewCategoryModel().GetUrlPrefix(cat.Catid)
		if !strings.HasPrefix(pageName, prefix) {
			c.Logger().Debug("地址前缀无法匹配", c.Ctx().Path())
			c.Ctx().Abort(http.StatusNotFound)
			return
		}
		c.Ctx().Params().Set("tid", strconv.Itoa(int(cat.Catid)))
		if isDetail {
			c.Detail(pageName)
		} else if cat.Type == 0 {
			c.List(pageName)
		} else {
			c.Page(pageName)
		}
	}
}

// 统计数据
func (c *IndexController) statistics(cacheHandler cache.AbstractCache) {
	ck := "pinecms_statistics"
	val := c.Ctx().GetCookie(ck)

	today := time.Now().In(helper.GetLocation()).Format("01-02")
	if len(val) == 0 || val != today {
		c.Ctx().SetCookie(ck, today, 48*3600)
		var statistics = map[string]uint8{}
		_ = cacheHandler.GetWithUnmarshal(controllers.CacheStatistics, &statistics)
		if _, ok := statistics[today]; !ok {
			statistics[today] = 1
		} else {
			statistics[today] += 1
		}
		_ = cacheHandler.SetWithMarshal(controllers.CacheStatistics, &statistics)
		var defaultRefers = map[string]int{"baidu": 0, "google": 0, "so": 0, "bing": 0, "sougou": 0, "other": 0}
		var refers = map[string]int{}
		// 访问来源
		referer := string(c.Ctx().Referer())
		matchs := regexp.MustCompile("^https?://.*(baidu|google|so|bing|sougou).(com|cn)").FindAllStringSubmatch(referer, 1)
		_ = cacheHandler.GetWithUnmarshal(controllers.CacheRefer, &refers)
		if len(matchs) > 0 {
			if len(refers) == 0 {
				refers = defaultRefers
			}
			refers[matchs[0][1]] += 1
		} else {
			refers["other"] += 1
		}
		_ = cacheHandler.SetWithMarshal(controllers.CacheRefer, &refers)
	}
}
