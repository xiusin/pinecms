package server

import (
	"context"
	"os"
	"os/exec"

	"github.com/xiusin/pinecms/src/config"

	"github.com/spf13/cobra"
	"github.com/xiusin/pinecms/src/server"
	"github.com/xiusin/reload"
	"github.com/xiusin/reload/util"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "开发更新服务",
	Run: func(cmd *cobra.Command, args []string) {
		fe, _ := cmd.Flags().GetBool("fe")
		var cbs []func(ctx context.Context)
		if !util.IsChildMode() {
			if fe {
				cbs = append(cbs, func(ctx context.Context) {
					apidocCmd := exec.Command("yarn", "serve")
					apidocCmd.Dir = "apidoc-ui"
					apidocCmd.Stderr = os.Stdout
					apidocCmd.Stdout = os.Stdout
					apidocCmd.Run()
				}, func(ctx context.Context) {
					adminCmd := exec.Command("yarn", "dev")
					adminCmd.Dir = "admin"
					adminCmd.Stdout = os.Stdout
					adminCmd.Stderr = os.Stdout
					adminCmd.Run()
				})
			}

			// runtime.SetFinalizer(engine, func(engine *Engine) {
			// 	_ = engine.Close()
			// })
		}

		reload.Loop(func() error {
			config.InitDB()
			server.Server()
			return nil
		}, &reload.Conf{Cmd: &reload.CmdConf{
			Params:       []string{"serve", "dev"},
			SubProcessCb: cbs,
		}})
	},
}

func init() {
	devCmd.Flags().Bool("fe", false, "是否启动fe进程")
}
