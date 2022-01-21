package wechat

import (
	"errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

func GetOfficialAccount(appid string) (*officialaccount.OfficialAccount, *tables.WechatAccount) {
	accountData := &tables.WechatAccount{}
	orm := helper.GetORM()
	orm.Where("app_id = ?", appid).Get(accountData)
	if accountData.Id == 0 {
		panic(errors.New("公众号" + appid + "不存在"))
	}
	wc, memory := wechat.NewWechat(), &WechatTokenCacher{AbstractCache: di.MustGet(controllers.ServiceICache).(cache.AbstractCache)}
	cfg := &offConfig.Config{
		AppID:          accountData.AppId,
		AppSecret:      accountData.Secret,
		Token:          accountData.Token,
		EncodingAESKey: accountData.AesKey,
		Cache:          memory,
	}
	account := wc.GetOfficialAccount(cfg)
	return account, accountData
}

func SaveCacheMaterialListKey(key string, cacher cache.AbstractCache) {
	var cacheKeys []string
	cacher.GetWithUnmarshal(CacheKeyWechatMaterialListKeys, &cacheKeys)
	for _, cacheKey := range cacheKeys {
		if cacheKey == key {
			return
		}
	}
	cacheKeys = append(cacheKeys, key)
	cacher.SetWithMarshal(CacheKeyWechatMaterialListKeys, &cacheKeys, CacheTimeSecs)
}
