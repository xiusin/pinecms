package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"strings"
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

// 构建模型表单
func buildModelForm(orm *xorm.Engine, mid int64) string {
	model := models.NewDocumentModel(orm)
	documentModel := model.GetByID(mid)
	if documentModel == nil {
		panic("模型不存在")
	}
	fields := models.NewDocumentFieldDslModel(orm).GetList(mid)
	h := "<form><input type='hidden' value='"+documentModel.Table+"'><table cellpadding=\"5\" style=\"width: 100%;\">"
	for _, field := range fields {
		h += `<tr><td style="width: 150px;">` +field.FormName+ `:</td><td>`
		// 判断属性与组件
		attrs := []string{"name='"+field.TableField+"'"}
		var options []string
		if field.Required == 1 {
			options = append(options, "required:true")
			if field.RequiredTips != "" {
				options = append(options, "missingMessage:'"+field.RequiredTips+"'")
			}
		}

		if field.Validator != "" {
			options = append(options, "validType:"+field.Validator)
			options = append(options, "invalidMessage:'"+field.RequiredTips+"'")
		}
		if len(options) >0 {
			attrs = append(attrs, `data-options="`+strings.Join(options, ", ")+`"`)
		}
		value := ""
		field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), 1)
		field.Html = strings.Replace(field.Html, "{{value}}", value, 1)

		// todo 匹配loop和数据源接入

		h += field.Html
		h += "</td></tr>"

	}
	h += "</form>"
	return h
}
