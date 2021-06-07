package backend

import (
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/common/storage"
	"math/rand"
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

func clearMenuCache(cache cache.AbstractCache, xorm *xorm.Engine) {
	var roles []*tables.AdminRole
	var menus []*tables.Menu
	xorm.Where("parentid = ?", 0).Find(&menus)
	xorm.Find(&roles)
	for _, role := range roles {
		for _, menu := range menus {
			cacheKey := fmt.Sprintf(controllers.CacheAdminMenuByRoleIdAndMenuId, role.Id, menu.Id)
			cache.Delete(cacheKey)
		}
	}
}

func getEditor(field, val, attrs string, required bool, formName, defaultVal, RequiredTips string) string {
	rid := "component_editor_um_" + strconv.Itoa(rand.Int())
	if RequiredTips == "" {
		RequiredTips = formName + "必须填写"
	}
	var requiredFunc = ""
	if required {
		requiredFunc = `editors.push(function(){ if (!` + rid + `.hasContents()) {$('#` + rid + `_tip').html("` + RequiredTips + `"); return false; } $('#` + rid + `_tip').html(''); return true; });`
	}
	return `<textarea id="` + rid + `" ` + attrs + ` style="width:100%;height:360px" name="` + field + `">` + val + `</textarea>
<div id='` + rid + `_tip' class='errtips'></div>
<script>var ` + rid + ` = UE.getEditor('` + rid + `'); ` + requiredFunc + `</script>`
}

func getStorageEngine(settingData map[string]string) storage.Uploader {
	uploadDir := settingData["UPLOAD_DIR"]
	urlPrefixDir := settingData["UPLOAD_URL_PREFIX"]
	engine := settingData["UPLOAD_ENGINE"]
	var uploader storage.Uploader
	switch engine {
	case "OSS存储":
		uploader = storage.NewOssUploader(settingData)
	default:
		uploader = storage.NewFileUploader(urlPrefixDir, uploadDir)
	}
	return uploader
}

func strFirstToUpper(str string) string {
	temp := strings.Split(strings.ReplaceAll(str, "_", "-"), "-")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

func inArray(val interface{}, vals []interface{}) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}
	return false
}

func ucwords(str string) string {
	str = strings.ToLower(str)
	vv := []rune(str)
	for i := 0; i < len(str); i++ {
		if i == 0 {
			vv[i] -= 32
		}
	}
	return string(vv)
}

func parseParam(ctx *pine.Context, param interface{}) error {
	return ctx.BindJSON(param)
}
