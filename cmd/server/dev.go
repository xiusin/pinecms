package server

import (
	"context"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	config "github.com/xiusin/pinecms/src/server"
	"github.com/xiusin/reload"
	"github.com/xiusin/reload/util"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "开发更新服务",
	Run: func(cmd *cobra.Command, args []string) {
		var cbs []func(ctx context.Context)
		if !util.IsChildMode() {
			cbs = append(cbs, func(ctx context.Context) {
				apidocCmd := exec.Command("yarn", "serve")
				apidocCmd.Dir = "apidoc-ui"
				apidocCmd.Stderr = os.Stdout
				apidocCmd.Run()
			}, func(ctx context.Context) {
				adminCmd := exec.Command("yarn", "serve", "dev")
				adminCmd.Dir = "admin"
				adminCmd.Stderr = os.Stdout
				adminCmd.Run()
			})
		}

		reload.Loop(func() error {
			config.InitDB()
			config.Server()
			return nil
		}, &reload.Conf{Cmd: &reload.CmdConf{
			Params:       []string{"serve", "dev"},
			SubProcessCb: cbs,
		}})
	},
}
