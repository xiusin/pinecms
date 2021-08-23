package wechat

import (
	"fmt"
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
	c.BaseController.Construct()
}

// PostSync 同步粉丝
func (c *WechatUserController) PostSync() {
	var q = &tables.WechatAccount{}
	if err := c.Ctx().BindJSON(q); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	fmt.Println()

	account, _, err := GetOfficialAccount(q.AppId)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	//if !data.Verified {
	//	helper.Ajax("公众号没有接入无法同步", 1, c.Ctx())
	//	return
	//}
	nextOpenId, exit := "", false
	//var ch = make(chan struct{}, 10) todo 并发携程控制
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
			exist, _ := c.Orm.Where("openid = ?", openid).Exist(&tables.WechatMember{})
			if exist {
				continue
			}
			u, err := account.GetUser().GetUserInfo(openid)
			if err != nil {
				pine.Logger().Error("获取微信会员信息失败", err)
				continue
			}
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
				TagidList:      nil,
				SubscribeScene: u.SubscribeScene,
				QrSceneStr:     u.QrSceneStr,
			})
		}

	}
	helper.Ajax("同步完成", 0, c.Ctx())
}

// PostSavePoster 保存海报 vue-canvas-poster
func (c *WechatUserController) PostSavePoster()  {

}
