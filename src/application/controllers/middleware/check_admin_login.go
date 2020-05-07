package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

func CheckAdminLoginAndAccess(xorm *xorm.Engine, iCache cache.AbstractCache) pine.Handler {
	return func(this *pine.Context) {
		this.Render().ViewData("staticDir", "/assets/backend/static")
		this.Render().ViewData("baseDir", "/assets/backend")

		if strings.Contains(this.Path(), "login") {
			this.Next()
			return
		}
		aid, _ := strconv.Atoi(this.Session().Get("adminid"))
		roleId, _ := strconv.Atoi(this.Session().Get("roleid"))
		if aid > 0 && roleId > 0 {
			//放置一些数据到全局可取
			this.Set("adminid", int64(aid))
			this.Set("roleid", int64(roleId))
			this.Set("username", this.Session().Get("username"))

			pathString := strings.Split(strings.Trim(this.Path(), "/"), "/")
			//public 或check 开始的路由不检测权限
			if !(len(pathString) == 3 && (strings.Contains(pathString[2], "public-") ||
				strings.Contains(pathString[2], "check-") || pathString[1] == "index")) {
				if roleId > 1 && CheckPriv(this, xorm, iCache) == false {
					helper.Ajax("您没有操作权限", 1, this)
					return
				}
			}
			this.Next()
		} else {
			this.Redirect("/b/login/index", 302)
			return
		}
	}
}

//检查权限
func CheckPriv(this *pine.Context, xorm *xorm.Engine, cache cache.AbstractCache) bool {
	pathinfo := strings.Split(strings.Trim(this.Path(), "/"), "/")
	roleId, err := strconv.Atoi(this.Session().Get("roleid"))
	if err != nil || len(pathinfo) < 3 {
		return false
	}
	// 用户权限放到缓存内
	key := fmt.Sprintf(controllers.CacheAdminPriv, roleId)
	data, err := cache.Get(key)
	ha := map[string]struct{}{}
	if err != nil || json.Unmarshal(data, &ha) != nil {
		var privs []*tables.AdminRolePriv
		// 读取所有用户权限
		err := xorm.Where("roleid = ?", roleId).Find(&privs)
		if err != nil {
			pine.Logger().Error(helper.GetCallerFuncName(), err)
		}
		for _, priv := range privs {
			ha[strings.ToLower(priv.C+"-"+priv.A)] = struct{}{}
		}
		strs, err := json.Marshal(&ha)
		if err != nil {
			pine.Logger().Error(helper.GetCallerFuncName(), "编码json失败", err)
		}
		data = strs
		if err := cache.Set(key, strs); err != nil {
			pine.Logger().Error("保存权限缓存失败", err.Error())
		}
	}

	if _, ok := ha[strings.ToLower(pathinfo[1]+"-"+pathinfo[2])]; ok {
		return true
	}
	return false
}
