package helper

import (
	"regexp"
	"strconv"
	"strings"
)

type EasyuiGridfields map[string]map[string]string
type EasyuiOptions map[string]string

/**
@property id 表ID
@property url 远程数据地址
@property tableoptions 数据表配置
@property field 数据列配置
@oroperty other 其他配置 暂时预留
@demo Datagrid("aaa", "http://www.baidu.com/", map[string]string{
	"field": "ck", "width": "45", "checkbox": "true", "formatter": "contentNewsListOperateFormatter",
     },map[string]map[string]string{
	"选中": {"field": "ck", "width": "45", "checkbox": "true", "formatter": "contentNewsListOperateFormatter"},
	"标题": {"field": "title", "width": "45", "checkbox": "true", "formatter": "contentNewsListOperateFormatter"},
	"描述": {"field": "desc", "width": "45", "checkbox": "true", "formatter": "contentNewsListOperateFormatter"},
     })
*/

func Datagrid(id, url string, tableoptions EasyuiOptions, field EasyuiGridfields, other ...string) string {
	dataOptions := map[string]interface{}{
		"border":       "false",
		"fit":          "true",
		"fitColumns":   "true",
		"rownumbers":   "true",
		"singleSelect": "true",
		"pagination":   "true",
		"pageList":     [5]int{20, 30, 50, 80, 100},
		"pageSize":     "25",
	}
	style := "width:100%;height:100%;"
	tabopt := []string{}
	if len(tableoptions) != 0 {
		for tabk, tabv := range tableoptions {
			if tabk == "title" {
				tabopt = append(tabopt, tabk+":'"+tabv+"'")
			} else {
				tabopt = append(tabopt, tabk+":"+tabv)
			}
		}
	}
	taboptstr := ""
	if len(tabopt) != 0 {
		taboptstr = "," + strings.Join(tabopt, ",")
	}

	count_field := len(field)
	i := 0
	ths := make([]string, count_field)
	for k, v := range field {
		troptions := []string{}
		currentIndex := 0
		for k1, v1 := range v {
			if k1 == "index" {
				index, err := strconv.Atoi(v1)
				if err != nil {
					return (err.Error())
				}
				currentIndex = index
				if currentIndex < 0 || currentIndex >= count_field {
					return ("index字段值超出范围")
				}
			}
			if k1 == "field" {
				troptions = append(troptions, k1+":'"+v1+"'")
			} else {
				troptions = append(troptions, k1+":"+v1)
			}
		}
		ths[currentIndex] = `<th data-options="` + strings.Join(troptions, ",") + `">` + k + `</th>`
		i++
	}
	th := strings.Join(ths, "\r\n")
	str := `
	<table id="` + id + `" class="easyui-datagrid" style="` + style + `"
           data-options="
           singleSelect:` + dataOptions["singleSelect"].(string) + `,
           url:'` + url + `',
           method:'get',
           border:` + dataOptions["pageSize"].(string) + `,
           fit:` + dataOptions["fit"].(string) + `,
           fitColumns:` + dataOptions["fitColumns"].(string) + `,
           pageSize:` + dataOptions["pageSize"].(string) + `,
           rownumbers:` + dataOptions["rownumbers"].(string) + `,
           autoRowHeight:false,
           singleSelect:` + dataOptions["singleSelect"].(string) + `,
           pagination:` + dataOptions["pagination"].(string) + taboptstr + `
           ">
        <thead>
        <tr>` + th + `</tr>
        </thead>
    </table>
	`
	return str
}

func Propertygrid(id string, options EasyuiOptions) string {

	title, ok1 := options["title"]
	url, ok2 := options["url"]
	toolbar, ok3 := options["toolbar"]
	if !ok1 || !ok2 || !ok3 {
		return "Propertygrid Error : 必须包含元素 : title,url,toolbar"
	}
	str := `<table id="` + id + `" class="easyui-propertygrid"
	data-options="
	border:false,
	fix:true,
	title:'` + title + `',
	url:'` + url + `',
	toolbar:` + toolbar + `,
	showHeader:true,
	showGroup:true,
	columns: [[{ field: 'name', title: '属性名称', sortable:true,width:80},{ field: 'value',title: '属性值',width:200}]],
	scrollbarSize:0" ></table>`
	return str
}

/**
 */
func Treegrid(id string, url string, options EasyuiOptions, field EasyuiGridfields) string {

	dataOptions := map[string]interface{}{
		"border":       "false",
		"fit":          "true",
		"fitColumns":   "true",
		"rownumbers":   "true",
		"singleSelect": "true",
		"animate":      "true",
	}

	tabopt := []string{}
	if len(options) != 0 {
		for tabk, tabv := range options {
			if isNumOrIsBool(tabv) {
				tabopt = append(tabopt, tabk+":"+tabv)
			} else {
				if tabk != "toolbar" {
					tabopt = append(tabopt, tabk+":'"+tabv+"'")
				} else {
					tabopt = append(tabopt, tabk+":"+tabv)
				}
			}
		}
	}
	taboptstr := ""
	if len(tabopt) != 0 {
		taboptstr = ",\r\n" + strings.Join(tabopt, ",\r\n")
	}
	//加入field排序字段index
	count_field := len(field)
	i := 0
	ths := make([]string, count_field)
	for k, v := range field {
		troptions := []string{}
		currentIndex := 0
		for k1, v1 := range v {
			if k1 == "index" {
				index, err := strconv.Atoi(v1)
				if err != nil {
					return (err.Error())
				}
				currentIndex = index
				if currentIndex < 0 || currentIndex >= count_field {
					return ("index字段值超出范围")
				}
			}
			if k1 != "formatter" {
				if isNumOrIsBool(v1) {
					troptions = append(troptions, k1+":"+v1)
				} else {
					troptions = append(troptions, k1+":'"+v1+"'")
				}
			} else {
				troptions = append(troptions, k1+":"+v1)
			}
		}
		ths[currentIndex] = `<th data-options="` + strings.Join(troptions, ",") + `">` + k + `</th>`
		i++
	}
	th := strings.Join(ths, "\r\n")
	style := "width:100%;height:100%;"
	str := `<table  id="` + id + `"
			   class="easyui-treegrid"
			   data-options="
			   	   url: '` + url + `',
			   	   singleSelect:` + dataOptions["singleSelect"].(string) + `,
				   method:'get',
				   fit:` + dataOptions["fit"].(string) + `,
				   fitColumns:` + dataOptions["fitColumns"].(string) + `,
				   rownumbers:` + dataOptions["rownumbers"].(string) + `,
				   singleSelect:` + dataOptions["singleSelect"].(string) + `,
				   animate:` + dataOptions["animate"].(string) + taboptstr + `
			"
			style="` + style + `"><thead><tr>` + th + `</tr></thead></table>`
	return str
}

func isNumOrIsBool(str string) bool {
	if str == "true" || str == "false" {
		return true
	}
	bol, _ := regexp.MatchString(`^\d+$`, str)
	return bol
}
