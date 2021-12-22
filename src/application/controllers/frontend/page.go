package frontend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (c *IndexController) Page(pathname string) {
	c.setTemplateData()
	pageFilePath := GetStaticFile(pathname)
	tid, _ := c.Ctx().Params().GetInt64("tid")
	if tid < 1 {
		c.Ctx().Abort(404, "tid failed")
		return
	}
	category, err := models.NewCategoryModel().GetCategoryFByIdForBE(tid)
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(404)
		return
	}
	page := category.Page
	if page == nil {
		page = &tables.Page{Title: category.Catname}
	}
	if page.Keywords == "" {
		page.Keywords = category.Keywords
	}
	if len(page.Description) > 0 {
		page.Description = category.Description
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
	tpl := "page.jet"
	if len(category.DetailTpl) > 0 {
		tpl = category.DetailTpl
	}
	temp, err := pineJet.GetTemplate(template(tpl))
	if err != nil {
		c.Ctx().WriteString(err.Error())
		return
	}
	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), struct {
		Field       *tables.Page // 单页信息
		Position    string
		TypeID      int64
		ArtID       int64
		TopCategory *tables.Category // 顶级栏目信息
	}{
		Field:  page,
		TypeID: tid,
	})
	if err != nil {
		c.Ctx().WriteString(err.Error())
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().WriteHTMLBytes(data)
}
