package cmd

import (
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "导入其他CMS数据",
	Long: `从其他CMS源中匹配已知数据. 支持如下CMS:
1. DEDECMS (usage: pinecms import dede [command])`,
TraverseChildren:true,
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

func init() {
	rootCmd.AddCommand(importCmd)
}
