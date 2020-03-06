package backend

import (
	"encoding/json"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"html/template"
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
	b.ANY( "/setting/site", "Site")
}

func (c *SettingController) Site(iCache cache.ICache) {
	if c.Ctx().IsPost() {
		var setting []*tables.IriscmsSetting
		act := c.Ctx().URLParam("dosubmit")
		var setval []ConfigItem
		if act == "" {
			if err := c.Ctx().Value("orm").(*xorm.Engine).Asc("listorder").Find(&setting); err != nil {
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
			c.Ctx().Value("orm").(*xorm.Engine).Table(new(tables.IriscmsSetting)).Where("`key`=?", k).MustCols("value").Update(&tables.IriscmsSetting{Value: v[0]})
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
