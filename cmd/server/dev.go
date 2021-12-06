package server

import (
	"github.com/spf13/cobra"
	config "github.com/xiusin/pinecms/src/server"
	"github.com/xiusin/reload"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "开发更新服务",
	Run: func(cmd *cobra.Command, args []string) {
		reload.Loop(func() error {
			config.InitDB()
			config.Server()
			return nil
		}, &reload.CmdConf{Template: []string{"serve", "dev"}})
	},
}
