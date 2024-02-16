package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
		if _, err := os.Stat(scriptName); err != nil {
			panic(err)
		}

		configJson := filepath.Join(sourcePluginDir, name, configName)

		if conf, err := os.ReadFile(configJson); err != nil {
			panic(err)
		} else {
			if err := os.WriteFile(filepath.Join(buildPluginDir, configName), conf, os.ModePerm); err != nil {
				panic(err)
			}
		}

		outPluginName := filepath.Join(buildPluginDir, name+".so")
		buildCmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outPluginName, scriptName)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stdout
		buildCmd.Env = os.Environ()
		buildCmd.Dir = util.AppPath()

		if err := buildCmd.Run(); err != nil {
			panic(err)
		}
		fmt.Println("构建插件", outPluginName, "成功")
	},
}

func init() {
	buildPluginCmd.Flags().String("name", "", "传入需要构建的插件文件名称")
}
