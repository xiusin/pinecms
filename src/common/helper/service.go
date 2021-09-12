package helper

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
)

func XormEngine() *xorm.Engine {
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine)
}

func AbstractCache() cache.AbstractCache {
	return pine.Make(controllers.ServiceICache).(cache.AbstractCache)
}
