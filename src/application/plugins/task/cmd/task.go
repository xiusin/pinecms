package cmd

import (
	"github.com/spf13/cobra"
)

const taskDir = "tasks"

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "任务管理模块",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	TaskCmd.AddCommand(makeTaskCmd)
}
