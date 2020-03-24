package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"reflect"
	"strings"
)

/**
 * 标签：{{yield query(sql="") content}} {{end}}
 * 作用：特殊标签，SQL查询标签
 * 用法示例： {{yield query(sql="SELECT * FROM tables") content}} .. HTML ..{{end}}
 * 参数说明：
 * 	sql SQL语句，只用于select类型语句
 */
func Query(args jet.Arguments) reflect.Value {
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("Query Failed", err)
		}
	}()
	sess := pine.Make("*xorm.Engine").(*xorm.Engine)
	query := strings.Trim(args.Get(0).String(), " ")
	// 只允许查询操作
	if strings.HasPrefix(query, "SELECT") || strings.HasPrefix(query, "select") {
		rest, err := sess.QueryString(query)
		if err != nil {
			panic(err)
		}
		if rest != nil {
			return reflect.ValueOf(rest)
		}
	}
	return reflect.ValueOf([]map[string]string{})
}
