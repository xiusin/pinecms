package plugin

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "plugin",
	Short: "插件模块管理",
}

func init() {
	Cmd.AddCommand(makePluginCmd, buildPluginCmd)
}
