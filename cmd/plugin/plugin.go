package plugin

import (
	"github.com/spf13/cobra"
)

var PluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "插件模块管理",
}

func init() {
	PluginCmd.AddCommand(makePluginCmd, buildPluginCmd)
}
