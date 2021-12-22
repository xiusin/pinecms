package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "dev 0.2.0"

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: " + Version)
	},
}
