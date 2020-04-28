package cmd

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动pinecms服务器",
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
