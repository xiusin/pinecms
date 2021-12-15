package server

import (
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
		if !util.IsChildMode() {
			go func() {
				apidocCmd := exec.Command("yarn", "serve")
				apidocCmd.Dir = "apidoc-ui"
				apidocCmd.Stdout = os.Stdout
				apidocCmd.Stderr = os.Stdout
				apidocCmd.Run()
			}()
		}

		reload.Loop(func() error {
			config.InitDB()
			config.Server()
			return nil
		}, &reload.CmdConf{Template: []string{"serve", "dev"}})
	},
}
