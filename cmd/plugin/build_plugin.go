package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/reload/util"

	"github.com/spf13/cobra"
)

var buildPluginCmd = &cobra.Command{
	Use:   "build",
	Short: "构建插件",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if len(name) == 0 {
			_ = cmd.Usage()
			return
		}
		buildPluginDir := filepath.Join(outputPluginDir, name)
		_ = os.MkdirAll(buildPluginDir, os.ModePerm)

		scriptName := filepath.Join(sourcePluginDir, name, name+".go")

		_, err := os.Stat(scriptName)
		helper.PanicErr(err)

		if conf, err := os.ReadFile(filepath.Join(sourcePluginDir, name, configName)); err != nil {
			helper.PanicErr(err)
		} else {
			helper.PanicErr(os.WriteFile(filepath.Join(buildPluginDir, configName), conf, os.ModePerm))
		}

		outPluginName := filepath.Join(buildPluginDir, name+".so")
		buildCmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outPluginName, scriptName)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stdout
		buildCmd.Env = os.Environ()
		buildCmd.Dir = util.AppPath()

		helper.PanicErr(buildCmd.Run())

		fmt.Println("构建插件", outPluginName, "成功")
	},
}

func init() {
	buildPluginCmd.Flags().String("name", "", "传入需要构建的插件文件名称")
}
