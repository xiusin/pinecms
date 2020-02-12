package backend

import (
	"encoding/json"
	"html/template"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"
)

type ConfigItem map[string]interface{}

type ConfigStruct map[string]ConfigItem

type SettingController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Cache   cache.ICache
	Session *sessions.Session
}

func (c *SettingController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/setting/site", "Site")
}

func (c *SettingController) Site() {
	if c.Ctx.Method() == "POST" {
		var setting []*tables.IriscmsSetting
		act := c.Ctx.URLParam("dosubmit")
		var setval []ConfigItem
		if act == "" {
			if err := c.Orm.Asc("listorder").Find(&setting); err != nil {
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
			c.Ctx.JSON(result)
			return
		}
		post := c.Ctx.FormValues()
		for k, v := range post {
			if k == "dosubmit" || len(v) == 0 {
				continue
			}
			c.Orm.Table(new(tables.IriscmsSetting)).Where("`key`=?", k).MustCols("value").Update(&tables.IriscmsSetting{Value: v[0]})
		}
		c.Cache.Delete(controllers.CacheSetting)
		helper.Ajax("更新配置信息成功", 0, c.Ctx)
		return
	}
	menuid, err := c.Ctx.URLParamInt64("menuid")
	if err != nil {
		menuid = 0
	}
	currentpos := models.NewMenuModel(c.Orm).CurrentPos(menuid)
	grid := helper.Propertygrid("setting_site_propertygrid", helper.EasyuiOptions{
		"title":   currentpos,
		"url":     "/b/setting/site?grid=propertygrid",
		"toolbar": "setting_site_propertygrid_toolbar",
	})
	c.Ctx.ViewData("grid", template.HTML(grid))
	c.Ctx.View("backend/setting_site.html")
}
