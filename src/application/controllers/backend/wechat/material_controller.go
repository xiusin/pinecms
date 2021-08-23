package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/common/helper"
)

type WechatMaterialController struct {
	backend.BaseController
}

func (c *WechatMaterialController) PostList(cacher cache.AbstractCache) {
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
	var list material.ArticleList
	cacheKey := fmt.Sprintf(CacheKeyWechatMaterialList, q.Appid, q.Type, q.Page)
	err = cacher.Remember(cacheKey, &list, func() (interface{}, error) {
		data, err := account.GetMaterial().BatchGetMaterial(material.PermanentMaterialType(q.Type), (q.Page-1)*q.Size, q.Size)
		if err == nil {
			SaveCacheMaterialListKey(cacheKey, cacher)
		}
		return data, err
	}, CacheTimeSecs)

	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(list, 0, c.Ctx())
}

func (c *WechatMaterialController) PostTotal(cacher cache.AbstractCache) {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	var ret material.ResMaterialCount
	cacheKey := fmt.Sprintf(CacheKeyWechatMaterialCount, appid)
	cacher.Remember(cacheKey, &ret, func() (interface{}, error) {
		data, err := account.GetMaterial().GetMaterialCount()
		if err == nil {
			SaveCacheMaterialListKey(cacheKey, cacher)
		}
		return data, err
	}, CacheTimeSecs)

	helper.Ajax(ret, 0, c.Ctx())
}

func (c *WechatMaterialController) PostClear(cacher cache.AbstractCache)  {
	var cacheKeys []string
	cacher.GetWithUnmarshal(CacheKeyWechatMaterialListKeys, &cacheKeys)
	for _, cacheKey := range cacheKeys {
		cacher.Delete(cacheKey)
	}
	cacher.Delete(CacheKeyWechatMaterialListKeys)
	helper.Ajax("清除微信素材", 0, c.Ctx())
}

func (c *WechatMaterialController) PostUpload() {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	filename := c.Input().Get("filename").String()
	mediaId, url, err := account.GetMaterial().AddMaterial(material.MediaTypeImage, filename)

	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	helper.Ajax(pine.H{"mediaId": mediaId, "url": url}, 0, c.Ctx())
}
