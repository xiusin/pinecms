package frontend

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/contracts"
	"github.com/xiusin/pine/render/engine/pjet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DetailPageCacheData struct {
	Article     map[string]string
	PrevArticle map[string]string
	NextArticle map[string]string
	Category    *tables.Category
}

func (c *IndexController) Detail(pathname string) {
	c.setTemplateData()
	pageFilePath := GetStaticFile(pathname)
	aid, _ := c.Ctx().Params().GetInt64("aid")
	tid, _ := c.Ctx().Params().GetInt64("tid")
	var err error
	if tid < 1 || aid < 1 {
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	// 直接读缓存
	cacher := pine.Make(controllers.ServiceICache).(contracts.Cache)
	cacheKey := fmt.Sprintf(controllers.CacheCategoryContentPrefix, tid, aid)

	var cacheData DetailPageCacheData
	err = cacher.GetWithUnmarshal(cacheKey, &cacheData)

	m := models.NewCategoryModel()
	var category *tables.Category
	var article map[string]string
	var prevArticle map[string]string
	var nextArticle map[string]string

	if err == nil && cacheData.Article != nil { // Cache hit
		article = cacheData.Article
		prevArticle = cacheData.PrevArticle
		nextArticle = cacheData.NextArticle
		category = cacheData.Category
	} else { // Cache miss
		category, err = m.GetCategoryFByIdForBE(tid)
		if err != nil {
			pine.Logger().Error(err.Error())
			c.Ctx().Abort(http.StatusNotFound)
			return
		}
		if category.Model.Enabled == 0 {
			pine.Logger().Warn("模型内容已被禁止查看")
			c.Ctx().Abort(404)
			return
		}

		sess := getOrmSess(category.Model)

		articleResult, err := sess.Clone().Where("id = ? AND catid = ?", aid, tid).Limit(1).QueryString()
		if err != nil || len(articleResult) == 0 {
			pine.Logger().Error(fmt.Sprintf("读取模型数据表:%s 错误: %s", category.Model.Table, err))
			c.Ctx().Abort(http.StatusNotFound)
			return
		}
		article = articleResult[0]
		article["typename"] = category.Catname
		posArr, _ := m.GetPosArr(tid)
		article["typelink"] = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(posArr))
		article["click"] = article["visit_count"]

		prevArticleResult, _ := sess.Clone().Where("id < ? AND catid = ?", aid, tid).Desc("id").Limit(1).QueryString()
		if len(prevArticleResult) > 0 {
			prevArticle = prevArticleResult[0]
		}

		nextArticleResult, _ := sess.Clone().Where("id > ? AND catid = ?", aid, tid).Asc("id").Limit(1).QueryString()
		if len(nextArticleResult) > 0 {
			nextArticle = nextArticleResult[0]
		}

		// Store everything in cache
		cacheData = DetailPageCacheData{
			Article:     article,
			PrevArticle: prevArticle,
			NextArticle: nextArticle,
			Category:    category,
		}
		_ = cacher.SetWithMarshal(cacheKey, &cacheData)
	}

	// Now that we have all data, render the template
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
		pine.Logger().Error(err.Error())
		c.Ctx().Abort(http.StatusNotFound)
		return
	}
	defer f.Close()
	pineJet := pine.Make(controllers.ServiceJetEngine).(*pjet.PineJet)
	temp, err := pineJet.GetTemplate(template(tpl))
	if err != nil {
		pine.Logger().Error(err.Error())
		c.Ctx().Abort(http.StatusNotFound)
		return
	}

	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), struct {
		Field       map[string]string
		PrevArticle map[string]string
		NextArticle map[string]string
		TypeID      int64
		ArtID       int64
	}{
		Field:       article,
		PrevArticle: prevArticle,
		NextArticle: nextArticle,
		TypeID:      tid,
		ArtID:       aid,
	})
	if err != nil {
		pine.Logger().Error(err.Error())
		c.Ctx().Abort(http.StatusInternalServerError)
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().WriteHTMLBytes(data)
}
