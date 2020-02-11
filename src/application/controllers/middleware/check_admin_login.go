package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/controllers"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"

	"github.com/kataras/iris/v12"
)

func CheckAdminLoginAndAccess(sess *sessions.Sessions, cache cache.ICache, xorm *xorm.Engine) func(this iris.Context) {
	return func(this iris.Context) {
		sess := sess.Start(this)
		if strings.Contains(this.Path(), "login") {
			this.Next()
			return
		}
		aid, err := sess.GetInt64("adminid") //检测是否设置过session
		if err != nil {
			golog.Debug("check login failed", err)
		}
		if aid == -1 {
			sess.Clear()
			this.Redirect("/b/login/index", 302)
		} else {
			//检查权限
			roleId, _ := sess.GetInt64("roleid")
			//放置一些数据到全局可取
			this.Values().Set("adminid", aid)
			this.Values().Set("roleid", roleId)
			this.Values().Set("username", sess.Get("username"))
			if roleId == -1 {
				sess.Clear()
				this.Redirect("/b/login/index", 302)
				this.StopExecution()
				return
			}
			pathString := strings.Split(strings.Trim(this.Path(), "/"), "/")
			//public 或check 开始的路由不检测权限
			if len(pathString) == 3 && (strings.Contains(pathString[2], "public-") || strings.Contains(pathString[2], "check-") || pathString[1] == "index") {
				this.Next()
			} else {
				if roleId > 1 && CheckPriv(this, sess, cache, xorm) == false {
					helper.Ajax("您没有操作权限", 1, this)
					this.StopExecution()
					return
				}
				//go ManageLog(this, xorm)
			}
			this.Next()
		}
	}
}

//检查权限
func CheckPriv(this iris.Context, sess *sessions.Session, cache cache.ICache, xorm *xorm.Engine) bool {
	pathinfo := strings.Split(strings.Trim(this.Path(), "/"), "/")
	roleId, err := sess.GetInt64("roleid")
	if err != nil || len(pathinfo) < 3 {
		return false
	}
	// 用户权限放到缓存内
	key := fmt.Sprintf(controllers.CacheAdminPriv, roleId)
	data := cache.Get(key)
	ha := map[string]struct{}{}
	if data == "" || json.Unmarshal([]byte(data), &ha) != nil {
		var privs []*tables.IriscmsAdminRolePriv
		// 读取所有用户权限
		err := xorm.Where("roleid = ?", roleId).Find(&privs)
		if err != nil {
			golog.Error(helper.GetCallerFuncName(), err)
		}
		for _, priv := range privs {
			ha[strings.ToLower(priv.C+"-"+priv.A)] = struct{}{}
		}
		strs, err := json.Marshal(&ha)
		if err != nil {
			golog.Error(helper.GetCallerFuncName(), "编码json失败", err)
		}
		data = string(strs)
		if cache.Set(key, strs) != nil {
			golog.Error("保存权限缓存失败")
		}
	}

	if _, ok := ha[strings.ToLower(pathinfo[1]+"-"+pathinfo[2])]; ok {
		return true
	}
	return false
}

//func ManageLog(this iris.Context, xorm *xorm.Engine) {
//	pathinfo := strings.Split(strings.Trim(this.Path(), "/"), "/")
//	if len(pathinfo) == 3 {
//		aid, _ := this.Values().Get("adminid").(int64)
//		ip := this.RemoteAddr()
//		username := this.Values().GetString("username")
//		time := helper.NowDate("Y-m-d H:i:s")
//		uri := this.Request().RequestURI
//		log := tables.IriscmsLog{
//			Ip:          ip,
//			Username:    username,
//			Querystring: uri,
//			Time:        time,
//			Controller:  pathinfo[1],
//			Action:      pathinfo[2],
//			Userid:      aid,
//		}
//		_, err := xorm.Insert(log)
//		if err != nil {
//			golog.Error(helper.GetCallerFuncName(), err)
//		}
//	}
//}
