package server

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "守护进程方式打开服务",
	Run: func(cmd *cobra.Command, args []string) {
		serv.Manage([]string{"start"}, cmd.UsageString())
	},
}
