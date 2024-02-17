package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiusin/pinecms/cmd/crud"
	"github.com/xiusin/pinecms/cmd/dede"
	"github.com/xiusin/pinecms/cmd/plugin"
	servCmd "github.com/xiusin/pinecms/cmd/server"
	"github.com/xiusin/pinecms/cmd/version"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"github.com/xiusin/pinecms/src/server"
)

var rootCmd = &cobra.Command{
	Use:               "pinecms",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true, DisableNoDescFlag: true},
	Long: `       _                             
      (_)                            
 ____  _ ____   ____ ____ ____   ___ 
|  _ \| |  _ \ / _  ) ___)    \ /___)
| | | | | | | ( (/ ( (___| | | |___ |
| ||_/|_|_| |_|\____)____)_|_|_(___/ 
|_|     		      version: ` + version.Version,

	Run: func(cmd *cobra.Command, args []string) {
		config.InitDB()
		server.Server()
	},
}

func Execute() {
	helper.PanicErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(plugin.Cmd)
	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(servCmd.ServeCmd)
	rootCmd.AddCommand(crud.Cmd)
	rootCmd.AddCommand(menuCmd)
	rootCmd.AddCommand(dede.Cmd)

	server.InitApp()
}
