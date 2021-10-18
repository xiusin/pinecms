package wechat

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

type WechatUserController struct {
	backend.BaseController
}

func (c *WechatUserController) Construct() {
	c.Table = &tables.WechatMember{}
	c.Entries = &[]tables.WechatMember{}

	c.SearchFields = []backend.SearchFieldDsl{
		{Field: "appid"},
		{Field: "province", Op: "LIKE", DataExp: "%$?%"},
		{Field: "tagid_list"},
		{Field: "city", Op: "LIKE", DataExp: "%$?%"},
		{Field: "nickname", Op: "LIKE", DataExp: "%$?%"},
		{Field: "remark", Op: "LIKE", DataExp: "%$?%"},
		{Field: "qr_scene_str", Op: "LIKE", DataExp: "%$?%"},
	}

	c.BaseController.Construct()
	c.OpAfter = c.After
	c.OpBefore = c.Before
}

func (c *WechatUserController) Before(act int, params interface{}) error {
	if act == backend.OpList {
		appid, _ := c.Input().GetString("appid")
		if len(appid) > 0 {
			(params.(*xorm.Session)).Where("appid = ?", appid)
		}
	}
	return nil
}

func (c *WechatUserController) After(act int, _ interface{}) error {
	if act == backend.OpEdit {
		account, _ := GetOfficialAccount(c.Table.(*tables.WechatMember).Appid)
		account.GetUser().UpdateRemark(c.Table.(*tables.WechatMember).Openid, c.Table.(*tables.WechatMember).Remark)
	}
	return nil
}

// PostSync 同步粉丝
func (c *WechatUserController) PostSync() {
	var q = &tables.WechatAccount{}
	if err := c.Ctx().BindJSON(q); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	if len(q.AppId) == 0 {
		helper.Ajax("请先选择一个公众号", 1, c.Ctx())
		return
	}
	account, data := GetOfficialAccount(q.AppId)
	if !data.Verified {
		helper.Ajax("公众号没有接入无法同步", 1, c.Ctx())
		return
	}
	nextOpenId, exit := "", false
	for !exit {
		users, err := account.GetUser().ListUserOpenIDs(nextOpenId)
		if err != nil {
			helper.Ajax(err, 1, c.Ctx())
			return
		}
		if users.Count < 1000 {
			exit = true
		}
		nextOpenId = users.NextOpenID

		for _, openid := range users.Data.OpenIDs {
			u, err := account.GetUser().GetUserInfo(openid)
			if err != nil {
				pine.Logger().Error("获取微信会员信息失败", err)
				continue
			}
			if exist, _ := c.Orm.Where("openid = ?", openid).Exist(&tables.WechatMember{}); exist {
				c.Orm.Where("openid = ?", openid).Cols("tagid_list").Update(&tables.WechatMember{
					TagidList: u.TagIDList,
					Remark:    u.Remark,
				})
			} else {
				c.Orm.InsertOne(&tables.WechatMember{
					Appid:          q.AppId,
					Openid:         u.OpenID,
					Nickname:       u.Nickname,
					Sex:            int(u.Sex),
					City:           u.City,
					Province:       u.Province,
					Headimgurl:     u.Headimgurl,
					SubscribeTime:  time.Unix(int64(u.SubscribeTime), 0).Format(helper.TimeFormat),
					Subscribe:      u.Subscribe > 0,
					Unionid:        u.UnionID,
					Remark:         u.Remark,
					TagidList:      u.TagIDList,
					SubscribeScene: u.SubscribeScene,
					QrSceneStr:     u.QrSceneStr,
				})
			}

		}

	}
	helper.Ajax("同步完成", 0, c.Ctx())
}

// PostSavePoster 保存海报 vue-canvas-poster
func (c *WechatUserController) PostSavePoster() {

}
