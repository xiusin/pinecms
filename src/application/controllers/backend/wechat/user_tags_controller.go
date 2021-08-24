package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type WechatUserTagsController struct {
	backend.BaseController
}

func (c *WechatUserTagsController) Construct() {
	c.Table = &tables.WechatMember{}
	c.Entries = &[]tables.WechatMember{}
	c.BaseController.Construct()
}

func (c *WechatUserTagsController) PostList(cacher cache.AbstractCache) {
	appid := "wxe43df03110f5981b"
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, appid)
	var tags []*user.TagInfo
	if err := cacher.Remember(cacheKey, &tags, func() (interface{}, error) {
		account, _, err := GetOfficialAccount(appid)
		if err != nil {
			return nil, err
		}
		return account.GetUser().GetTag()
	}, CacheTimeSecs); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(tags, 0, c.Ctx())
}

func (c *WechatUserTagsController) PostDelete(cacher cache.AbstractCache) {
	appid := "wxe43df03110f5981b"
	id, _ := c.Input().Get("id").Int64()
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, appid)
	var tags []*user.TagInfo
	err := cacher.GetWithUnmarshal(cacheKey, &tags)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	var idx = -1
	for i, tag := range tags {
		if tag.ID == int32(id) {
			idx = i
			break
		}
	}
	if idx > -1 {
		account, _, err := GetOfficialAccount(appid)
		if err != nil {
			helper.Ajax(err, 1, c.Ctx())
			return
		}
		err = account.GetUser().DeleteTag(int32(id))
		if err != nil {
			helper.Ajax(err, 1, c.Ctx())
		} else {
			tags = append(tags[:idx], tags[idx+1:]...)
			cacher.SetWithMarshal(cacheKey, &tags, CacheTimeSecs)
			helper.Ajax("删除标签成功", 0, c.Ctx())
		}
	} else {
		helper.Ajax("标签不存在或已被删除", 1, c.Ctx())
	}
}

func (c *WechatUserTagsController) PostEdit(cacher cache.AbstractCache) {
	appid := "wxe43df03110f5981b"
	id, _ := c.Input().Get("id").Int64()
	name := c.Input().Get("name").String()
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, appid)
	var tags []*user.TagInfo
	if err := cacher.GetWithUnmarshal(cacheKey, &tags); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	for i, tag := range tags {
		if tag.ID == int32(id) {
			tag.Name = name
			tags[i] = tag
			account, _, err := GetOfficialAccount(appid)
			if err != nil {
				helper.Ajax(err, 1, c.Ctx())
				return
			}
			if err = account.GetUser().UpdateTag(int32(id), name); err != nil {
				helper.Ajax(err, 1, c.Ctx())
				return
			} else {
				cacher.SetWithMarshal(cacheKey, &tags, CacheTimeSecs)
				helper.Ajax("修改标签成功", 0, c.Ctx())
				return
			}
		}
	}
	helper.Ajax("没有找到标签", 1, c.Ctx())
}

// PostTagging 批量打标签
func (c *WechatUserTagsController) PostTagging() {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	account.GetUser().BatchTag([]string{}, 0)
}

// PostUntagging 批量解除标签
func (c *WechatUserTagsController) PostUntagging() {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	account.GetUser().BatchUntag([]string{}, 0)
}
