package wechat

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

type WechatMaterialController struct {
	backend.BaseController
}

func (c *WechatMaterialController) Construct() {
	c.Table = &tables.WechatMaterial{}
	c.Entries = &[]tables.WechatMaterial{}
	c.BaseController.Construct()
}

func (c *WechatMaterialController) PostList(cacher cache.AbstractCache) {
	var p = struct {
		Appid string `json:"appid"`
		Type  string `json:"type"`
		Page  int    `json:"page"`
		Size  int    `json:"size"`
	}{}
	if err := c.Ctx().BindJSON(&p); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	p.Size = 20

	count, err := c.Orm.Where("`type` = ?", p.Type).Where("appid = ?", p.Appid).Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(c.Entries)

	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	if count == 0 {
		c.PostSync()
	}

	helper.Ajax(pine.H{
		"list": c.Entries,
		"pagination": pine.H{
			"page":  p.Page,
			"size":  p.Size,
			"total": count,
		},
	}, 0, c.Ctx())
}

func (c *WechatMaterialController) PostSync() {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	_, err = c.Orm.Transaction(func(session *xorm.Session) (interface{}, error) {
		session.Where("id > 0").Delete(c.Table)
		types := []material.MediaType{material.MediaTypeImage, material.MediaTypeThumb, material.MediaTypeVoice, material.MediaTypeVideo}
		for _, mediaType := range types {
			page, size, finished := 1, 20, false
			for !finished {
				data, err := account.GetMaterial().BatchGetMaterial(material.PermanentMaterialType(mediaType), int64((page-1)*size), int64(size))
				if err != nil {
					return nil, err
				} else {
					for _, item := range data.Item {
						session.InsertOne(&tables.WechatMaterial{
							Appid:      appid,
							Type:       string(mediaType),
							MediaId:    item.MediaID,
							Url:        item.URL,
							UpdateTime: tables.LocalTime(time.Unix(item.UpdateTime, 0)),
						})
					}
				}
				page++
				finished = data.ItemCount == 0
			}
		}
		return nil, nil
	})
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
	} else {
		helper.Ajax("同步素材类型成功", 0, c.Ctx())
	}
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

func (c *WechatMaterialController) PostClear(cacher cache.AbstractCache) {
	var cacheKeys []string
	cacher.GetWithUnmarshal(CacheKeyWechatMaterialListKeys, &cacheKeys)
	for _, cacheKey := range cacheKeys {
		cacher.Delete(cacheKey)
	}
	cacher.Delete(CacheKeyWechatMaterialListKeys)
	helper.Ajax("清除微信素材", 0, c.Ctx())
}

// PostDelete 删除素材
func (c *WechatMaterialController) PostDelete() {
	err := c.Ctx().BindJSON(c.Table)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	appid := "wxe43df03110f5981b"

	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	if err = account.GetMaterial().DeleteMaterial(c.Table.(*tables.WechatMaterial).MediaId); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	c.Orm.Where("media_id = ?").Delete(c.Table)
	helper.Ajax("删除资源成功", 0, c.Ctx())
}

func (c *WechatMaterialController) PostUpload() {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	fileName := c.Ctx().FormValue("fileName")
	mediaType := c.Ctx().FormValue("mediaType")
	var mediaId, url string
	if material.MediaTypeVideo == material.MediaType(mediaType) {
		title := c.Ctx().FormValue("title")
		introduction := c.Ctx().FormValue("introduction")
		mediaId, url, err = account.GetMaterial().AddVideo(fileName, title, introduction)
	} else {
		mediaId, url, err = account.GetMaterial().AddMaterial(material.MediaType(mediaType), fileName)
	}
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(pine.H{"mediaId": mediaId, "url": url}, 0, c.Ctx())
}
