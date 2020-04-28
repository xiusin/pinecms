package frontend

import (
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (c *IndexController) Detail(pathname string) {
	c.setTemplateData()
	pageFilePath := GetStaticFile(pathname)
	aid, _ := c.Param().GetInt64("aid")
	tid, _ := c.Param().GetInt64("tid")
	var err error
	if tid < 1 || aid < 1 {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}

	// 直接读缓存
	cacher := pine.Make("cache.AbstractCache").(cache.AbstractCache)
	cacheKey := fmt.Sprintf(controllers.CacheCategoryContentPrefix, tid, aid)
	var article = map[string]string{}
	cacher.GetWithUnmarshal(cacheKey, &article)
	if len(article) == 0 {

	}
	m := models.NewCategoryModel()
	category, err := m.GetCategoryFByIdForBE(tid)
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	if category.Model.Enabled == 0 {
		pine.Logger().Warning("模型内容已被禁止查看")
		c.Ctx().Abort(404)
		return
	}
	if len(article) == 0 {
		sess := getOrmSess(category.Model).Where("id = ?", aid).Where("catid = ?", tid).Limit(1)
		result, err := sess.QueryString()
		if err != nil || len(result) == 0 {
			pine.Logger().Errorf("读取模型数据表:%s 错误: %s", category.Model.Table, err)
			c.Ctx().Abort(http.StatusNotFound)
			return
		}
		article = result[0]
		article["typename"] = category.Catname
		article["typelink"] = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(m.GetPosArr(tid)))
		article["click"] = article["visit_count"]
		cacher.SetWithMarshal(cacheKey, &article)
	}
	detailUrlFunc := c.Ctx().Value("detail_url").(func(string, ...string) string)
	tpl := "article_" + category.Model.Table + ".jet"
	if len(category.Model.FeTplDetail) > 0 {
		tpl = category.Model.FeTplDetail
	}
	if len(category.DetailTpl) > 0 {
		tpl = category.DetailTpl
	}
	_ = os.MkdirAll(filepath.Dir(pageFilePath), os.ModePerm)
	f, err := os.OpenFile(pageFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	defer f.Close()
	pineJet := pine.Make(controllers.ServiceJetEngine).(*jet.PineJet)
	temp, err := pineJet.GetTemplate(template(tpl))
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusNotFound)
		return
	}

	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), struct {
		Field    map[string]string
		TypeID   int64
		ArtID    int64
		PrevNext func(string, tpl string) string
	}{
		Field:  article,
		TypeID: tid,
		ArtID:  aid,
		PrevNext: func(conf string, tpl string) string {
			var str string
			confs := strings.Split(conf, ",")
			hasNext, hasPrev := false, false
			for _, v := range confs {
				if !hasNext || !hasPrev {
					var data []map[string]string
					var titleFlag string
					if v == "prev" || v == "pre" {
						titleFlag, hasPrev = "上一篇: ", true
						data, _ = getOrmSess(category.Model).Where("id < ?", aid).Desc("id").Limit(1).QueryString()
					} else if v == "next" {
						titleFlag, hasNext = "下一篇: ", true
						data, _ = getOrmSess(category.Model).Where("id > ?", aid).Desc("id").Limit(1).QueryString()
					}
					if len(data) > 0 {
						if tpl == "" {
							str += "<p><a href='" + detailUrlFunc(data[0]["id"], data[0]["catid"]) + "'>" + titleFlag + ": " + data[0]["title"] + "</a></p>"
						} else {
							str += strings.ReplaceAll(tpl, "~arturl~", detailUrlFunc(data[0]["id"], data[0]["catid"]))
							str = strings.ReplaceAll(str, "~title~", data[0]["title"])
						}
					}
				}
			}
			return str
		}})
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusInternalServerError)
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().Writer().Write(data)
}
