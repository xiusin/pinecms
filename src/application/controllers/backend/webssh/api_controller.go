package webssh

import (
	"fmt"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/Apiform"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/common"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/errcode"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

type ApiController struct {
	backend.BaseController
}

func (c *ApiController) PostLogin() {
	var resp Apiform.Resp
	resp.Code = errcode.C_from_err
	resp.Msg = "手机号和验证码不能为空！"
	var user Apiform.Login
	if err := c.Ctx().BindForm(&user); err == nil {
		if user.Phone > 0 {
			var userinfo tables.SSHUser
			exists, err := c.Orm.Where("phone = ?", user.Phone).Get(&userinfo)
			if !exists {
				userinfo = tables.SSHUser{Phone: user.Phone, Password: helper.GetMd5("123456")}
				id, _ := c.Orm.InsertOne(&userinfo)
				userinfo.Id = id
			}
			newToken, err := common.ReleaseToken(userinfo.Id)
			fmt.Println(newToken, err, userinfo)
			if err == nil && userinfo.Id > 0 {
				resp.Code = errcode.C_nil_err
				resp.Msg = "登陆成功"
				resp.Data = userinfo
				resp.Token = newToken
			} else {
				resp.Code = errcode.S_auth_err
				if err == nil {
					resp.Msg = "Token创建失败"
				} else {
					resp.Msg = err.Error()
				}
			}
		} else {
			resp.Msg = "手机号错误"
		}
	} else {
		resp.Msg = err.Error()
	}
	c.Render().JSON(resp)
}

func (c *ApiController) GetTerm() {

}

func (c *ApiController) GetSftp() {

}

func (c *ApiController) GetUserinfo() {
	var resp Apiform.Resp
	newToken := c.Ctx().Value("token").(string)
	if newToken != "" { //更新Token逻辑
		resp.Token = newToken
	}
	uid := c.Ctx().Value("uid").(int64)
	var limit Apiform.Slist
	if uid > 0 {
		if err := c.Ctx().BindForm(&limit); err == nil {

			var list Apiform.List_resp
			var server tables.SSHServer
			server.BindUser = uid
			count, _ := c.Orm.Table(&tables.SSHServer{}).Where("bind_user = ?", uid).Limit(limit.Limit, (limit.Page-1)*limit.Limit).FindAndCount(&list.List)
			list.Count = uint(count)
			resp.Code = 200
			resp.Data = list
			resp.Msg = "查询成功"
		} else {
			resp.Code = errcode.C_from_err
			resp.Msg = err.Error()
		}
	} else {
		resp.Code = errcode.S_Verify_err
		resp.Msg = "Token信息错误"
	}
	c.Render().JSON(resp)
}

func (c *ApiController) PostNickname() {

}

func (c *ApiController) PostAddser() {
	var resp Apiform.Resp
	newToken := c.Ctx().Value("token").(string)
	if newToken != "" {
		resp.Token = newToken
	}
	uid := c.Ctx().Value("uid").(int64)
	var info Apiform.Addser
	resp.Code = errcode.C_from_err
	resp.Msg = "数据错误"
	if err := c.Ctx().BindForm(&info); err == nil {
		id, _ := c.Orm.InsertOne(&tables.SSHServer{Ip: info.Ip, Port: info.Port, Username: info.Username, Password: info.Password, Nickname: info.Nickname, BindUser: uid})
		if id > 0 {
			resp.Code = errcode.C_nil_err
			resp.Msg = "保存成功"
		} else {
			resp.Code = errcode.S_Db_err
			resp.Msg = "保存失败"
		}
	} else {
		resp.Msg = err.Error()
	}
	c.Render().JSON(resp)
}

func (c *ApiController) PostRepass() {
	var resp Apiform.Resp
	var edit Apiform.EditPass
	newToken := c.Ctx().Value("token").(string)
	if newToken != "" { //更新Token逻辑
		resp.Token = newToken
	}
	uid := c.Ctx().Value("uid").(int64)

	if c.Ctx().BindForm(&edit) == nil {
		//server.Nickname = nickname
		var server tables.SSHServer
		server.Id = edit.ID
		server.BindUser = uid
		result, _ := c.Orm.Table(&tables.SSHServer{}).Where("id = ?", server.Id).Where("bind_user = ?", uid).Update(tables.SSHServer{Password: edit.Password})
		if result > 0 {
			resp.Code = errcode.C_nil_err
			resp.Msg = "保存成功"
		} else {
			resp.Code = errcode.S_Db_err
			resp.Msg = "修改失败"
		}
	} else {
		resp.Code = errcode.C_from_err
		resp.Msg = "提交字段错误"
	}
	c.Render().JSON(resp)
}

func (c *ApiController) PostDelete() {
	var resp Apiform.Resp
	var del Apiform.Edit
	newToken := c.Ctx().Value("token").(string)
	if newToken != "" { //更新Token逻辑
		resp.Token = newToken
	}
	uid := c.Ctx().Value("uid").(int64)
	if err := c.Ctx().BindForm(&del); err == nil {
		var server tables.SSHServer
		server.Id = del.ID
		server.BindUser = uid
		result, _ := c.Orm.Where("id = ?", server.Id).Where("bind_user = ?", uid).Delete(&tables.SSHServer{})
		if result > 0 {
			resp.Code = errcode.C_nil_err
			resp.Msg = "删除成功"
		} else {
			resp.Code = errcode.S_Db_err
			resp.Msg = "操作失败"
		}
	} else {
		resp.Code = errcode.C_from_err
		resp.Msg = err.Error()
	}
	c.Render().JSON(resp)
}

func (c *ApiController) PostGetterm() {
	var resp Apiform.Resp
	var term Apiform.GetTerm
	newToken := c.Ctx().Value("token").(string)
	if newToken != "" { //更新Token逻辑
		resp.Token = newToken
	}
	uid := c.Ctx().Value("uid").(int64)
	resp.Code = errcode.C_from_err
	resp.Msg = "表单错误"
	if c.Ctx().BindForm(&term) == nil {
		var server tables.SSHServer
		server.Id = term.ID
		server.BindUser = uid
		result, _ := c.Orm.Where("id = ?", server.Id).Where("bind_user = ?", uid).Get(&server)
		if result {
			c.Orm.Where("id = ?", server.Id).Where("bind_user = ?", uid).Update(&tables.SSHServer{BeforeTime: time.Now()})
			sid, err := term.Decode(server)
			if err == nil {
				resp.Code = errcode.C_nil_err
				resp.Data = sid
				resp.Msg = "OK"
			} else {
				resp.Code = errcode.S_Verify_err
				resp.Msg = "秘钥解密失败"
			}
		} else {
			resp.Code = errcode.S_Db_err
			resp.Msg = "服务器信息检索失败"
		}
	}
	c.Render().JSON(resp)
}
