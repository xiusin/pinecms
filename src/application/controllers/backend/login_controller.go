package backend

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"time"
)

type LoginController struct {
	pine.Controller
}

/**
todo 系统无法兼容: /dist分组 和 /dist/category分组
dist-category 路由段无法解析 /v2/dict-category/list
*/

func (c *LoginController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/login", "Login")
}

func (c *LoginController) Login() {
	var p loginUserParam
	apidoc.SetApiEntity(c.Ctx(), &apidoc.Entity{
		ApiParam: &p,
		AppId:    "admin",
		Group:    "登录模块",
		SubGroup: "系统登录",
		Title:    "登录系统",
		Desc:     "账号密码登录系统， 并且返回JWT凭证",
	})

	if err := parseParam(c.Ctx(), &p); err != nil {
		helper.Ajax("参数不能为空", 1, c.Ctx())
		return
	}

	if helper.IsFalse(p.Username, p.Password) {
		helper.Ajax("参数不能为空", 1, c.Ctx())
		return
	}
	var hs = jwt.NewHS256([]byte(config.AppConfig().JwtKey))
	now := time.Now()
	pl := controllers.LoginAdminPayload{
		Payload: jwt.Payload{
			Subject:        "PineCMS",
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
		},
		Id:        1,
		AdminId:   1,       //admin.Userid,
		RoleID:    1,       //admin.Roleid,
		AdminName: "admin", //admin.Username,
	}
	token, err := jwt.Sign(pl, hs)
	if err != nil {
		helper.Ajax("登录失败： 授权失败", 1, c.Ctx())
		return
	}
	helper.Ajax(pine.H{
		"role_name":  "超级管理员",
		"role_id":    1,
		"admin_id":   1,
		"id":         1,
		"admin_name": "admin",
		"token":      string(token),
	}, 0, c.Ctx())
}
