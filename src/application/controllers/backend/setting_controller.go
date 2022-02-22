package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/common/message"
	"github.com/xiusin/pinecms/src/config"
	"xorm.io/xorm"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Construct() {
	c.Table = &tables.Setting{}
	c.Entries = &[]*tables.Setting{}
	c.Group = "系统配置"
	c.SubGroup = "配置模块"
	c.ApiEntityName = "配置"

	c.BaseController.Construct()

	c.apiEntities = map[string]apidoc.Entity{
		"list":   {Title: "配置列表", Desc: "查询系统指定分组配置信息列表"},
		"add":    {Title: "新增配置", Desc: "新增一个新的配置"},
		"edit":   {Title: "修改配置", Desc: "编辑已存在的配置项"},
		"del":    {Title: "删除配置", Desc: "删除一个或多个配置项"},
		"info":   {Title: "配置详情", Desc: "获取指定配置详情信息"},
		"groups": {Title: "配置分组列表", Desc: "获取所有配置分组列表"},
	}
	c.setApiEntity()

	c.OpBefore = c.before
	c.OpAfter = c.after
}

func (c *SettingController) before(act int, params interface{}) error {
	if act == OpList {
		pars := c.Input().Get("params")
		if pars == nil {
			panic("必须选择分组")
		}
		params.(*xorm.Session).Where("`group` = ?", pars.(map[string]interface{})["group"])
	}
	return nil
}

func (c *SettingController) after(act int, _ interface{}) error {
	if act == OpEdit {
		helper.AbstractCache().Delete(controllers.CacheSetting)
		config.SiteConfig()
	}
	return nil
}

// PostGroups 获取新分组
func (c *SettingController) PostGroups() {
	var groups []tables.Setting
	c.Orm.Table(&tables.Setting{}).GroupBy("`group`").Find(&groups)
	var kvs = []KV{}

	for _, group := range groups {
		kvs = append(kvs, KV{Label: group.Group, Value: group.Group})
	}
	helper.Ajax(kvs, 0, c.Ctx())
}

func (c *SettingController) PostTest() {
	email, _ := c.Input().GetString("email")
	title, _ := c.Input().GetString("title")
	content, _ := c.Input().GetString("content")

	emailMessage := &message.EmailMessage{}
	if err := emailMessage.Init(); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	if err := emailMessage.Send([]string{email}, title, content); err != nil {
		helper.Ajax(err, 1, c.Ctx())
	} else {
		helper.Ajax("发送邮箱成功", 0, c.Ctx())
	}
}
