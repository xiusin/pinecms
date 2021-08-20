package server

import "github.com/spf13/cobra"

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装服务",
	Run: func(cmd *cobra.Command, args []string) {
		serv.Manage([]string{"install"}, cmd.UsageString())
	},
}
