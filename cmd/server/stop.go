package server

import "github.com/spf13/cobra"

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "移除服务",
	Run: func(cmd *cobra.Command, args []string) {
		serv.Manage([]string{"remove"}, cmd.UsageString())
	},
}
