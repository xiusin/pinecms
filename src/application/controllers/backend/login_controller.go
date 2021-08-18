package backend

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"time"
)

type LoginController struct {
	pine.Controller
}

func (c *LoginController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/login", "Login")
}

func (c *LoginController) Login(orm *xorm.Engine) {
	var p loginUserParam
	apidoc.SetApiEntity(c.Ctx(), &apidoc.Entity{
		ApiParam: &p,
		AppId:    "admin",
		Group:    "登录模块",
		SubGroup: "系统登录",
		Title:    "登录系统",
		Desc:     "账号密码登录系统， 并且返回JWT凭证",
	})

	if err := parseParam(c.Ctx(), &p); err != nil || helper.IsFalse(p.Username, p.Password) {
		helper.Ajax("参数不能为空", 1, c.Ctx())
		return
	}

	sessVerifyCode := c.Session().Get(controllers.CacheVerifyCode)

	if sessVerifyCode != p.CaptchaId {
		helper.Ajax("验证码错误", 1, c.Ctx())
		return
	}

	// 读取登录人信息
	admin, err := models.NewAdminModel().Login(p.Username, p.Password, c.Ctx().ClientIP())
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	var hs = jwt.NewHS256([]byte(config.AppConfig().JwtKey))
	now := time.Now()
	pl := controllers.LoginAdminPayload{
		Payload: jwt.Payload{
			Subject:        "PineCMS",
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
		},
		Id:        admin.Userid,
		AdminId:   admin.Userid,
		RoleID:    admin.Roleid,
		AdminName: admin.Username,
	}
	if token, err := jwt.Sign(pl, hs); err != nil {
		helper.Ajax("登录失败", 1, c.Ctx())
	} else {
		helper.Ajax(pine.H{
			"role_name":  "超级管理员",
			"role_id":    admin.Roleid,
			"admin_id":   admin.Userid,
			"id":         admin.Userid,
			"admin_name": admin.Username,
			"token":      string(token),
		}, 0, c.Ctx())
	}

}
