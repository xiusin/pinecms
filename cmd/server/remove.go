package server

import "github.com/spf13/cobra"

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止守护进程",
	Run: func(cmd *cobra.Command, args []string) {
		serv.Manage([]string{"stop"}, cmd.UsageString())
	},
}
