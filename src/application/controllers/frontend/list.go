package frontend

import (
	"github.com/valyala/fasthttp"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/pjet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (c *IndexController) List(pageFilePath string) {
	c.setTemplateData()
	pageFilePath = GetStaticFile(pageFilePath)
	queryTid, _ := c.Ctx().GetInt64("tid")
	tid, _ := c.Ctx().Params().GetInt64("tid", queryTid)
	if tid < 1 {
		c.Ctx().Abort(404)
		return
	}
	category, err := models.NewCategoryModel().GetCategoryFByIdForBE(tid)
	if err != nil || category.Model.Enabled == 0 {
		if err == nil {
			pine.Logger().Error("模型禁用,无可查看", c.Ctx().Path())
		} else {
			pine.Logger().Error(err)
		}
		c.Ctx().Abort(404)
		return
	}
	page, _ := c.Ctx().Params().GetInt("page", 1)
	if page < 1 {
		page = 1
	}
	total, _ := getOrmSess(category.Model).In("id", models.NewCategoryModel().GetNextCategoryOnlyCatids(tid, true)).Count()
	tpl := "list_" + category.Model.Table + ".jet" // default tpl
	if len(category.Model.FeTplList) > 0 {         // model tpl
		tpl = category.Model.FeTplList
	}
	if len(category.ListTpl) > 0 { // category tpl
		tpl = category.ListTpl
	}
	_ = os.MkdirAll(filepath.Dir(pageFilePath), os.ModePerm)
	f, err := os.Create(pageFilePath)
	if err != nil {
		c.Ctx().Abort(fasthttp.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	jet := pine.Make(controllers.ServiceJetEngine).(*pjet.PineJet)
	temp, err := jet.GetTemplate(template(tpl))
	if err != nil {
		c.Ctx().Abort(fasthttp.StatusInternalServerError, err.Error())
		return
	}

	err = temp.Execute(f, viewDataToJetMap(c.Render().GetViewData()), struct {
		Field     *tables.Category
		TypeID    int64
		ArtCount  int64
		ModelName string
		QP        map[string]interface{}
		PageNum   int64
	}{
		Field:     category,
		TypeID:    tid,
		ArtCount:  total,
		PageNum:   int64(page),
		ModelName: category.Model.Table,
		QP:        c.Ctx().All(),
	})
	if err != nil {
		c.Ctx().Abort(fasthttp.StatusInternalServerError, err.Error())
		return
	}
	data, _ := ioutil.ReadFile(pageFilePath)
	c.Ctx().WriteHTMLBytes(data)
}
