package middleware

import (
	"fmt"
	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

func Casbin(engine *xorm.Engine) pine.Handler {
	enforcer := casbin.NewEnforcer(
		helper.GetRootPath("resources/configs/rbac_models.conf"), xd.NewAdapterByEngine(engine))
	di.Set(controllers.ServiceCasbinEnforcer, func(builder di.AbstractBuilder) (interface{}, error) {
		return enforcer, nil
	}, true)
	addPolicy(engine, enforcer)
	return func(ctx *pine.Context) {
		roleidIntf := ctx.Value("adminid")
		if roleidIntf != nil { // 存在角色ID 则检查权限
			roleId := roleidIntf.(int64)
			fmt.Println("roleid", roleId, "path", ctx.Path())
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
				enforcer.AddPolicy(fmt.Sprintf("%d", role.Id), priv.C, priv.A, fmt.Sprintf("%d", priv.MenuId))
			}
		}
		enforcer.SavePolicy()
	} else {
		_ = enforcer.LoadPolicy()
	}
}
