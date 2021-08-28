package wechat

import (
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"time"
)

type WechatQrcodeController struct {
	backend.BaseController
}

func (c *WechatQrcodeController) Construct() {
	c.Table = &tables.WechatQrcode{}
	c.Entries = &[]tables.WechatQrcode{}
	c.BaseController.Construct()
	c.OpAfter = c.after
}

func (c *WechatQrcodeController) after(act int, params interface{}) error {
	if act == backend.OpAdd {
		appid := "wxe43df03110f5981b"
		account, _ := GetOfficialAccount(appid)

		data := c.Table.(*tables.WechatQrcode)
		var req *basic.Request
		if data.IsTemp {
			var t time.Duration
			if time.Time(data.ExpireTime).Before(time.Now()) {
				t = 0
			} else {
				t = time.Time(data.ExpireTime).Sub(time.Now())
			}
			req = basic.NewTmpQrRequest(t, data.SceneStr)
		} else {
			req = basic.NewLimitQrRequest(data.SceneStr)
		}
		ticket, err := account.GetBasic().GetQRTicket(req)
		if err != nil {
			pine.Logger().Error("获取传参二维码失败", err)
			return err
		}
		data.Ticket = ticket.Ticket

		data.Url = ticket.URL

		c.Orm.Where("id = ?", data.Id).Update(data)

	}
	return nil
}
