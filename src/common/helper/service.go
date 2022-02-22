package helper

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
)

func AbstractCache() cache.AbstractCache {
	return pine.Make(controllers.ServiceICache).(cache.AbstractCache)
}

func App() *pine.Application {
	return pine.Make(controllers.ServiceApplication).(*pine.Application)
}
