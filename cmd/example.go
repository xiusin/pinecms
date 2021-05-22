package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiusin/logger"
	"os"
	"os/exec"
	"strings"
)


var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "获取案例demo数据库以及资源",
	Run: func(cmd *cobra.Command, args []string)  {
		var input string
		fmt.Print("下载demo会覆盖本地数据资源以及清理本地缓存,确定要执行吗?[Y/n] ")
		fmt.Scanln(&input)
		if strings.ToLower(input) == "n" || input == "" {
			return
		}
		os.RemoveAll("pinecms_demo")
		gitCmd := exec.Command("git", "clone", "https://github.com/xiusin/pinecms_demo.git")
		gitCmd.Stdout = os.Stdout
		gitCmd.Stdin = os.Stdin
		if err := gitCmd.Run(); err != nil {
			logger.Error(err)
			return
		}
		os.Rename("pinecms_demo/data.db.demo", "data.db.demo")
		os.Rename("pinecms_demo/resources/themes/example", "resources/themes/example")
		os.Rename("pinecms_demo/resources/assets/example", "resources/assets/example")
		os.Rename("pinecms_demo/resources/uploads", "resources/uploads")
		os.RemoveAll("pinecms_demo")
		fmt.Println(`1. 请修改配置文件database.yml的数据源为: data.db.demo
2. 请配置application.yml主题为: example`)
		err := os.RemoveAll("runtime/cache.db")
		if err != nil {
			logger.Print("删除缓存文件失败, 请手动删除缓存文件")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
