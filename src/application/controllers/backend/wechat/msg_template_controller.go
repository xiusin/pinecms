package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/util"
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

func (c *WechatMsgTemplateController) PostDelete() {
	c.Ctx().BindJSON(c.Table)
	t := c.Table.(*tables.WechatMsgTemplate)
	if t.TemplateId == "" || t.Appid == "" {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	t.Appid = "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(t.Appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	var accessToken string

	accessToken, err = account.GetTemplate().GetAccessToken()
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", templateDeleteURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	var res = struct{ util.CommonError }{}
	err = util.DecodeWithError(response, &res, "DeleteTemplate")
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax("删除成功", 0, c.Ctx())
}

func (c *WechatMsgTemplateController) PostSync() {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

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

func (c *WechatMsgTemplateController) PostSend()  {
	appid := "wxe43df03110f5981b"
	account, _, err := GetOfficialAccount(appid)
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	fmt.Println(account.GetTemplate().Send(&message.TemplateMessage{
		ToUser:     "op-oIuJztLwLOoNBX6hNoOzHFEws",
		TemplateID: "97_xBnPYcKDeGjUaq9BxDvUQneDAM4-mAEd8Jt0VRu4",
		Data: map[string]*message.TemplateDataItem{
			"first": {"欢迎注册成为新会员", ""},
			"keyword1": {"陈xiusin", ""},
			"keyword2": {"16601313660", ""},
			"keyword3": {"10", ""},
			"keyword4": {"VIP.1", ""},
			"keyword5": {"1110000011000001", ""},
			"remark": {"会员享有系统内所有特权!", ""},
		},
	}))
}
