package backend

import (
	"encoding/json"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/application/plugins"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

type PluginController struct {
	BaseController
}

func (c *PluginController) Construct() {
	c.Table = &tables.Plugin{}
	c.Entries = &[]*tables.Plugin{}
	c.ApiEntityName = "插件"
	c.Group = "插件管理"
	c.SubGroup = "插件管理"
	c.BaseController.Construct()
	c.OpAfter = c.after
}

func (c *PluginController) after(act int, _ interface{}) error {
	if act == OpList {
		// todo 对比插件信息, 展示注册页面信息功能
		mgr := plugins.PluginMgr()

		entities := c.Entries.(*[]*tables.Plugin)

		for _, plugin := range *entities {
			if len(plugin.Config) > 0 {
				plugin.Status = 1
			}
		}

		return mgr.Iter(func(name string, pluginEntity plugins.PluginIntf) error {
			var view []interface{}
			if err := json.Unmarshal([]byte(pluginEntity.View()), &view); err != nil {
				return nil
			}
			pluginDb := &tables.Plugin{
				Name:        pluginEntity.Name(),
				Author:      pluginEntity.Author(),
				Contact:     pluginEntity.Contact(),
				Description: pluginEntity.Description(),
				Version:     pluginEntity.Version(),
				Sign:        pluginEntity.Sign(),
				View:        interface{}(view),
				Path:        name,
				CreatedAt:   tables.LocalTime(time.Now()),
			}
			existData := &tables.Plugin{}
			_, _ = c.Orm.Where("sign = ?", pluginDb.Sign).Get(existData)
			var err error
			var rest int64
			var firstInstall bool
			if existData.Id == 0 {
				rest, err = c.Orm.InsertOne(pluginDb)
				firstInstall = true
			} else {
				rest, err = c.Orm.Where("sign = ?", pluginDb.Sign).Update(existData)
			}
			if rest > 0 {
				if firstInstall {
					pluginEntity.Install()
				}
			} else {
				return err
			}
			return nil
		})

	}
	return nil
}

func (c *PluginController) PostEnable() {
	sign := c.Input().GetStringBytes("sign")
	if sign == nil || len(sign) == 0 {
		helper.Ajax("请传入标识参数", 1, c.Ctx())
		return
	}
	enable := c.Input().GetBool("enable")
	res, _ := c.Orm.Where("sign = ?", string(sign)).Cols("enable").Update(&tables.Plugin{
		Enable:    enable,
		UpdatedAt: tables.LocalTime(time.Now()),
	})
	if res > 0 {
		helper.Ajax("success", 0, c.Ctx())
	} else {
		helper.Ajax("failed", 1, c.Ctx())
	}
}

func (c *PluginController) GetConfig() {
	plugin := &tables.Plugin{}
	c.Orm.Where("sign = ?", c.Ctx().GetString("sign")).Get(plugin)
	helper.Ajax(plugin.Config, 0, c.Ctx())
}

func (c *PluginController) PostConfig() {
	plugin := &tables.Plugin{}
	if err := c.Ctx().BindJSON(plugin); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	plugin.UpdatedAt = tables.LocalTime(time.Now())
	res, _ := c.Orm.Where("sign = ?", plugin.Sign).Update(plugin)
	if res > 0 {
		helper.Ajax("success", 0, c.Ctx())
	} else {
		helper.Ajax("failed", 1, c.Ctx())
	}
}
