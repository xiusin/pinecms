package taglibs

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
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
	panic("helper.go: 参数必须为整型: " + str + ":" + val.String())
}

func getOrmSess() *xorm.Session {
	return pine.Make("*xorm.Engine").(*xorm.Engine).Table("iriscms_articles")
}