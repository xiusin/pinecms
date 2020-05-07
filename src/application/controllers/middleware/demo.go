package middleware

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"strings"
)

func Demo() pine.Handler {
	return func(ctx *pine.Context) {
		if !strings.Contains(ctx.Path(), "login") {
			if ctx.PostArgs().Len() > 0 {
				helper.Ajax("演示数据,无法修改", 1, ctx)
				ctx.Stop()
				return
			}
		}
		ctx.Next()
	}
}
