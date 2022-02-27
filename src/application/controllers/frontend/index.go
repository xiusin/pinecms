package frontend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/pjet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func (c *IndexController) Index() {
	c.setTemplateData()
	indexPage := "editor.tpl"
	pageFilePath := GetStaticFile(indexPage)
	os.MkdirAll(filepath.Dir(pageFilePath), os.ModePerm)
	f, err := os.OpenFile(pageFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		c.Logger().Error(err)
		return
	}
	defer f.Close()
	jet := pine.Make(controllers.ServiceJetEngine).(*pjet.PineJet)
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
	startTime := time.Now()
	c.Ctx().WriteHTMLBytes(data)
	pine.Logger().Debug("渲染模板总耗时", time.Now().Sub(startTime))
}
