package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"strings"
)

var defaultArrReturnVal = reflect.ValueOf([]interface{}{})

func checkArgType(args *jet.Arguments) bool {
	l := args.NumOfArguments()
	for i := 0; i< l; i++ {
		t := args.Get(i)
		if !isNumber(t) && t.Type().String() != "string"  {
			pine.Logger().Errorf("参数类型不支持: idx: %d -> type: %s -> val: %s",i, t, args.Get(i))
			return false
		}
	}
	return true
}

func isNumber(val reflect.Value) bool {
	return strings.Contains(val.String(), "float")  || strings.Contains(val.String(), "int")
}

func getNumber(val reflect.Value) int64 {
	if strings.Contains(val.Type().String(), "float") {
		return int64(val.Float())
	} else if strings.Contains(val.Type().String(), "int") {
		return val.Int()
	}
	return 0
}

func getOrmSess(table ...string) *xorm.Session {
	if len(table) == 0 {
		table = []string{"articles"}
	}
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(controllers.GetTableName(table[0]))
}

func getCategoryOrm() *xorm.Session {
	return helper.GetORM().Table(&tables.Category{}).Where("ismenu = 1")
}