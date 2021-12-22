package dede

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:              "dede",
	Short:            "导入织梦CMS数据",
	Long:             `从织梦CMS源中匹配已知数据`,
	TraverseChildren: true,
}
var SqlFieldTypeMap = map[string]string{
	"varchar": "varchar(100)",
	"int":     "int(10)",
}
var SqlLite3FieldTypeMap = map[string]string{
	"varchar": "TEXT",
	"int":     "INTEGER",
}
var ExtraFields = []map[string]string{
	{
		"COLUMN_NAME":    "catid",
		"EXTRA":          "",
		"COLUMN_TYPE":    "int",
		"IS_NULLABLE":    "NO",
		"COLUMN_COMMENT": "所属栏目ID",
		"COLUMN_DEFAULT": "0",
	},
	{
		"COLUMN_NAME":    "mid",
		"EXTRA":          "",
		"COLUMN_TYPE":    "int",
		"IS_NULLABLE":    "NO",
		"COLUMN_COMMENT": "模型ID",
		"COLUMN_DEFAULT": "0",
	},
	{
		"COLUMN_NAME":    "created_time",
		"EXTRA":          "",
		"COLUMN_TYPE":    "datetime",
		"IS_NULLABLE":    "YES",
		"COLUMN_COMMENT": "",
		"COLUMN_DEFAULT": "",
	},
	{
		"COLUMN_NAME":    "updated_time",
		"EXTRA":          "",
		"COLUMN_TYPE":    "datetime",
		"IS_NULLABLE":    "YES",
		"COLUMN_COMMENT": "",
		"COLUMN_DEFAULT": "",
	},
	{
		"COLUMN_NAME":    "deleted_time",
		"EXTRA":          "",
		"COLUMN_TYPE":    "datetime",
		"IS_NULLABLE":    "YES",
		"COLUMN_COMMENT": "",
		"COLUMN_DEFAULT": "",
	},
}
