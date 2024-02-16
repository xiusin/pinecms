package helper

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
)

// 获取缓存服务
func AbstractCache() cache.AbstractCache {
	return pine.Make(controllers.ServiceICache).(cache.AbstractCache)
}

// 获取应用实例
func App() *pine.Application {
	return pine.Make(controllers.ServiceApplication).(*pine.Application)
}
