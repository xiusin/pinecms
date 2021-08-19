package crud

import (
	"fmt"
	"github.com/xiusin/pinecms/cmd/util"
	"strconv"
	"strings"
)

var searchFieldDsl string

// SQLTable 表名结构体
type SQLTable struct {
	Name string
	Cols []SQLColumn
}

// SQLColumn 描述结构体
type SQLColumn struct {
	Name, Type    string
	IsPrimaryKey  bool
	IsUnique      bool
	Length        string
	EnumValues    []string
	AutoIncrement bool
	NotNull       bool
	Default       string
}

func (t *SQLTable) toXorm(print bool, tableName string, frontendPath string) string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("type %s struct {\n", util.CamelString(tableName)))

	var tableDsl []map[string]interface{}
	var formDsl []map[string]interface{}
	var filterDsl []map[string]interface{}

	for _, col := range t.Cols {
		tableField := util.SnakeString(col.Name)
		coreCol := cols[tableField]

		// ↓↓↓↓↓↓↓↓ 解析生成table结构 开始
		{
			str.WriteRune('\t')
			str.WriteString(util.CamelString(col.Name))

			var goType string
			switch col.Type {
			case "varchar", "text", "char", "longtext", "mediumtext", "smalltext", "tinytext", "set":
				goType = "string"
			case "enum", "tinyint":
				goType = "int"
			case "int", "bigint":
				goType = "int64"
			case "double", "float", "decimal":
				goType = "float64"
			case "year", "date", "datetime", "time", "timestamp":
				goType = "LocalTime"
			case "blob":
				goType = "[]byte"
			default:
				panic(col.Name + " 是一个未知类型")
			}
			str.WriteString(" " + goType)
			str.WriteString(" `xorm:\"")

			str.WriteString(col.Type)

			if len(col.EnumValues) > 0 {
				str.WriteRune('(')
				for i, en := range col.EnumValues {
					str.WriteString(en)
					if i != len(col.EnumValues)-1 {
						str.WriteRune(',')
					}
				}
				str.WriteRune(')')
			} else if len(col.Length) > 0 {
				str.WriteString("(" + col.Length + ")")
			}

			if col.AutoIncrement {
				str.WriteString(" autoincr")
			}

			if col.NotNull {
				str.WriteString(" not null")
			}

			if len(col.Default) > 0 {
				if strings.Contains(goType, "int") || strings.Contains(goType, "float") || col.Default == "null" {
					str.WriteString(" default " + col.Default)
				} else {
					str.WriteString(" default '" + col.Default + "'")
				}
			}
			if col.IsPrimaryKey {
				str.WriteString(" pk")
			}
			if col.IsUnique {
				str.WriteString(" unique")
			}
			str.WriteString(" '" + col.Name + "'")

			if len(cols[col.Name].Comment) > 0 {
				clearComment := strings.ReplaceAll(strings.ReplaceAll(cols[col.Name].Comment, "\r\n", " "), "'", "")
				clearComment = strings.ReplaceAll(clearComment, "\"", "")
				str.WriteString(" comment('" + clearComment + "')")
			}

			str.WriteString("\"")

			// 设置JsonTag
			str.WriteString(" json:\"" + tableField + "\"")

			// 添加验证规则选项
			if !col.IsPrimaryKey {
				str.WriteString(" validate:\"required\"")
			}
		}
		// ↑↑↑↑↑↑↑↑ 解析生成table结构 结束

		labelName, elFieldType, elProps := getLabelAndFieldTypeAndProps(col, coreCol)

		tableItem := map[string]interface{}{"prop": tableField, "label": labelName} // 列表字段

		comp := map[string]interface{}{"name": elFieldType, "props": elProps}                     // 渲染组件
		item := map[string]interface{}{"prop": tableField, "label": labelName, "component": comp} // upsert 组件

		// 设置字段默认值
		if !coreCol.DefaultIsEmpty {
			if elProps["is_number"].(bool) || elProps["is_float"].(bool) {
				if elProps["is_float"].(bool) {
					item["value"], _ = strconv.ParseFloat(coreCol.Default, 64)
				} else {
					item["value"], _ = strconv.Atoi(coreCol.Default)
				}
			} else {
				item["value"] = coreCol.Default
			}
		}

		// 过滤大字段或不可确定字段
		if elFieldType != "cl-upload-space" && elFieldType != "cl-editor-quill" && elFieldType != "el-switch" {
			filterType := elFieldType // 头部筛选字段
			switch elFieldType {
			case "el-checkbox", "el-switch", "el-select", "cms-checkbox", "cms-radio":
				filterType = "el-select"
			}
			props := map[string]interface{}{}
			filterItem := map[string]interface{}{
				"name":  util.SnakeString(col.Name),
				"label": labelName,
				"type":  filterType,
			}
			if item["options"] != nil {
				filterItem["options"] = item["options"]
				tableItem["dict"] = item["options"]
			}
			switch col.Type {
			case "datetime":
				props["type"] = "datetimerange"
			case "date":
				props["type"] = "date"
			default:
				if elFieldType == "tinyint" && elProps["options"] != nil && comp["name"] == "el-input-number" {
					filterItem["type"] = "el-select"
					filterItem["options"] = elProps["options"]
					tableItem["dict"] = filterItem["options"]
					comp["name"] = "cms-radio" // 单选
					elProps = map[string]interface{}{}
					elProps["options"] = filterItem["options"]
					elProps["size"] = "mini"
					comp["props"] = elProps
				}

				if (elFieldType == "varchar" || elFieldType == "char") && elProps["options"] != nil {
					filterItem["type"] = "el-select"
					filterItem["options"] = elProps["options"]
					tableItem["dict"] = filterItem["options"]
					comp["name"] = "cms-select" // 多选
					elProps = map[string]interface{}{}
					elProps["options"] = filterItem["options"]
					elProps["size"] = "mini"
					comp["props"] = elProps
				}
			}
			if strings.HasSuffix(filterItem["type"].(string), "-range") {
				searchFieldDsl += `		"` + filterItem["name"].(string) + `":{Op: "range"},`
			} else if coreCol.SQLType.Name == "set" {
				searchFieldDsl += `		"` + filterItem["name"].(string) + `":{Op: "set"},` // findInSet
			} else {
				searchFieldDsl += `		"` + filterItem["name"].(string) + `":{Op: "="},`
			}
			searchFieldDsl += "\r\n"
			filterDsl = append(filterDsl, filterItem)
			//tableItem["type"] = tableFieldType
			tableDsl = append(tableDsl, tableItem)
		}

		str.WriteString("`\n")
		delete(elProps, "is_number")
		delete(elProps, "is_float")
		if !col.IsPrimaryKey {
			formDsl = append(formDsl, item)
		}
	}
	str.WriteString("}")

	genFrontendFile(tableName, frontendPath, tableDsl, formDsl, filterDsl)

	return str.String()
}
