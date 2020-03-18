package taglibs

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"reflect"
	"strings"
)

func getInt(val reflect.Value) int {
	str := val.Type().String()
	if strings.HasPrefix(str, "float") {
		return int(val.Float())
	} else if strings.Contains(str, "int") {
		return int(val.Int())
	}
	pine.Logger().Error("helper.go: 参数必须为整型: " + str + ":" + val.String())
	panic("getInt failed")
}

func getOrmSess(table ...string) *xorm.Session {
	if len(table) == 0 {
		table = []string{"articles"}
	}
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(controllers.GetTableName(table[0]))
}
