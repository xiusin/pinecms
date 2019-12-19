package backend

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
)

func clearMenuCache(cache cache.ICache, xorm *xorm.Engine) {
	var roles []*tables.IriscmsAdminRole
	var menus []*tables.IriscmsMenu

	xorm.Where("parentid = ?", 0).Find(&menus)
	xorm.Find(&roles)
	for _, role := range roles {
		for _, menu := range menus {
			cacheKey := fmt.Sprintf(controllers.CacheAdminMenuByRoleIdAndMenuId, role.Roleid, menu.Id)
			cache.Delete(cacheKey)
		}

	}
}
