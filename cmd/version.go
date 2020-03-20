package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "dev 0.1.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: " + Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
