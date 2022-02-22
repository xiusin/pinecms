package taglibs

import (
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/config"
)

/**
 * 标签：{{yield query(sql="") content}} {{end}}
 * 作用：特殊标签，SQL查询标签
 * 用法示例： {{yield query(sql="SELECT * FROM #@_tables") content}} .. HTML ..{{end}}
 * 参数说明：
 * 	sql SQL语句，只用于select类型语句
 */
func Query(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("Query Failed", err)
		}
	}()
	sess := helper.GetORM()
	query := strings.Trim(args.Get(0).String(), " \n\t")
	// 只允许查询操作
	conf := config.DB()
	if strings.HasPrefix(query, "SELECT") || strings.HasPrefix(query, "select") {
		rest, err := sess.QueryString(strings.ReplaceAll(query, "#@_", conf.Db.DbPrefix))
		if err != nil {
			panic(err)
		}
		if rest != nil {
			return reflect.ValueOf(rest)
		}
	}
	return reflect.ValueOf([]map[string]string{})
}
