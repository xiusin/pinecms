package cmd

import (
	"fmt"
	"strings"
)

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

func (t *SQLTable) toXorm(print bool, tableName string) string {

	var str strings.Builder
	str.WriteString(fmt.Sprintf("type %s struct {\n", camelString(tableName)))

	var tableDsl = []map[string]interface{}{}
	var formDsl = []map[string]interface{}{}
	var filterDsl = []map[string]interface{}{}

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
			str.WriteString(" default '" + col.Default + "'")
		}
		if col.IsPrimaryKey {
			str.WriteString(" pk")
		}
		if col.IsUnique {
			str.WriteString(" unique")
		}
		str.WriteString(" '" + col.Name + "'")

		str.WriteString("\"")
		// 添加Json和schematag
		str.WriteString(" json:\"" + snakeString(col.Name) + "\"")
		str.WriteString(" schema:\"" + snakeString(col.Name) + "\"")
		// 添加验证规则选项
		if !col.IsPrimaryKey {
			str.WriteString(" validate:\"required\"")
		}
		amisType := getFieldType(col.Name, col.Type)

		labelName, xormCol := col.Name, cols[col.Name]
		if xormCol.Comment != "" {
			c := strings.Split(xormCol.Comment, ":")
			labelName = c[0]
		}
		tableItem := map[string]interface{}{
			"name":  snakeString(col.Name),
			"label": labelName,
		}

		item := map[string]interface{}{
			"name":  snakeString(col.Name),
			"label": labelName,
			"type":  amisType,
		}
		getFieldFormExtra(amisType, item)
		if !col.IsPrimaryKey {
			formDsl = append(formDsl, item)
		}
		tableFieldType := "text"
		// 过滤大字段或不可确定字段
		if amisType != "rich-text" && amisType != "textarea" && amisType != "file" {
			filterAmisType := amisType
			switch amisType {
			case "checkboxes", "radios", "select":
				filterAmisType = "select"
			}
			if amisType == "image" {
				tableFieldType = "images"
				tableItem["enlargeAble"] = true
			} else {
				filterItem := map[string]interface{}{
					"name":  snakeString(col.Name),
					"label": "　　" + labelName + ":",
					"type":  filterAmisType,
				}
				if item["options"] != nil {
					filterItem["options"] = item["options"]
					filterItem["clearable"] = true
				}
				switch col.Type {
				case "datetime", "timestamp", "date":
					filterAmisType += "-range"
					if col.Type == "date" {
						filterItem["format"] = "YYYY-MM-DD"
					} else {
						filterItem["format"] = "YYYY-MM-DD HH:mm:ss"
					}
					filterItem["type"] = filterAmisType
				case "tinyint": // 如果仅仅为tinyint但是备注有可以解析的配置信息
					vmap := parseCommentInfo(xormCol.Comment)
					if len(vmap) > 0 {
						opts := []map[string]interface{}{}
						for k, v := range vmap {
							opts = append(opts, map[string]interface{}{"label": v, "value": k})
						}
						filterItem["type"] = "select"
						filterItem["options"] = opts
						FormatEnum(col.Name, opts, tableItem)
					}
				}
				filterDsl = append(filterDsl, filterItem)
			}
			tableItem["type"] = tableFieldType

			var opts interface{}
			var ok bool
			opts, ok = item["options"] // 查看是否存在options
			if !ok {
				opts = []map[string]interface{}{}
			}
			if col.Type == "enum" {
				FormatEnum(col.Name, opts.([]map[string]interface{}), tableItem)
			} else if col.Type == "set" {
				FormatSet(col.Name, opts.([]map[string]interface{}), tableItem)
			}
			tableDsl = append(tableDsl, tableItem)
		}

		str.WriteString("`\n")
	}
	str.WriteString("}")

	genFrontendFile(print, tableName, tableDsl, formDsl, filterDsl)

	return str.String()
}
