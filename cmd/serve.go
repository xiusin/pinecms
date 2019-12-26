package cmd

import (
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动iriscms服务器",
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
