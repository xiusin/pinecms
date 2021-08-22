package wechat

import (
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/common/helper"
)

type WechatMaterialController struct {
	backend.BaseController
}

func (c *WechatMaterialController) PostList() {
	var q = struct {
		Appid string `json:"appid"`
		Type  string `json:"type"`
		Page  int64  `json:"page"`
		Size  int64  `json:"size"`
	}{}
	if err := c.Ctx().BindJSON(&q); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	account, _, err := GetOfficialAccount(q.Appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	list, err := account.GetMaterial().BatchGetMaterial(material.PermanentMaterialType(q.Type), (q.Page - 1) * q.Size, q.Size)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(list, 0, c.Ctx())
}

//
//func (c *WechatMaterialController) materialTotal() (material.ResMaterialCount, error) {
//	account, _, err := GetOfficialAccount(c.appId)
//	if err != nil {
//		return material.ResMaterialCount{}, err
//	}
//	return account.GetMaterial().GetMaterialCount()
//}
