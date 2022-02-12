package webssh

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/Apiform"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/common"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/common/core"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/common/sftp_clients"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/errcode"
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"log"
	"time"
)

var upGrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

type AuthMsg struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type sftpReq struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type sftpResp struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// MFA https://github.com/renmcc/koko/blob/ab95a0a40d2b851424aab06125483957cf64a219/pkg/auth/server.go
var mfaInstruction = "Please enter 6 digits."
var mfaQuestion = "[MFA auth]: "


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
	var auth Apiform.WsAuth
	if c.Ctx().BindForm(&auth) != nil {
		return
	}
	if err := upGrader.Upgrade(c.Ctx().RequestCtx, func(wsConn *websocket.Conn) {
		cols, rows := 120, 32
		var serInfo Apiform.SerInfo //接收反序列化数据
		for {
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				pine.Logger().Print(err)
				_ = wsConn.Close()
				return
			}
			msgObj := AuthMsg{}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				log.Println("Auth : unmarshal websocket message failed:", string(wsData))
				pine.Logger().Print(err)
				continue
			}
			token := msgObj.Token
			claims, err := common.ParseToken(token)
			valid := claims.Valid()
			if valid != nil || err != nil {
				fmt.Println("身份验证失败")
				_ = wsConn.WriteMessage(websocket.BinaryMessage, []byte("身份验证失败\r\n"))
				_ = wsConn.Close()
				return
			}
			sInfo, err := helper.AbstractCache().Get(auth.Sid)
			if err != nil || len(sInfo) == 0 {
				fmt.Println("连接超时，请重试！")
				_ = wsConn.WriteMessage(websocket.BinaryMessage, []byte("连接超时，请重试！\r\n"))
				_ = wsConn.Close()
				return
			}
			if json.Unmarshal(sInfo, &serInfo) != nil {
				fmt.Println("服务器信息获取失败，请重试！")
				_ = wsConn.WriteMessage(websocket.BinaryMessage, []byte("服务器信息获取失败，请重试！\r\n"))
				_ = wsConn.Close()
				return
			}
			//log.Println(ser_info)
			if claims.Userid != serInfo.BindUser { //验证权限
				fmt.Println("权限验证失败，请重试！")
				_ = wsConn.WriteMessage(websocket.BinaryMessage, []byte("权限验证失败，请重试！\r\n"))
				_ = wsConn.Close()
				return
			}
			break
			//break
		}
		client, err := core.NewSshClient(core.Server{Ip: serInfo.Ip, Port: serInfo.Port, User: serInfo.Username, Passwd: serInfo.Password})
		if err != nil {
			pine.Logger().Print(err)
			return
		}

		defer client.Close()
		ssConn, err := core.NewSshConn(cols, rows, client) //加入sftp客户端
		if err != nil {
			pine.Logger().Print(err)
			return
		}
		sftp_clients.Client.Lock()
		sftp_clients.Client.C[auth.Sid] = &sftp_clients.MyClient{uint(serInfo.BindUser), ssConn.SftpClient}
		sftp_clients.Client.Unlock()
		defer func() {
			sftp_clients.Client.Lock()
			delete(sftp_clients.Client.C, auth.Sid) //释放SFTP客户端
			sftp_clients.Client.Unlock()
		}()
		defer ssConn.Close()
		quitChan := make(chan bool, 3)

		// most messages are ssh output, not webSocket input
		go ssConn.ReceiveWsMsg(wsConn, quitChan)
		go ssConn.SendComboOutput(wsConn, quitChan)
		go ssConn.SessionWait(quitChan)

		<-quitChan //任意协程退出则结束
		log.Println("websocket finished")
	}); err != nil {
		var resp Apiform.Resp
		resp.Code = errcode.C_from_err
		resp.Msg = err.Error()
		c.Render().JSON(resp)
	}
}

func (c *ApiController) GetSftp() {
	var auth Apiform.WsAuth

	if err := c.Ctx().BindForm(&auth); err != nil {
		var resp Apiform.Resp
		resp.Code = errcode.C_from_err
		resp.Msg = err.Error()
		pine.Logger().Print(err)
		c.Render().JSON(resp)
		return
	}

	if err := upGrader.Upgrade(c.Ctx().RequestCtx, func(wsConn *websocket.Conn) {
		for {
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				_ = wsConn.Close()
				return
			}
			//unmashal bytes into struct
			msgObj := sftpReq{}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				log.Println("Auth : unmarshal websocket message failed:", string(wsData))
				continue
			}
			respMsg := sftpResp{}
			token := msgObj.Token
			claims, err := common.ParseToken(token)
			valid := claims.Valid()
			if valid != nil || err != nil {
				respMsg.Code = errcode.S_auth_fmt_err
				respMsg.Msg = "身份令牌校验不通过"
				respMsg.Data = err.Error()
				msg, _ := json.Marshal(respMsg)
				if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("sftp token fmt err:", err)
				}
				_ = wsConn.Close()
				return
			}
			if claims.Userid != int64(sftp_clients.Client.C[auth.Sid].Uid) { //身份与缓存不符合
				respMsg.Code = errcode.S_auth_fmt_err
				respMsg.Msg = "用户权限不通过"
				respMsg.Data = err.Error()
				msg, _ := json.Marshal(respMsg)
				if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("sftp server_user err:", err)
				}
				_ = wsConn.Close()
				return
			}

			path, err := sftp_clients.Client.C[auth.Sid].Sftp.Getwd()
			if err != nil {
				respMsg.Code = errcode.S_send_err
				respMsg.Type = "connect"
				respMsg.Msg = "SFTP连接失败"
				msg, _ := json.Marshal(respMsg)
				if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("sftp connect err:", err)
				}
				return
			}

			respMsg.Code = 200
			respMsg.Type = "connect"
			respMsg.Msg = "连接成功"
			respMsg.Data = path
			msg, _ := json.Marshal(respMsg)
			if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("sftp return err:", err)
				return
			}

			break
			//break
		}
		quitChan := make(chan bool, 2)
		go sftp_clients.Client.C[auth.Sid].ReceiveWsMsg(wsConn, quitChan)
		<-quitChan //任意协程退出则结束
		fmt.Println("Sftp Exit")
		log.Println("sftp websocket finished")
	}); err != nil {
		var resp Apiform.Resp
		resp.Code = errcode.C_from_err
		resp.Msg = err.Error()
		c.Render().JSON(resp)
	}
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
	if err := c.Ctx().BindForm(&term); err == nil {
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
	} else {
		resp.Code = errcode.S_Verify_err
		resp.Msg = err.Error()
	}
	c.Render().JSON(resp)
}
