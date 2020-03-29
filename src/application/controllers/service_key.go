package controllers

import (
	"fmt"
	"github.com/xiusin/logger"
	"reflect"
)

const ServiceConfig = "pinecms.config"

const ServiceICache = "cache.AbstractCache"

const ServiceTablePrefix = "pinecms.table_prefix"

const ServiceJetEngine = "pinecms.jet"

const ServiceXorm = "*xorm.Engine"

func init()  {
	fmt.Println("直接通过Nil反射类型", reflect.TypeOf((*logger.AbstractLogger)(nil)).Elem())
}
