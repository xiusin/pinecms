package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/common/helper"
	"strings"
)

type WechatUserTagsController struct {
	backend.BaseController
	p UserTags
}

type UserTags struct {
	Appid string `json:"appid"`
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

func (c *WechatUserTagsController) Construct() {
	c.BaseController.Construct()
}

func (c *WechatUserTagsController) PostList(cacher cache.AbstractCache) {
	appid := strings.Trim(c.Input().Get("appid").String(), `"`)
	var tags []*user.TagInfo
	if appid == "" {
		helper.Ajax("请输入公众号ID", 1, c.Ctx())
		return
	}
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, appid)
	cacher.Remember(cacheKey, &tags, func() (interface{}, error) {
		account, _ := GetOfficialAccount(appid)
		return account.GetUser().GetTag()
	}, CacheTimeSecs)

	helper.Ajax(tags, 0, c.Ctx())
}

func (c *WechatUserTagsController) PostDelete(cacher cache.AbstractCache) {
	appid := strings.Trim(c.Input().Get("appid").String(), `"`)
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
		account, _ := GetOfficialAccount(appid)
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

func (c *WechatUserTagsController) PostAdd(cacher cache.AbstractCache) {
	c.Ctx().BindJSON(&c.p)
	if c.p.Appid == "" || c.p.Name == "" {
		helper.Ajax("必要参数为空", 1, c.Ctx())
		return
	}
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, c.p.Appid)
	account, _ := GetOfficialAccount(c.p.Appid)
	if tagInfo, err := account.GetUser().CreateTag(c.p.Name); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	} else {
		var tags []*user.TagInfo
		cacher.GetWithUnmarshal(cacheKey, &tags)
		tags = append(tags, tagInfo)
		cacher.SetWithMarshal(cacheKey, &tags, CacheTimeSecs)
		helper.Ajax("新增标签成功", 0, c.Ctx())
		return
	}
}

func (c *WechatUserTagsController) PostEdit(cacher cache.AbstractCache) {
	c.Ctx().BindJSON(&c.p)
	if c.p.Appid == "" || c.p.Name == "" || c.p.Id == 0 {
		helper.Ajax("必要参数为空", 1, c.Ctx())
		return
	}
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, c.p.Appid)
	var tags []*user.TagInfo
	if err := cacher.GetWithUnmarshal(cacheKey, &tags); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	for i, tag := range tags {
		if tag.ID == int32(c.p.Id) {
			tag.Name = c.p.Name
			tags[i] = tag
			account, _ := GetOfficialAccount(c.p.Appid)
			if err := account.GetUser().UpdateTag(int32(c.p.Id), c.p.Name); err != nil {
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
	account, _ := GetOfficialAccount(appid)

	account.GetUser().BatchTag([]string{}, 0)
}

// PostUntagging 批量解除标签
func (c *WechatUserTagsController) PostUntagging() {
	appid := "wxe43df03110f5981b"
	account, _ := GetOfficialAccount(appid)

	account.GetUser().BatchUntag([]string{}, 0)
}
