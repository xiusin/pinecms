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
	"sync"
)

type ConfigItem map[string]interface{}

type ConfigStruct map[string]ConfigItem

type SettingController struct {
	Ctx iris.Context
	Orm *xorm.Engine
	Session *sessions.Session
}

var settingWg sync.WaitGroup

func (c *SettingController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY","/setting/site","Site")
	b.Handle("ANY","/setting/site-default", "SiteDefault")
}

var SiteConfig = ConfigStruct{
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
	"SITE_OPEN": {
		"name":    "站点开启",
		"group":   "前台设置",
		"editor":  ConfigItem{
			"type":"checkbox",
			"options": map[string]interface{}{
				"on":"开启",
				"off": "关闭",
			},
		},
		"default": "开启",
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

	"WX_APPID": {
		"name":    "APPID",
		"group":   "微信配置",
		"editor":  "text",
		"default": "",
	},
	"WX_APPSECRET":{
		"name":    "APPSECTET",
		"group":   "微信配置",
		"editor":  "text",
		"default": "",
	},

	"WX_TOKEN": {
		"name":    "TOKEN",
		"group":   "微信配置",
		"editor":  "text",
		"default": "",
	},
	"WX_AESKEY": {
		"name":    "AESKEY",
		"group":   "微信配置",
		"editor":  "text",
		"default": "",
	},


	"OSS_ENDPOINT": {
		"name":    "ENDPOINT",
		"group":   "OSS存储配置",
		"editor":  "text",
		"default": "",
	},

	"OSS_KEYID": {
		"name":    "KEYID",
		"group":   "OSS存储配置",
		"editor":  "text",
		"default": "",
	},

	"OSS_KEYSECRET": {
		"name":    "SECRET",
		"group":   "OSS存储配置",
		"editor":  "text",
		"default": "",
	},

	"OSS_BUCKETNAME": {
		"name":    "BUCKETNAME",
		"group":   "OSS存储配置",
		"editor":  "text",
		"default": "",
	},

	"OSS_HOST": {
		"name":    "HOST",
		"group":   "OSS存储配置",
		"editor":  "text",
		"default": "",
	},

}

//系统配置 -> 站点配置
func (this *SettingController) Site() {
	if this.Ctx.Method() == "POST" {
		var setting []tables.IriscmsSetting
		act := this.Ctx.URLParam("dosubmit")
		var setval []ConfigItem
		if act == "" {
			this.Orm.Find(&setting)
			var keys []string
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

		for k, v := range post {
			if k == "dosubmit" || len(v) == 0 {
				continue
			}
			//更新数据
			go func(k,v string) {
				settingWg.Add(1)
				setting := tables.IriscmsSetting{Key: k}
				bol, _ := this.Orm.Get(&setting)	//逐个查找,判断添加还是修改配置
				if bol {
					this.Orm.Table(new(tables.IriscmsSetting)).Where("`key`=?", k).Update(&tables.IriscmsSetting{Value: v})
				} else {
					setting.Value = v
					this.Orm.Insert(setting)
				}
				defer settingWg.Done()
			}(k,v[0])
		}
		settingWg.Wait()
		helper.Ajax("更新配置信息成功", 0, this.Ctx)
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
