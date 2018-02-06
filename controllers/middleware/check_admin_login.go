package middleware

import (
"fmt"
"strings"


"iriscms/controllers/backend/helper"
"iriscms/models/tables"

"github.com/go-xorm/xorm"
"github.com/kataras/iris/context"
"github.com/kataras/iris/sessions"

)

func CheckAdminLoginAndAccess(sess *sessions.Sessions,xorm *xorm.Engine) func(this context.Context) {
	return func(this context.Context) {
		sess := sess.Start(this)
		if strings.Contains(this.Path(), "login") {
			this.Next()
			return
		}
		aid, err := sess.GetInt64("adminid") //检测是否设置过session
		if err != nil || aid == -1 {
			sess.Clear()
			this.Redirect("/b/login/index", 302)
			return
		} else {
			//检查权限
			roleId, err := sess.GetInt64("roleid")
			//放置一些数据到全局可取

			this.Values().Set("adminid", int64(aid))
			this.Values().Set("roleid", int64(roleId))
			this.Values().Set("username", sess.Get("username"))
			if err != nil || roleId == -1 {
				sess.Clear()
				this.Redirect("/b/login/index", 302)
				return
			}

			pathString := strings.Split(strings.Trim(this.Path(), "/"), "/")
			//public 或check 开始的路由不检测权限
			if len(pathString) == 3 && (strings.Contains(pathString[2], "public-") || strings.Contains(pathString[2], "check-") || pathString[1] == "index") {
				this.Next()
			} else {
				if roleId > 1 && CheckPriv(this, sess,xorm) == false {
					helper.Ajax("您没有操作权限", 1, this)
				} else {
					ManageLog(this,xorm)
					this.Next()
				}
			}
		}
		this.Next()
	}
}

//检查权限
func CheckPriv(this context.Context, sess *sessions.Session,xorm *xorm.Engine) bool {
	pathinfo := strings.Split(strings.Trim(this.Path(), "/"), "/")
	roleId, err := sess.GetInt64("roleid")
	if err != nil || len(pathinfo) < 3 {
		return true
	}
	has, _ := xorm.Get(&tables.IriscmsAdminRolePriv{C: pathinfo[1], A: pathinfo[2], Roleid: int64(roleId)})
	if !has {
		return false
	}
	return true
}

func ManageLog(this context.Context,xorm *xorm.Engine) {
	pathinfo := strings.Split(strings.Trim(this.Path(), "/"), "/")
	if len(pathinfo) == 3 {
		aid, _ := this.Values().Get("adminid").(int64)
		ip := this.RemoteAddr()
		username := this.Values().GetString("username")
		time := helper.NowDate("Y-m-d H:i:s")
		uri := string(this.Request().RequestURI)
		log := tables.IriscmsLog{
			Ip:          ip,
			Username:    username,
			Querystring: uri,
			Time:        time,
			Controller:  pathinfo[1],
			Action:      pathinfo[2],
			Userid:      aid,
		}
		_, err := xorm.Insert(log)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
