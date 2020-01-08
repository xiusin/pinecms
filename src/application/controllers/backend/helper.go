package backend

import (
	"encoding/json"
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
	if documentModel == nil || documentModel.Id < 1 {
		panic("模型不存在")
	}
	fields := models.NewDocumentFieldDslModel(orm).GetList(mid)
	h := "<form method='POST' id='content_page_form' enctype='multipart/form-data'>" +
		"	<input type='hidden' value='" + documentModel.Table + "'>" +
		"		<table cellpadding='2' class='dialogtable' style='width: 100%;'>"
	for _, field := range fields {
		h += `			<tr><td style="width: 150px;">` + field.FormName + `:</td><td>
`
		if strings.Contains(field.Html, "easyui-") {
			h += easyUIComponents(&field)
		} else {
			h += domOrCustomTagComponents(&field)
		}
		h += "</td></tr>"
	}
	h += "</form>"
	return h
}
func domOrCustomTagComponents(field *tables.IriscmsDocumentModelDsl) string {
	attrs := []string{"name='" + field.TableField + "'"}
	isEditor := strings.HasPrefix(field.Html, "<editor")
	isImageUpload := strings.HasPrefix(field.Html, "<images")
	isMulImageUpload := strings.HasPrefix(field.Html, "<mul-images")
	value := ""
	if isEditor {
		field.Html = getEditor(value, strings.Join(attrs, " "), true)
	} else if isImageUpload {
		field.Html = helper.SiginUpload("", field.TableField)
	} else if isMulImageUpload {
		field.Html = helper.MultiUpload([]string{}, 5)
	} else {
		field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), 1)
		field.Html = strings.Replace(field.Html, "{{value}}", value, 1)
	}
	if field.RequiredTips != "" {
		field.Html += field.RequiredTips
	}
	return field.Html
}

func easyUIComponents(field *tables.IriscmsDocumentModelDsl) string {
	var options []string
	attrs := []string{"name='" + field.TableField + "'"}
	if strings.Contains(field.Html, "multiline") {
		options = append(options, "multiline:true")
		field.Html = strings.Replace(field.Html, "multiline", "", 1)
	}
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

	if field.Datasource != "" {
		// 多选  单选 下拉 联动需要在这里加载数据源
		options = append(options, "valueField:'value'")
		options = append(options, "textField:'label'")
		var datas []struct {
			Value string `json:"value"`
			Label string `json:"label"`
		}
		var dataSourceJson interface{}
		if strings.HasPrefix(field.Datasource, "[") || strings.HasPrefix(field.Datasource, "{") {
			err := json.Unmarshal([]byte(field.Datasource), &dataSourceJson)
			if err == nil { // 能解出来json // 只支持kv格式,  符合ComboBox的JSON规范
				switch dataSourceJson.(type) {
				case map[string]interface{}:
					for k, v := range dataSourceJson.(map[string]interface{}) {
						datas = append(datas, struct {
							Value string `json:"value"`
							Label string `json:"label"`
						}{Value: k, Label: fmt.Sprintf("%s", v)})
					}
				case []interface{}:
					for _, v := range dataSourceJson.([]interface{}) {
						datas = append(datas, struct {
							Value string `json:"value"`
							Label string `json:"label"`
						}{Value: fmt.Sprintf("%s", v), Label: fmt.Sprintf("%s", v)})
					}
				}
			} else {
				panic("解码失败:" + err.Error())
			}
		}
		if field.FieldType == 5 || field.FieldType == 6 {
			if len(datas) == 0 {
				options = append(options, "url: '"+field.Datasource+"'")
			} else {
				jsonstr, err := json.Marshal(&datas)
				if err == nil {
					s := strings.Replace(string(jsonstr), `"`, "'", -1)
					s = strings.Replace(s, `'value'`, "value", -1)
					s = strings.Replace(s, `'label'`, "label", -1)
					options = append(options, "onSelect:function(record){ console.log(record) }")
					options = append(options, "data: "+s)
				} else {
					panic("序列化数据失败:" + err.Error())
				}
			}
		} else if field.FieldType == 7 || field.FieldType == 8 {
			var htm = ""
			for _, item := range datas {
				if item.Value == field.Default {
					field.Html = strings.Replace(field.Html, "{{default}}", "checked", 1)
				} else {
					field.Html = strings.Replace(field.Html, "{{default}}", "", 1)
				}
				htm += strings.Replace(field.Html, "{{value}}", item.Value, -1) + strings.Repeat("&nbsp;", 3) + item.Label + strings.Repeat("&nbsp;", 8)
			}
			field.Html = htm
		}
	}
	if len(options) > 0 {
		attrs = append(attrs, `data-options="`+strings.Join(options, ", ")+`"`)
	}
	value := ""
	field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), -1)
	field.Html = strings.Replace(field.Html, "{{value}}", value, -1)
	return field.Html
}

func getEditor(val, attrs string, required bool) string {
	rid := strconv.Itoa(rand.Int())
	return `<textarea id="component_editor_um_` + rid + `" ` + attrs + ` style="width:100%;height:360px" >` + val + `</textarea><script> UM.getEditor('component_editor_um_` + rid + `');</script>`
}
