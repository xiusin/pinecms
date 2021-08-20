package server

import "github.com/spf13/cobra"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "启动后端服务",
	Run: func(cmd *cobra.Command, args []string) {
		serv.Manage([]string{"run"}, cmd.UsageString())
	},
}
