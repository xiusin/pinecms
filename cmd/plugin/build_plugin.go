package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var buildPluginCmd = &cobra.Command{
	Use:   "build",
	Short: "构建插件",
	Run: func(cmd *cobra.Command, args []string) {
		name,_ := cmd.Flags().GetString("name")
		if len(name) == 0 {
			cmd.Usage()
			return
		}
		os.Mkdir(outputPluginDir, os.ModePerm)
		//查找插件文件
		scriptName := filepath.Join(sourcePluginDir, name, name + ".go")
		if _, err := os.Stat(scriptName); err != nil {
			panic(err)
		}
		outPluginName := filepath.Join(outputPluginDir, name + ".so")
		buildCmd := exec.Command("go","build", "-buildmode=plugin", "-o", outPluginName, scriptName)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			panic(err)
		}
		fmt.Println("构建插件", outPluginName, "成功")
	},
}

func init() {
	buildPluginCmd.Flags().String("name","", "传入需要构建的插件文件名称")
}
