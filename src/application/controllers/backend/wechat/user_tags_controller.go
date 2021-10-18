package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type WechatUserTagsController struct {
	backend.BaseController
	p UserTags
}

func (c *WechatUserTagsController) Construct() {
	c.BaseController.Construct()
}

func (c *WechatUserTagsController) PostList(cacher cache.AbstractCache) {
	appid, _ := c.Input().GetString("appid")
	var tags = []*user.TagInfo{}
	if appid == "" {
		helper.Ajax("请输入公众号ID", 1, c.Ctx())
		return
	}
	cacheKey := fmt.Sprintf(CacheKeyWechatUserTags, appid)
	err := cacher.Remember(cacheKey, &tags, func() (interface{}, error) {
		account, _ := GetOfficialAccount(appid)
		return account.GetUser().GetTag()
	}, CacheTimeSecs)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(tags, 0, c.Ctx())
}

func (c *WechatUserTagsController) PostDelete(cacher cache.AbstractCache) {
	appid, _ := c.Input().GetString("appid")
	id, _ := c.Input().GetInt64("id")
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

func (c *WechatUserTagsController) PostTagging() {
	c.Ctx().BindJSON(&c.p)
	if c.p.Appid == "" || len(c.p.Openids) == 0 || c.p.Id == 0 || c.p.Action == "" {
		helper.Ajax("必要参数为空", 1, c.Ctx())
		return
	}
	account, _ := GetOfficialAccount(c.p.Appid)
	var err error
	if c.p.Action == "tagging" {
		err = account.GetUser().BatchTag(c.p.Openids, int32(c.p.Id))
	} else {
		err = account.GetUser().BatchUntag(c.p.Openids, int32(c.p.Id))
	}
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	c.refreshUserInfo(account)
	helper.Ajax("操作成功", 0, c.Ctx())
}

func (c *WechatUserTagsController) refreshUserInfo(account *officialaccount.OfficialAccount)  {
	for _, openid := range c.p.Openids {
		u, err := account.GetUser().GetUserInfo(openid)
		if err == nil {
			_, err = c.Orm.Where("openid = ?", openid).Cols("tagid_list").Update(&tables.WechatMember{TagidList: u.TagIDList})
			if err != nil {
				pine.Logger().Warning("更新用户", openid, "失败: ", err.Error())
			}
		} else {
			pine.Logger().Warning("更新用户", openid, "失败: ", err.Error())
		}
	}
}
