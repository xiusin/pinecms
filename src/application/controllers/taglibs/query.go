package taglibs

import (
	"reflect"
	"strings"

	"github.com/xiusin/pinecms/src/common/helper"

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
	var rest = []map[string]string{}
	helper.Cache().Remember("pinecms:tag:query:"+getTagHash(args), &rest, func() (any, error) {
		if !checkArgType(&args) {
			return &rest, nil
		}
		sess := helper.GetORM()
		query := strings.Trim(args.Get(0).String(), " \n\t")
		// 只允许查询操作
		conf := config.DB()
		if !strings.HasPrefix(strings.ToLower(query), "select") {
			return &rest, pine.NewError("只允许执行SELECT查询")
		}
		if strings.Contains(strings.ToLower(query), "into") {
			return &rest, pine.NewError("不允许执行INTO查询")
		}
		var err error
		rest, err = sess.QueryString(strings.ReplaceAll(query, "#@_", conf.Db.DbPrefix))
		if err != nil {
			pine.Logger().Error(err)
			return &rest, err
		}
		return &rest, nil
	})

	return reflect.ValueOf(rest)
}
