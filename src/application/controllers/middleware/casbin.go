package middleware

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

func Casbin(engine *xorm.Engine, conf string) pine.Handler {
	enforcer := casbin.NewEnforcer(helper.GetRootPath(conf), xd.NewAdapterByEngine(engine))
	di.Set(controllers.ServiceCasbinEnforcer, func(builder di.AbstractBuilder) (interface{}, error) {
		return enforcer, nil
	}, true)
	addPolicy(engine, enforcer)
	return func(ctx *pine.Context) {
		roleidIntf := ctx.Value("adminid")
		if roleidIntf != nil {
			// TODO 添加缓存, 记录账户权限
			var admin = &tables.Admin{}
			exist, _ := engine.Where("id = ?", roleidIntf).Get(admin)
			pathString := strings.Split(strings.Trim(ctx.Path(), "/"), "/")

			if exist && len(pathString) >= 3 && pathString[0] == "v2" {
				roles := admin.RoleIdList

				var passable bool
				for _, role := range roles {
					if passable = enforcer.Enforce(fmt.Sprintf("%d", role), pathString[1], pathString[2]); passable {
						ctx.Next()
						return
					}
				}
			}
			helper.Ajax("无节点操作权限", 1, ctx)
			return
		}
		ctx.Next()
	}
}

// 根据角色注入权限
func addPolicy(engine *xorm.Engine, enforcer *casbin.Enforcer) {
	if count, _ := engine.Table(&xd.CasbinRule{}).Count(); count == 0 {
		var roles []tables.AdminRole
		engine.Find(&roles)
		for _, role := range roles {
			var privs []tables.AdminRolePriv
			engine.Where("roleid = ?", role.Id).Find(&privs)
			for _, priv := range privs {
				enforcer.AddPolicy(fmt.Sprintf("%d", role.Id), fmt.Sprintf("%d", priv.MenuId))
			}
		}
		enforcer.SavePolicy()
	} else {
		_ = enforcer.LoadPolicy()
	}
}
