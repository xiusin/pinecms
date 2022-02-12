package middleware

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/Apiform"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/common"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/errcode"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"strings"
	"time"
)

func Auth() pine.Handler {
	return func(c *pine.Context) {
		if strings.Contains(c.Path(), "/ui") {
			c.Next()
			return
		}

		// 放行接口不验证
		if strings.Contains(c.Path(), "/login") {
			c.Next()
			return
		}

		var resp Apiform.Resp
		jwtToken := c.Header("Authorization")
		if jwtToken == "" || !strings.HasPrefix(jwtToken, "Bearer ") {
			resp.Code = errcode.S_auth_fmt_err
			resp.Msg = "Token不正确"
			c.Render().JSON(resp)
			return
		}
		jwtToken = jwtToken[7:]
		claims, err := common.ParseToken(jwtToken)
		if err != nil {
			resp.Code = errcode.S_auth_err
			resp.Msg = "Token错误，请重新登录"
			c.Render().JSON(resp)
			return
		}
		valid := claims.Valid()
		if valid != nil {
			resp.Code = errcode.S_auth_err
			resp.Msg = "用户登录超时，请重新登录"
			c.Render().JSON(resp)
			return
		}
		var userInfo tables.SSHUser

		helper.GetORM().Where("id = ?", claims.Userid).Get(&userInfo)
		if userInfo.Phone == 0 {
			resp.Code = errcode.S_auth_err
			resp.Msg = "用户不存在，请重新登录"
			c.Render().JSON(resp)
			return
		}
		c.Set("uid", claims.Userid)
		c.Set("token", "")
		newToken, err := common.ReleaseToken(claims.Userid)
		if err != nil {
			resp.Code = errcode.S_auth_err
			resp.Msg = err.Error()
			c.Render().JSON(resp)
			return
		}
		if time.Now().Add(24*time.Hour).Unix() > claims.ExpiresAt {
			c.Set("token", newToken)
		}
		c.Next()
	}
}
