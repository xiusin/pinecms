package wechat

import (
	"bytes"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	helper.Ajax(pine.H{
		"list": c.Entries,
		"pagination": pine.H{
			"page":  p.Page,
			"size":  p.Size,
			"total": count,
		},
	}, 0, c.Ctx())
}

func (c *WechatMaterialController) GetPreview() {
	c.Ctx().Response.Header.Set("Cache-Control", "max-age=78400")
	url, _ := c.Ctx().GetString("url")
	resp, err := http.Get(url)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	defer resp.Body.Close()

	c.Ctx().SetContentType("img/png")
	d, _ := io.ReadAll(resp.Body)
	c.Ctx().Write(d)
}

func (c *WechatMaterialController) PostSync() {
	appid := "wxe43df03110f5981b"
	account, _ := GetOfficialAccount(appid)

	_, err := c.Orm.Transaction(func(session *xorm.Session) (interface{}, error) {
		session.Where("id > 0").Delete(c.Table)
		types := []material.PermanentMaterialType{
			material.PermanentMaterialTypeImage,
			material.PermanentMaterialTypeVideo,
			material.PermanentMaterialTypeVoice,
			material.PermanentMaterialTypeNews,
		}
		for _, mediaType := range types {
			page, size, finished := 1, 20, false
			for !finished {
				data, err := account.GetMaterial().BatchGetMaterial(mediaType, int64((page-1)*size), int64(size))
				if err != nil {
					return nil, err
				} else {
					for _, item := range data.Item {
						fmt.Println(session.InsertOne(&tables.WechatMaterial{
							Appid:      appid,
							Type:       string(mediaType),
							MediaId:    item.MediaID,
							Url:        item.URL,
							UpdateTime: tables.LocalTime(time.Unix(item.UpdateTime, 0)),
						}))
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
	account, _ := GetOfficialAccount(appid)

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
	var data tables.WechatMaterial
	c.Ctx().BindJSON(&data)
	if data.MediaId == "" {
		helper.Ajax("请选择要删除的素材", 1, c.Ctx())
		return
	}
	c.Orm.Where("media_id = ?", data.MediaId).Get(&data)
	if data.Id == 0 {
		helper.Ajax("素材不存在或已删除", 1, c.Ctx())
		return
	}
	account, _ := GetOfficialAccount(data.Appid)
	if err := account.GetMaterial().DeleteMaterial(data.MediaId); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	c.Orm.Where("media_id = ?", data.MediaId).Delete(c.Table)
	helper.Ajax("删除资源成功", 0, c.Ctx())
}

func (c *WechatMaterialController) PostUpload() {
	var form MaterialUploadForm
	c.Ctx().BindForm(&form)
	if len(form.Appid) == 0 || form.MediaType == "" || form.FileName == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	mf, err := c.Ctx().MultipartForm()
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	fss, ok := mf.File["file"]
	if !ok {
		helper.Ajax("打开上传临时文件失败", 1, c.Ctx())
		return
	}
	fs := fss[0]
	uploadFile, _ := fs.Open()
	defer uploadFile.Close()
	data, err := io.ReadAll(uploadFile)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	tmpnam := filepath.Join(os.TempDir(), fs.Filename)
	f, err := os.Create(tmpnam)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	defer func() {
		f.Close()
		os.Remove(tmpnam)
	}()

	if _, err = io.Copy(f, bytes.NewBuffer(data)); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	account, _ := GetOfficialAccount(form.Appid)
	var mediaId, url string
	if material.MediaTypeVideo == material.MediaType(form.MediaType) {
		mediaId, url, err = account.GetMaterial().AddVideo(f.Name(), form.Title, form.Introduction)
	} else {
		mediaId, url, err = account.GetMaterial().AddMaterial(material.MediaType(form.MediaType), f.Name())
	}
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	// 入库数据
	c.Orm.Insert(&tables.WechatMaterial{
		Appid:      form.Appid,
		Type:       form.MediaType,
		MediaId:    mediaId,
		Url:        url,
		UpdateTime: tables.LocalTime(time.Now()),
	})

	helper.Ajax(pine.H{"mediaId": mediaId, "url": url}, 0, c.Ctx())
}
