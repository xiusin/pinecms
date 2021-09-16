package middleware

import (
	"strings"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)


// Demo 演示项目示例限制操作
func Demo() pine.Handler {
	return func(ctx *pine.Context) {
		uri := ctx.Path()
		if strings.Contains(uri, "del") || strings.Contains(uri, "edit") || strings.Contains(uri, "update") || strings.Contains(uri, "save") {
			helper.Ajax("演示项目不允许修改信息数据", 1, ctx)
			return
		}
		ctx.Next()
	}
}
