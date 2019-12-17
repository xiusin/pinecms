package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
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

func (this *LoginController) Index() {
	if this.Ctx.Method() == "POST" {
		username := this.Ctx.PostValue("username")
		password := this.Ctx.PostValue("password")
		code := this.Ctx.PostValue("code")
		//verify, _ := this.Session.GetFlash("verify_code").(string)
		if helper.IsFalse(username, password, code) {
			helper.Ajax("参数不能为空", 1, this.Ctx)
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
		admin, err := models.NewAdminModel(this.Orm).Login(username, password, this.Ctx.RemoteAddr())
		if err != nil {
			helper.Ajax(err.Error(), 1, this.Ctx)
		} else {
			this.Session.Set("roleid", admin.Roleid)
			this.Session.Set("adminid", admin.Userid)
			this.Session.Set("username", admin.Username)
			helper.Ajax("登录成功", 0, this.Ctx)
		}
		return
	}

	if err := this.Ctx.View("backend/login_index.html"); err != nil {
		this.Ctx.WriteString(err.Error())
	}
}

//退出系统
func (this *LoginController) Logout() {
	this.Session.Delete("adminid")
	this.Session.Delete("roleid")
	this.Ctx.RemoveCookie("username")
	this.Ctx.RemoveCookie("userid")
	this.Ctx.Redirect("/b/login/index", iris.StatusFound)
}
