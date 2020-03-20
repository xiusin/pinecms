package middleware

import (
	"strconv"
	"strings"

	"github.com/xiusin/pine"
)

func CheckDatabaseBackupDownload() func(ctx *pine.Context) {
	return func(ctx *pine.Context) {
		if strings.Contains(ctx.Request().URL.Path, "database/backup/"){
			aid, _ := strconv.Atoi(ctx.Session().Get("adminid"))
			roleId, _ := strconv.Atoi(ctx.Session().Get("roleid"))
			if aid ==0 || roleId == 0 {
				ctx.Abort(404)
				ctx.Stop()
			}
		}
		ctx.Next()
	}
}
