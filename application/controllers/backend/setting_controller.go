package backend

import (
	"github.com/go-xorm/xorm"
	"iriscms/common/helper"
	"html/template"
	"iriscms/application/models/tables"
	"strings"
	"iriscms/application/models"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
)

type ConfigItem map[string]string

type ConfigStruct map[string]ConfigItem

type SettingController struct {
	Ctx iris.Context
	Orm *xorm.Engine
	Session *sessions.Session
}



func (c *SettingController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY","/setting/site","Site")
	b.Handle("ANY","/setting/site-default", "SiteDefault")
}

var SiteConfig ConfigStruct = ConfigStruct{
	"SITE_TITLE": {
		"name":    "站点标题",
		"group":   "前台设置",
		"editor":  "text",
		"default": "iriscms项目",
	},
	"SITE_KEYWORDS": {
		"name":    "关键字",
		"group":   "前台设置",
		"editor":  "text",
		"default": "iriscms项目",
	},
	"SITE_DESCRIPTION": {
		"name":    "描述",
		"group":   "前台设置",
		"editor":  "textarea",
		"default": "iriscms项目",
	},
	"SITE_ICP": {
		"name":    "备案号",
		"group":   "前台设置",
		"editor":  "text",
		"default": "",
	},

	"DATAGRID_PAGE_SIZE": {
		"name":    "列表默认分页数",
		"group":   "后台设置",
		"editor":  "numberbox",
		"default": "25",
	},
	"EMAIL_SMTP": {
		"name":    "SMTP",
		"group":   "邮箱设置",
		"editor":  "text",
		"default": "",
	},
	"EMAIL_PORT": {
		"name":    "端口",
		"group":   "邮箱设置",
		"editor":  "numberbox",
		"default": "25",
	},
	"EMAIL_EMAIL": {
		"name":    "邮箱地址",
		"group":   "邮箱设置",
		"editor":  "text",
		"default": "",
	},
	"EMAIL_USER": {
		"name":    "用户名",
		"group":   "邮箱设置",
		"editor":  "text",
		"default": "",
	},
	"EMAIL_PWD": {
		"name":    "密码",
		"group":   "邮箱设置",
		"editor":  "text",
		"default": "",
	},
}

//系统配置 -> 站点配置
func (this *SettingController) Site() {
	if this.Ctx.Method() == "POST" {
		setting := []tables.IriscmsSetting{}
		act := this.Ctx.URLParam("dosubmit")
		setval := []ConfigItem{}
		if act == "" {
			this.Orm.Find(&setting)
			var keys []string = []string{}
			if len(setting) != 0 {
				for _, v := range setting {
					if _, ok := SiteConfig[v.Key]; !ok {
						continue
					}
					keys = append(keys, v.Key)
					setval = append(setval, ConfigItem{
						"name":    SiteConfig[v.Key]["name"],
						"key":     v.Key,
						"group":   SiteConfig[v.Key]["group"],
						"editor":  SiteConfig[v.Key]["editor"],
						"default": SiteConfig[v.Key]["default"],
						"value":   v.Value,
					})
				}
			}
			var keysStr string
			if len(keys) > 0 {
				keysStr = strings.Join(keys, "|")
			} else {
				keysStr = ""
			}

			//没有值的配置项自动合并
			for k, v := range SiteConfig {
				if keysStr != "" && strings.Contains(keysStr, k) {
					continue
				}
				setval = append(setval, ConfigItem{
					"key":     k,
					"name":    v["name"],
					"group":   v["group"],
					"editor":  v["editor"],
					"default": v["default"],
					"value":   v["default"],
				})
			}

			result := map[string]interface{}{
				"rows":  setval,
				"total": len(setval),
			}
			this.Ctx.JSON(result)
			return
		}
		post := this.Ctx.FormValues()

		flag := false
		for k, v := range post {
			if k == "dosubmit" || len(v) == 0 {
				continue
			}
			//更新数据
			setting := tables.IriscmsSetting{Key: k}
			bol, _ := this.Orm.Get(&setting)	//逐个查找,判断添加还是修改配置
			if bol {
				res, _ := this.Orm.Table(new(tables.IriscmsSetting)).Where("`key`=?", k).Update(&tables.IriscmsSetting{Value: v[0]})
				if res > 0 && flag == false {
					flag = true
				}
			} else {
				setting.Value = v[0]
				res, _ := this.Orm.Insert(setting)
				if res > 0 && flag == false {
					flag = true
				}
			}

		}
		if flag {
			helper.Ajax("更新配置信息成功", 0, this.Ctx)
		} else {
			helper.Ajax("没有更新任何配置", 1, this.Ctx)
		}
		return
	}
	menuid, err := this.Ctx.URLParamInt64("menuid")
	if err != nil {
		menuid = 0
	}
	currentpos := models.NewMenuModel(this.Orm).CurrentPos(menuid)
	grid := helper.Propertygrid("setting_site_propertygrid", helper.EasyuiOptions{
		"title":   currentpos,
		"url":     "/b/setting/site?grid=propertygrid",
		"toolbar": "setting_site_propertygrid_toolbar",
	})
	this.Ctx.ViewData("grid",template.HTML(grid))
	this.Ctx.View("backend/setting_site.html")
}

//站点配置恢复默认
func (this *SettingController) SiteDefault() {
	ok, _ := this.Orm.Where("1").Delete(new(tables.IriscmsSetting))
	if ok == 0 {
		helper.Ajax("操作失败", 1, this.Ctx)
	} else {
		helper.Ajax("操作成功", 0, this.Ctx)
	}
}
