package cmd

import (
	"github.com/spf13/cobra"
)

var makeTaskCmd = &cobra.Command{
	Use:   "make",
	Short: "创建任务源文件",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
