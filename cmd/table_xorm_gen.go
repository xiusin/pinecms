package cmd

import (
	"fmt"
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
	str.WriteString(fmt.Sprintf("type %s struct {\n", camelString(tableName)))

	var tableDsl []map[string]interface{}
	var formDsl []map[string]interface{}
	var filterDsl []map[string]interface{}

	for _, col := range t.Cols {
		str.WriteRune('\t')
		str.WriteString(camelString(col.Name))
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

		// Type
		str.WriteString(col.Type)

		// Bracketed type metadata
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

		labelName, xormCol := col.Name, cols[col.Name]
		if len(xormCol.Comment) > 0 {
			c := strings.Split(xormCol.Comment, ":") // 如果填写了备注, 则拆分取第一个
			labelName = c[0]
			clearComment := strings.ReplaceAll(strings.ReplaceAll(xormCol.Comment, "\r\n", " "), "'", "")
			clearComment = strings.ReplaceAll(clearComment, "\"", "")
			str.WriteString(" comment('" + clearComment + "')")
		}

		str.WriteString("\"")
		// 添加Json和schematag
		str.WriteString(" json:\"" + snakeString(col.Name) + "\"")
		// 添加验证规则选项
		if !col.IsPrimaryKey {
			str.WriteString(" validate:\"required\"")
		}
		// ↑↑↑↑↑↑↑↑ 解析struct


		elFieldType, elProps := getFieldTypeAndProps(col.Name, col.Type)

		// 列表字段
		tableItem := map[string]interface{}{"prop": snakeString(col.Name), "label": labelName}

		// 表单渲染组件
		comp := map[string]interface{}{
			"name":  elFieldType,
			"props": elProps,
		}
		// 表单字段
		item := map[string]interface{}{
			"prop":  snakeString(col.Name),
			"label": labelName,
		}

		// 设置字段默认值
		colM := cols[item["prop"].(string)]
		if !colM.DefaultIsEmpty {
			item["value"] = colM.Default
		}

		// 绑定组件
		item["component"] = comp

		// 过滤大字段或不可确定字段
		if elFieldType != "cl-upload-space" && elFieldType != "cl-editor-quill" && elFieldType != "el-switch" {
			filterType := elFieldType
			switch elFieldType {
			case "el-checkbox", "el-switch", "el-select":
				filterType = "el-select"
			}
			props := map[string]interface{}{}
			filterItem := map[string]interface{}{
				"name":  snakeString(col.Name),
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
			case "tinyint":
				vmap := parseCommentInfo(xormCol.Comment)
				if len(vmap) > 0 {
					var opts []map[string]interface{}
					for k, v := range vmap {
						opts = append(opts, map[string]interface{}{"label": v, "key": k})
					}
					filterItem["type"] = "el-select"
					filterItem["options"] = opts
					tableItem["dict"] = filterItem["options"]

					elProps = map[string]interface{}{}
					elProps["name"] = "el-checkbox"
					elProps["options"] = filterItem["options"]
					comp["props"] = elProps
					item["component"] = elProps
				}
			}
			if strings.HasSuffix(filterItem["type"].(string), "-range") {
				searchFieldDsl += `		"` + filterItem["name"].(string) + `":{Op: "range"},`
			} else if xormCol.SQLType.Name == "set" {
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

		if !col.IsPrimaryKey {
			formDsl = append(formDsl, item)
		}
	}
	str.WriteString("}")

	genFrontendFile(tableName, frontendPath, tableDsl, formDsl, filterDsl)

	return str.String()
}
