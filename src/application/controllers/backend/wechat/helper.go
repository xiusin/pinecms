package wechat

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

func GetOfficialAccount(appid string) (*officialaccount.OfficialAccount, *tables.WechatAccount, error) {
	accountData := &tables.WechatAccount{}
	orm := di.MustGet(controllers.ServiceXorm).(*xorm.Engine)
	orm.Where("app_id = ?", appid).Get(accountData)
	if accountData.Id == 0 {
		return nil, nil, errors.New("公众号" + appid + "不存在")
	}
	wc, memory := wechat.NewWechat(), cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          accountData.AppId,
		AppSecret:      accountData.Secret,
		Token:          accountData.Token,
		EncodingAESKey: accountData.AesKey,
		Cache:          memory,
	}
	account := wc.GetOfficialAccount(cfg)
	return account, accountData, nil
}
