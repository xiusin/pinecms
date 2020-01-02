package backend

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"
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
	h := "<form><input type='hidden' value='" + documentModel.Table + "'><table cellpadding=\"2\" class=\"dialogtable\" style=\"width: 100%;\">"
	for _, field := range fields {
		h += `<tr><td style="width: 150px;">` + field.FormName + `:</td><td>`
		attrs := []string{"name='" + field.TableField + "'"}
		var options []string
		var domOpts []string
		//if field.Id == 2 {
		//	options = append(options, "multiline:true")
		//}

		if field.Required == 1 {
			options = append(options, "required:true")
			domOpts = append(domOpts, "required")
			if field.RequiredTips != "" {
				options = append(options, "missingMessage:'"+field.RequiredTips+"'")
			}
		}

		if field.Validator != "" {
			options = append(options, "validType:"+field.Validator)
			options = append(options, "invalidMessage:'"+field.RequiredTips+"'")
		}

		isEditor := strings.HasPrefix(field.Html, "<editor")
		isImageUpload := strings.HasPrefix(field.Html, "<images")
		isMulImageUpload := strings.HasPrefix(field.Html, "<mul-images")
		if len(options) > 0 {
			if isEditor {
				attrs = append(attrs, domOpts...)
			} else {
				attrs = append(attrs, `data-options="`+strings.Join(options, ", ")+`"`)
			}
		}

		value := ""
		if isEditor {
			field.Html = getEditor(value, strings.Join(attrs, " "), true)
		} else if isImageUpload {
			field.Html = helper.SiginUpload("", field.TableField)
		} else if isMulImageUpload {
			field.Html = helper.MultiUpload([]string{})
		} else {
			field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), 1)
			field.Html = strings.Replace(field.Html, "{{value}}", value, 1)
		}
		// todo 匹配loop和数据源接入
		h += field.Html
		h += "</td></tr>"

	}
	h += "</form>"
	return h
}

func getEditor(val, attrs string, required bool) string {
	rid := strconv.Itoa(rand.Int())
	return `<textarea id="component_editor_um_` + rid + `" ` + attrs + ` style="width:100%;height:360px" >` + val + `</textarea>
<script> UM.getEditor('component_editor_um_` + rid + `');</script>`
}
