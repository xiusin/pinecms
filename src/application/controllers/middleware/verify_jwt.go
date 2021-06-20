package middleware

import (
	"fmt"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/config"
	"strings"
	"time"
)

func VerifyJwtToken() pine.Handler {
	return func(ctx *pine.Context) {
		if !strings.Contains(ctx.Path(), "login") && !strings.Contains(ctx.Path(), "/public/") {
			token := ctx.Header("Authorization")
			var hs = jwt.NewHS256([]byte(config.AppConfig().JwtKey))
			// 验证token
			var pl controllers.LoginAdminPayload
			_, err := jwt.Verify([]byte(token), hs, &pl)
			if err != nil {
				_ = ctx.Render().JSON(pine.H{"code": 1, "msg": "授权失败, 请重新登录"})
				return
			}
			ctx.Set("roleid", pl.RoleID)
			ctx.Set("adminid", pl.AdminId)
			if strings.Contains(ctx.Path(), "user/info") {
				ctx.QueryArgs().Set("id", fmt.Sprintf("%d", pl.AdminId))
			}
			if !strings.Contains(ctx.Path(), "/log/list") {
				// 记录操作日志
				ctx.Value("orm").(*xorm.Engine).Insert(&tables.Log{
					Uri:      string(ctx.RequestURI()),
					Userid:   pl.AdminId,
					Params:   string(ctx.PostBody()),
					Username: pl.AdminName,
					Ip:       ctx.ClientIP(),
					Time:     time.Now(),
					Method:   string(ctx.Method()),
				})
			}
		}
		ctx.Next()
	}
}
