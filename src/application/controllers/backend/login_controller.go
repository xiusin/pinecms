package backend

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

type LoginController struct {
	pine.Controller
}

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

	parseParam(c.Ctx(), &p)

	if p.Password == "" || p.Username == "" || p.CaptchaId == "" || p.CaptchaValue == "" {
		helper.Ajax("参数不能为空", 1, c.Ctx())
		return
	}

	// if !captcha.Verify(p.CaptchaId, p.CaptchaValue) {
	// 	helper.Ajax("验证码错误, 请重新输入", 1, c.Ctx())
	// 	return
	// }

	// 读取登录人信息
	admin, err := models.NewAdminModel().Login(p.Username, p.Password, c.Ctx().ClientIP())
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	var hs = jwt.NewHS256([]byte(config.App().JwtKey))
	now := time.Now()
	pl := controllers.LoginAdminPayload{
		Payload: jwt.Payload{
			Subject:        "PineCMS",
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
		},
		Id:        admin.Userid,
		AdminId:   admin.Userid,
		RoleID:    admin.RoleIdList,
		AdminName: admin.Username,
	}
	if token, err := jwt.Sign(pl, hs); err != nil {
		helper.Ajax("登录失败", 1, c.Ctx())
	} else {
		helper.Ajax(pine.H{
			"role_id":    admin.RoleIdList,
			"admin_id":   admin.Userid,
			"id":         admin.Userid,
			"admin_name": admin.Username,
			"token":      string(token),
		}, 0, c.Ctx())
	}
}
