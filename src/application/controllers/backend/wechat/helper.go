package wechat

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	cache2 "github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

const CacheKeyWechatMaterialCount = "pinecms.wechat.material.count.%s"

const CacheKeyWechatMaterialList = "pinecms.wechat.material.list.%s.t.%s.p.%d"

const CacheKeyWechatMaterialListKeys = "pinecms.wechat.material.list.key"

const CacheTimeSecs = 30 * 24 * 3600

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

func SaveCacheMaterialListKey(key string, cacher cache2.AbstractCache)  {
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
