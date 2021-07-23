package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/cmd/plugin"
	"github.com/xiusin/pinecms/cmd/server"
	"github.com/xiusin/pinecms/cmd/version"
	config "github.com/xiusin/pinecms/src/server"
)

// http://www.network-science.de/ascii/ Font: stop
var rootCmd = &cobra.Command{
	Use: "pinecms",
	Long: `       _                             
      (_)                            
 ____  _ ____   ____ ____ ____   ___ 
|  _ \| |  _ \ / _  ) ___)    \ /___)
| | | | | | | ( (/ ( (___| | | |___ |
| ||_/|_|_| |_|\____)____)_|_|_(___/ 
|_|     		      version: ` + version.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	di.Set("pinecms.cmd.root", func(builder di.AbstractBuilder) (interface{}, error) {
		return rootCmd, nil
	}, true)

	rootCmd.AddCommand(plugin.PluginCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(server.ServeCmd)

	// new cmd

	config.InitApp()
}
