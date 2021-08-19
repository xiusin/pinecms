package cmd

import (
	"github.com/spf13/cobra"
)

var mateCmd = &cobra.Command{
	Use:   "migration",
	Short: "数据迁移",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
