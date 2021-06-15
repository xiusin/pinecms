package cmd

import (
	"github.com/spf13/cobra"
)

var mateCmd = &cobra.Command{
	Use:   "migration",
	Short: "数据迁移",
	Run: func(cmd *cobra.Command, args []string) {
		//xormEngine := di.MustGet("*xorm.Engine").(*xorm.Engine)
		//xormEngine.DumpAll()
	},
}

func init() {
	rootCmd.AddCommand(mateCmd)
}
