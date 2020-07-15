package backend

import (
	"encoding/json"
	"fmt"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/common/storage"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

func clearMenuCache(cache cache.AbstractCache, xorm *xorm.Engine) {
	var roles []*tables.AdminRole
	var menus []*tables.Menu
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
func buildModelForm(mid int64, data map[string]string) string {
	if data == nil {
		data = map[string]string{}
	}
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

	// 判断是否显示
	var showInPage = map[string]controllers.FieldShowInPageList{}
	_ = json.Unmarshal([]byte(documentModel.FieldShowInList), &showInPage)

	h := "<form method='POST' id='content_page_form' enctype='multipart/form-data'>" +
		"<input type='hidden' name='table_name' value='" + documentModel.Table + "'>" +
		"<input type='hidden' name='mid' value='" + strconv.Itoa(int(mid)) + "'>" +
		idInput +
		"<table cellpadding='2' class='dialogtable' style='width: 100%;'>"
	for _, field := range fields {
		if v, ok := showInPage[field.TableField]; ok && v.FormShow {
			h += `<tr><td style="width: 100px;text-align:right;">` + field.FormName + `：</td><td>`
			val, _ := data[field.TableField] // 读取字段值
			if strings.Contains(field.Html, "easyui-") {
				h += easyUIComponents(&field, val)
			} else {
				h += domOrCustomTagComponents(&field, val)
			}
			h += "</td></tr>"
		}
	}
	h += `<tr><td colspan=2><a href="javascript:void(0);" onclick="submitForm()" class="easyui-linkbutton">` + buttonTxt + `</a></td></tr></table></form>`
	return h
}
func domOrCustomTagComponents(field *tables.DocumentModelDsl, val string) string {
	attrs := []string{"name='" + field.TableField + "'"}
	isEditor := strings.HasPrefix(field.Html, "<editor")
	isImageUpload := strings.HasPrefix(field.Html, "<images")
	isMulImageUpload := strings.HasPrefix(field.Html, "<mul-images")
	isTags := strings.HasPrefix(field.Html, "<tags")
	isAttr := strings.HasPrefix(field.Html, "<attr") //文档属性
	isFileUpload := strings.HasPrefix(field.Html, "<fileupload")
	if isEditor {
		field.Html = getEditor(field.TableField, val, strings.Join(attrs, " "), field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isImageUpload {
		field.Html = helper.SiginUpload(field.TableField, val, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isMulImageUpload {
		field.Html = helper.MultiUpload(field.TableField, strings.Split(val, ","), 5, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isTags {
		field.Html = helper.Tags(field.TableField, val, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isFileUpload {
		field.Html = helper.FileUpload(field.TableField, val, field.Required == 1, field.FormName, field.Default, field.RequiredTips)
	} else if isAttr {
		attrs := strings.Split(val, ",")
		kvpair := map[string]string{}
		var labels []string
		json.Unmarshal([]byte(field.Datasource), &kvpair)
		for k,v := range kvpair {
			checked := ""
			for _, val := range attrs {
				if val == k {
					checked = "checked"
					break
				}
			}
			labels = append(labels, `<label><input type="checkbox" name="attrs" value="`+k+`" ` + checked + `>`+v+`[`+k+`]</label>`)
		}
		field.Html =  strings.Join(labels, "&nbsp;&nbsp;&nbsp;")
	} else {
		field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), 1)
		field.Html = strings.Replace(field.Html, "{{value}}", val, 1)
	}
	return field.Html
}

func easyUIComponents(field *tables.DocumentModelDsl, val string) string {
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
	if field.TableField == "visit_count" && val == "" {
		rand.Seed(time.Now().Unix())
		val = strconv.Itoa(rand.Intn(5000))
	}

	if field.TableField == "pubtime" && val == "" {
		val = time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)
	}

	if field.TableField == "listorder" && val == "" {
		val = "30"
	}

	if field.Validator != "" {
		options = append(options, "validType:'"+field.Validator + "'")
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
		if field.FieldType == 5 || field.FieldType == 6 { // 下拉 联动
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
		} else if field.FieldType == 7 || field.FieldType == 8 { // 单选 多选
			options = append(options, "valueField:'value'")
			options = append(options, "textField:'label'")
			defaultVals := strings.Split(field.Default, "|")
			var htm = ""
			var htmlTmp string
			for _, item := range datas {
				if field.FieldType == 8 {
					for _, v := range defaultVals {
						htmlTmp = field.Html
						if val == item.Value || (val == "" && v == item.Value) {
							htmlTmp = strings.Replace(htmlTmp, "{{default}}", "checked", 1)
							break
						} else {
							htmlTmp = strings.Replace(htmlTmp, "{{default}}", "", 1)
						}
					}
				} else {
					htmlTmp = field.Html
					if val == item.Value || (val == "" && item.Value == field.Default) {
						htmlTmp = strings.Replace(htmlTmp, "{{default}}", "checked", 1)
					} else {
						htmlTmp = strings.Replace(htmlTmp, "{{default}}", "", 1)
					}
				}
				htm += strings.Replace(htmlTmp, "{{value}}", item.Value, -1) + " " + item.Label + strings.Repeat("&nbsp;", 5)
			}
			field.Html = htm
		}
	}

	if field.FieldType == 13 { // 单选按钮
		if val == "1" || (val == "" && field.Default == "1")   {
			field.Html = strings.Replace(field.Html, "{{default}}", "checked", 1)
		} else {
			field.Html = strings.Replace(field.Html, "{{default}}", "", 1)
		}
	}

	if len(options) > 0 {
		attrs = append(attrs, `data-options="`+strings.Join(options, ", ")+`"`)
	}
	field.Html = strings.Replace(field.Html, "{{attr}}", strings.Join(attrs, " "), -1)
	if val == "" {
		val = field.Default
	}
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
	var uploader storage.Uploader
	switch engine {
	case "OSS存储":
		uploader = storage.NewOssUploader(settingData)
	default:
		uploader = storage.NewFileUploader(urlPrefixDir, uploadDir)
	}
	return uploader
}
