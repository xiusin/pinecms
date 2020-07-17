package middleware

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/config"
	"strings"
)

func VerifyJwtToken() pine.Handler {
	return func(ctx *pine.Context) {
		if !strings.Contains(ctx.Path(), "login") &&
			!strings.Contains(ctx.Path(), "/public/") /*public控制器下的不校验Token*/{
			token := ctx.Header("X-TOKEN")
			var hs = jwt.NewHS256([]byte(config.AppConfig().JwtKey))
			// 验证token
			var pl controllers.LoginAdminPayload
			_, err := jwt.Verify([]byte(token), hs, &pl)
			if err != nil {
				panic(err)
			}
			ctx.Set("roleid", pl.RoleID)
			ctx.Set("adminid", pl.AdminId)
		}
		ctx.Next()
	}
}

