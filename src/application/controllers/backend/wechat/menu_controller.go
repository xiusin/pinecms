package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/common/helper"
)

type editParam struct {
	Appid string `json:"appid"`
	Menu  struct {
		Button []*menu.Button `json:"button"`
		MenuID int64         `json:"menuid"`
	} `json:"menu"`
}

type WechatMenuController struct {
	backend.BaseController
	key string
	p editParam
}

func (c *WechatMenuController) Construct() {
	c.BaseController.Construct()
	c.p = editParam{}
}

func (c *WechatMenuController) PostEdit() {
	c.Ctx().BindJSON(&c.p)
	if len(c.p.Appid) == 0 || len(c.p.Menu.Button) == 0 {
		helper.Ajax("发布参数错误", 1, c.Ctx())
		return
	}
	account, _ := GetOfficialAccount(c.p.Appid)
	if err := account.GetMenu().SetMenu(c.p.Menu.Button); err != nil {
		helper.Ajax("设置菜单失败: " + err.Error(), 1, c.Ctx())
	} else {
		helper.Ajax("发布菜单成功", 0, c.Ctx())
	}
}

func (c *WechatMenuController) GetInfo(cacher cache.AbstractCache) {
	appid := c.Ctx().GetString("appid")
	if len(appid) == 0 {
		helper.Ajax("请选择公众号", 1, c.Ctx())
		return
	}
	c.key = fmt.Sprintf(CacheKeyWechatMenu, appid)
	account, _ := GetOfficialAccount(appid)
	retMenu, err := account.GetMenu().GetMenu()
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	for i, btn := range retMenu.Menu.Button {
		if len(btn.SubButtons) == 0 {
			btn.SubButtons = []*menu.Button{} // 初始化子菜单结构
		}
		retMenu.Menu.Button[i] = btn
	}
	helper.Ajax(retMenu.Menu, 0, c.Ctx())
}
