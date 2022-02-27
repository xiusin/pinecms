
package server

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/xiusin/pinecms/src/config"

	"github.com/spf13/cobra"
	"github.com/xiusin/pinecms/src/server"
	"github.com/xiusin/reload"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "开发更新服务",
	Run: func(cmd *cobra.Command, args []string) {
		var cbs []func(ctx context.Context)

		fmt.Println(reload.Loop(func() error {
			config.InitDB()

			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout
			_ = cmd.Run()

			server.Server()
			return nil
		}, &reload.Conf{Cmd: &reload.CmdConf{
			Params:       []string{"serve", "dev"},
			SubProcessCb: cbs,
		}}))
	},
}
