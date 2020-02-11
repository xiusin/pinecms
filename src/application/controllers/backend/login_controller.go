package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/common/helper"
)

type LoginController struct {
	Orm     *xorm.Engine
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/login/index", "Index")
	b.Handle("ANY", "/login/logout", "Logout")
}

func (c *LoginController) Index() {
	if c.Ctx.Method() == "POST" {
		username := c.Ctx.PostValue("username")
		password := c.Ctx.PostValue("password")
		code := c.Ctx.PostValue("code")
		//verify, _ := this.Session.GetFlash("verify_code").(string)
		if helper.IsFalse(username, password, code) {
			helper.Ajax("参数不能为空", 1, c.Ctx)
			return
		}
		//if verify == "" {
		//	helper.Ajax("验证码过期,无法验证", 1, this.Ctx)
		//	return
		//}
		//if strings.ToLower(code) != strings.ToLower(verify) {
		//	helper.Ajax("验证码错误", 1, this.Ctx)
		//	return
		//}
		admin, err := models.NewAdminModel(c.Orm).Login(username, password, c.Ctx.RemoteAddr())
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx)
		} else {
			c.Session.Set("roleid", admin.Roleid)
			c.Session.Set("adminid", admin.Userid)
			c.Session.Set("username", admin.Username)
			helper.Ajax("登录成功", 0, c.Ctx)
		}
		return
	}

	c.Ctx.View("backend/login_index.html")
}

//退出系统
func (c *LoginController) Logout() {
	c.Session.Delete("adminid")
	c.Session.Delete("roleid")
	c.Ctx.RemoveCookie("username")
	c.Ctx.RemoveCookie("userid")
	c.Ctx.Redirect("/b/login/index", iris.StatusFound)
}
