package wechat

import (
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type WechatAccountController struct {
	backend.BaseController
}

func (c *WechatAccountController) Construct() {
	c.Table = &tables.WechatAccount{}
	c.Entries = &[]tables.WechatAccount{}
	c.BaseController.Construct()
}

func (c *WechatAccountController) PostClear() {
	appid, _ := c.Input().GetString("appid")
	account, _ := GetOfficialAccount(appid)

	if err := account.GetBasic().ClearQuota(); err != nil {
		helper.Ajax(err, 1, c.Ctx())
	} else {
		helper.Ajax("清除请求限制成功", 0, c.Ctx())
	}
}

func (c *WechatAccountController) PostSelect() {
	_ = c.Orm.Find(c.Entries)
	m := c.Entries.(*[]tables.WechatAccount)
	var kv []tables.KV
	for _, model := range *m {
		kv = append(kv, tables.KV{
			Label: model.Name,
			Value: model.AppId,
		})
	}
	helper.Ajax(kv, 0, c.Ctx())
}

// PostDistribution 会员分布
func (c *WechatAccountController) PostDistribution() {
	appid, _ := c.Input().GetString("appid")
	if len(appid) == 0 {
		helper.Ajax("请选择公众号", 1, c.Ctx())
		return
	}

	typ, _ := c.Input().GetInt("type")
	field := "province"
	if typ == 2 {
		field = "city"
	}

	data, _ := c.Orm.QueryString("SELECT " + field + ", COUNT(*) AS total FROM " +
		controllers.GetTableName("wechat_member") + " WHERE appid='" + appid + "' GROUP BY " + field)

	helper.Ajax(data, 0, c.Ctx())
}
