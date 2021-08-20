package server

import "github.com/spf13/cobra"

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看守护进程状态",
	Run: func(cmd *cobra.Command, args []string) {
		serv.Manage([]string{"status"}, cmd.UsageString())
	},
}
