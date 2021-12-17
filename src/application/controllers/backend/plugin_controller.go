package backend

import (
	"encoding/json"
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/application/plugins"
	"github.com/xiusin/pinecms/src/common/helper"
	"math/rand"
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
		// 合并安装插件和未安装插件
		mgr := plugins.PluginMgr()

		entities := c.Entries.(*[]*tables.Plugin)

		localPlugins := mgr.GetLocalPlugin()

		for _, plugin := range *entities {
			if len(plugin.Config) > 0 {
				plugin.Status = 1 // 改变插件状态
			}
			delete(localPlugins, plugin.Path) // 删除掉一些已经安装的插件信息
		}

		// 合并未安装的插件
		for plugin, conf := range localPlugins {
			item := &tables.Plugin{
				Id:          10000 + rand.Int63n(100000),
				Name:        conf.Name,
				Sign:        "未安装",
				Author:      conf.Author,
				Version:     conf.Version,
				Description: conf.Description,
				Contact:     conf.Contact,
				Path:        plugin,
				Page:        conf.Page,
				ErrMsg:      conf.Error,
				NoInstall:   true,
				Enable:      false,
				Status:      0,
			}
			if len(conf.Error) > 0 {
				item.Status = 3
			}
			*entities = append(*entities, item)
		}
	}
	return nil
}

func (c *PluginController) PostInstall() {
	path, _ := c.Input().GetString("path")
	if len(path) == 0 {
		helper.Ajax("请传入要安装的插件地址", 1, c.Ctx())
		return
	}
	if _, err := c.Orm.Transaction(func(session *xorm.Session) (interface{}, error) {
		mgr := plugins.PluginMgr()
		conf, err := mgr.GetPluginInfo(path)
		if err != nil {
			return nil, err
		}
		// 安装插件
		if entity, err := mgr.Install(path); err != nil {
			conf.Error = err.Error()
			return nil, err
		} else {
			mgr.Reload()
			var viewConf []map[string]interface{}
			err = json.Unmarshal([]byte(entity.View()), &viewConf)
			if err != nil {
				return nil, err
			}
			sign := entity.Sign()
			if len(sign) == 0 {
				return nil, errors.New("插件无签名, 无效")
			}
			pluginDb := &tables.Plugin{
				Name:        conf.Name,
				Author:      conf.Author,
				Contact:     conf.Contact,
				Description: conf.Description,
				Version:     conf.Version,
				Sign:        sign,
				View:        viewConf,
				Path:        path,
				CreatedAt:   tables.LocalTime(time.Now()),
			}
			_, err = c.Orm.Insert(pluginDb)
			if err != nil {
				return nil, err
			}
			entity.Menu(&tables.Menu{}, int(pluginDb.Id))
		}
		return nil, nil
	}); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
	} else {
		helper.Ajax("success", 0, c.Ctx())
	}
}

func (c *PluginController) PostEnable() {
	path, _ := c.Input().GetBytes("path")
	if path == nil || len(path) == 0 {
		helper.Ajax("请传入插件路径", 1, c.Ctx())
		return
	}
	enable,_ := c.Input().GetBool("enable")
	ret, _ := c.Orm.Where("path = ?", string(path)).Cols("enable").Update(&tables.Plugin{
		Enable:    enable,
		UpdatedAt: tables.LocalTime(time.Now()),
	})
	if ret > 0 {
		helper.Ajax("success", 0, c.Ctx())
	} else {
		helper.Ajax("failed", 1, c.Ctx())
	}
}

func (c *PluginController) GetConfig() {
	plugin := &tables.Plugin{}
	path, _ := c.Ctx().GetString("path")
	c.Orm.Where("path = ?", path).Get(plugin)
	helper.Ajax(plugin.Config, 0, c.Ctx())
}

func (c *PluginController) PostConfig() {
	plugin := &tables.Plugin{}
	if err := c.Ctx().BindJSON(plugin); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	plugin.UpdatedAt = tables.LocalTime(time.Now())
	res, _ := c.Orm.Where("path = ?", plugin.Path).Update(plugin)
	if res > 0 {
		helper.Ajax("success", 0, c.Ctx())
	} else {
		helper.Ajax("failed", 1, c.Ctx())
	}
}
