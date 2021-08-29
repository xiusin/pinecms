package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

const templateDeleteURL = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"

type WechatMsgTemplateController struct {
	backend.BaseController
}

func (c *WechatMsgTemplateController) Construct() {
	c.Table = &tables.WechatMsgTemplate{}
	c.Entries = &[]tables.WechatMsgTemplate{}
	c.BaseController.Construct()
}

//func (c *WechatMsgTemplateController) PostDelete() {
//	c.Ctx().BindJSON(c.Table)
//	t := c.Table.(*tables.WechatMsgTemplate)
//	if t.TemplateId == "" || t.Appid == "" {
//		helper.Ajax("参数错误", 1, c.Ctx())
//		return
//	}
//	t.Appid = "wxe43df03110f5981b"
//	account, _ := GetOfficialAccount(t.Appid)
//
//	var accessToken string
//
//	accessToken, err := account.GetTemplate().GetAccessToken()
//	if err != nil {
//		helper.Ajax(err, 1, c.Ctx())
//		return
//	}
//	uri := fmt.Sprintf("%s?access_token=%s", templateDeleteURL, accessToken)
//	var response []byte
//	response, err = util.HTTPGet(uri)
//	if err != nil {
//		helper.Ajax(err, 1, c.Ctx())
//		return
//	}
//	var res = struct{ util.CommonError }{}
//	err = util.DecodeWithError(response, &res, "DeleteTemplate")
//	if err != nil {
//		helper.Ajax(err, 1, c.Ctx())
//		return
//	}
//	helper.Ajax("删除成功", 0, c.Ctx())
//}

func (c *WechatMsgTemplateController) PostSync() {
	appid := "wxe43df03110f5981b"
	account, _ := GetOfficialAccount(appid)

	templateList, err := account.GetTemplate().List()
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	for _, item := range templateList {
		if exist, _ := c.Orm.Where("appid = ?", appid).Where("template_id = ?", item.TemplateID).Exist(); !exist {
			temp := tables.WechatMsgTemplate{
				Appid:           appid,
				TemplateId:      item.TemplateID,
				Title:           item.Title, // 标题
				PrimaryIndustry: item.PrimaryIndustry,
				DeputyIndustry:  item.DeputyIndustry,
				Content:         item.Content,
				Example:         item.Example,
				UpdatedAt:       tables.LocalTime(time.Now()),
			}
			c.Orm.InsertOne(&temp)
		}
	}
	helper.Ajax("同步成功", 0, c.Ctx())
}

func (c *WechatMsgTemplateController) PostSend() {
	p := struct {
		AppId      string `json:"appid"`
		TemplateId string `json:"template_id"`
		Data       []struct {
			Color string `json:"color"`
			Name  string `json:"name"`
			Value string `json:"value"`
		}
		Miniprogram struct {
			Appid    string `json:"appid"`
			PagePath string `json:"pagePath"`
		}
		Url                string `json:"url"`
		WxUserFilterParams map[string]interface{}
	}{}

	if err := c.Ctx().BindJSON(&p); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	if p.TemplateId == "" || p.AppId == "" || len(p.Data) == 0 || len(p.WxUserFilterParams) == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	account, _ := GetOfficialAccount(p.AppId)
	sess := c.Orm.Where("appid = ?", p.AppId)
	if nickname := p.WxUserFilterParams["nickname"].(string); len(nickname) > 0 {
		sess.Where("nickname like ?", "%"+nickname+"%")
	}
	if city := p.WxUserFilterParams["city"].(string); len(city) > 0 {
		sess.Where("city like ?", "%"+city+"%")
	}
	if province := p.WxUserFilterParams["province"].(string); len(province) > 0 {
		sess.Where("province like ?", "%"+province+"%")
	}
	if remark := p.WxUserFilterParams["remark"].(string); len(remark) > 0 {
		sess.Where("remark like ?", "%"+remark+"%")
	}
	if qrScene := p.WxUserFilterParams["qrScene"].(string); len(qrScene) > 0 {
		sess.Where("qr_scene_str like ?", "%"+qrScene+"%")
	}

	var users []tables.WechatMember
	sess.Cols("openid").Get(&users)

	pine.Logger().Print(sess.LastSQL())

	if len(users) == 0 {
		helper.Ajax("无法查找到相关推送用户", 1, c.Ctx())
		return
	}

	failed, msgInfo := 0, map[string]*message.TemplateDataItem{}
	for _, datum := range p.Data {
		msgInfo[datum.Name] = &message.TemplateDataItem{Value: datum.Value, Color: datum.Color}
	}

	for _, user := range users {
		_, err := account.GetTemplate().Send(&message.TemplateMessage{
			ToUser:     user.Openid,
			TemplateID: p.TemplateId,
			Data:       msgInfo,
		})
		if err != nil {
			failed++
			pine.Logger().Error(fmt.Sprintf("发送消息给%s失败", user.Openid), err)
		}
	}
	helper.Ajax(fmt.Sprintf("发送信息%d成功, %d失败", failed, len(users)-failed), 0, c.Ctx())
}
