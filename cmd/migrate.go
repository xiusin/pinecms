package cmd

import (
	"github.com/spf13/cobra"
)


var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "迁移数据",
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
