package backend

import (
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type LoginController struct {
	pine.Controller
}

func (c *LoginController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY( "/login/index", "Index")
	b.ANY( "/login/logout", "Logout")
}

func (c *LoginController) Index(orm *xorm.Engine) {
	if c.Ctx().IsPost() {
		username := c.Ctx().PostString("username")
		password := c.Ctx().PostString("password")
		code := c.Ctx().PostString("code")
		//verify := c.Session().Get("verify")
		if helper.IsFalse(username, password, code) {
			helper.Ajax("参数不能为空", 1, c.Ctx())
			return
		}
		//if verify == "" {
		//	helper.Ajax("验证码过期,无法验证", 1, c.Ctx())
		//	return
		//}
		//if strings.ToLower(code) != strings.ToLower(verify) {
		//	helper.Ajax("验证码错误", 1, c.Ctx())
		//	return
		//}
		admin, err := models.NewAdminModel().Login(username, password, c.Ctx().ClientIP())
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
		} else {
			c.Session().Set("roleid", strconv.Itoa(int(admin.Roleid)))
			role := &tables.AdminRole{Roleid:admin.Roleid}
			orm.Get(role)
			c.Session().Set("role_name", role.Rolename)
			c.Session().Set("adminid", strconv.Itoa(int(admin.Userid)))
			c.Session().Set("username", admin.Username)

			// 处理 V2
			//if c.Ctx().Value("IsV2").(bool) {
			//	helper.AjaxV2(pine.H{
			//		"key": "X-TOKEN",
			//		"token": 1,
			//	},0, c.Ctx())
			//	return
			//}

			helper.Ajax("登录成功", 0, c.Ctx())
		}
		return
	}

	c.Ctx().Render().HTML("backend/login_index.html")
}

func (c *LoginController) Logout() {
	c.Session().Remove("adminid")
	c.Session().Remove("roleid")
	c.Ctx().RemoveCookie("username")
	c.Ctx().RemoveCookie("userid")
	c.Session().Remove("role_name")
	c.Session().Remove("username")
	c.Ctx().Redirect("/b/login/index")
}
