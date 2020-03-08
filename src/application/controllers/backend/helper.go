package backend

import (
	"encoding/json"
	"fmt"
	"github.com/xiusin/pinecms/src/common/storage"
	"github.com/xiusin/pine/cache"
	"math/rand"
	"strconv"
	"strings"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

//todo 清理缓存， 结合最终
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
func buildModelForm( mid int64, data map[string]string) string {
	model := models.NewDocumentModel()
	documentModel := model.GetByID(mid)
	if documentModel == nil || documentModel.Id < 1 {
		panic("模型不存在")
	}
	idInput := ""
	buttonTxt := "发布"
	if id, ok := data["id"]; ok {
		idInput = "<input type='hidden' name='id' value='" + id + "'>"
		buttonTxt = "更新"
	}
	fields := models.NewDocumentFieldDslModel().GetList(mid)
	h := "<form method='POST' id='content_page_form' enctype='multipart/form-data'>" +
		"<input type='hidden' name='table_name' value='" + documentModel.Table + "'>" +
		"<input type='hidden' name='mid' value='" + strconv.Itoa(int(mid)) + "'>" +
		idInput +
		"<table cellpadding='2' class='dialogtable' style='width: 100%;'>"
	for _, field := range fields {
		h += `<tr><td style="width: 100px;text-align:right;">` + field.FormName + `：</td><td>`
		val, _ := data[field.TableField]	// 读取字段值
		if strings.Contains(field.Html, "easyui-") {
			h += easyUIComponents(&field, val)
		} else {
			h += domOrCustomTagComponents(&field, val)
		}
		h += "</td></tr>"
	}
	if data["status"] == "1" {
		h += "<tr><td style='text-align:right;'>状态：</td><td> <input class='easyui-switchbutton' checked name='status'></td></tr>"
	} else {
		h += "<tr><td style='text-align:right;'>状态：</td><td> <input class='easyui-switchbutton' name='status'></td></tr>"
	}
	h += `<tr><td colspan=2><a href="javascript:void(0);" onclick="submitForm()" class="easyui-linkbutton">`+buttonTxt+`</a></td></tr></table></form>`
	return h
}
func domOrCustomTagComponents(field *tables.IriscmsDocumentModelDsl, val string) string {
	attrs := []string{"name='" + field.TableField + "'"}
	isEditor := strings.HasPrefix(field.Html, "<editor")
	isImageUpload := strings.HasPrefix(field.Html, "<images")
	isMulImageUpload := strings.HasPrefix(field.Html, "<mul-images")
	isTags := strings.HasPrefix(field.Html, "<tags")
	if isEditor {
		field.Html = getEditor(field.TableField, val, strings.Join(attrs, " "), field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isImageUpload {
		field.Html = helper.SiginUpload(field.TableField, val, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isMulImageUpload {
		field.Html = helper.MultiUpload(field.TableField, strings.Split(val, ","), 5, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isTags{
		field.Html = helper.Tags(field.TableField, val, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else {
		field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), 1)
		field.Html = strings.Replace(field.Html, "{{value}}", val, 1)
	}
	return field.Html
}

func easyUIComponents(field *tables.IriscmsDocumentModelDsl, val string) string {
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
		} else {
			options = append(options, "missingMessage:'"+field.FormName+"必须填写'")
		}
	}
	if field.Validator != "" {
		options = append(options, "validType:"+field.Validator)
		options = append(options, "invalidMessage:'"+field.FormName+"输入数据内容无效'")
	}

	if field.Datasource != "" {

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
			options = append(options, "valueField:'value'")
			options = append(options, "textField:'label'")
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
			options = append(options, "valueField:'value'")
			options = append(options, "textField:'label'")
			defaultVals := strings.Split(field.Default, "|")
			var htm = ""
			var htmlTmp string
			for _, item := range datas {
				if field.FieldType == 8 {
					for _, v := range defaultVals {
						htmlTmp = field.Html
						if v == item.Value {
							htmlTmp = strings.Replace(htmlTmp, "{{default}}", "checked", 1)
							break
						} else {
							htmlTmp = strings.Replace(htmlTmp, "{{default}}", "", 1)
						}
					}
				} else {
					htmlTmp = field.Html
					if item.Value == field.Default {
						htmlTmp = strings.Replace(htmlTmp, "{{default}}", "checked", 1)
					} else {
						htmlTmp = strings.Replace(htmlTmp, "{{default}}", "", 1)
					}
				}
				htm += strings.Replace(htmlTmp, "{{value}}", item.Value, -1) + strings.Repeat("&nbsp;", 3) + item.Label + strings.Repeat("&nbsp;", 8)
			}
			field.Html = htm
		} else if field.FieldType == 13 && len(datas) >= 1 { // 单选按钮
			if field.Default != "" {
				if field.Default == datas[0].Value {
					field.Html = strings.Replace(field.Html, "{{default}}", "checked", 1)
				} else {
					field.Html = strings.Replace(field.Html, "{{default}}", "", 1)
				}
			}
			options = append(options, "onText: '"+datas[0].Value+"'")
			options = append(options, "offText: '"+datas[1].Value+"'")
		}
	}
	if len(options) > 0 {
		attrs = append(attrs, `data-options="`+strings.Join(options, ", ")+`"`)
	}
	field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), -1)
	field.Html = strings.Replace(field.Html, "{{value}}", val, -1)
	return field.Html
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
	fmt.Println("engine", settingData)
	var uploader storage.Uploader
	switch engine {
	case "OSS存储":
		uploader = storage.NewOssUploader(settingData)
	default :
		uploader = storage.NewFileUploader(urlPrefixDir, uploadDir)
	}
	return uploader
}
