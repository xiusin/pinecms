package server

import (
	"github.com/spf13/cobra"
	"github.com/takama/daemon"
	"runtime"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动pinecms服务器",
}

func init() {
	ServeCmd.AddCommand(installCmd)
	ServeCmd.AddCommand(removeCmd)
	ServeCmd.AddCommand(runCmd)
	ServeCmd.AddCommand(startCmd)
	ServeCmd.AddCommand(statusCmd)
	ServeCmd.AddCommand(stopCmd)

	daemonKind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		daemonKind = daemon.UserAgent
	}
	srv, err := daemon.New("pinecms", "pinecms 内容管理系统服务", daemonKind)
	if err != nil {
		panic(err)
	}
	serv = &Service{Daemon: srv}
}
