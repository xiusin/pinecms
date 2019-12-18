package cmd

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"github.com/xiusin/iriscms/src/config"
)

type row struct {
	Name  string `header:"Key"`
	Value string `header:"Val"`
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	Long: `启动iris服务: 实现为子命令的方式
iriscms serve [command]
`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		p := tableprinter.New(os.Stdout)
		p.BorderTop, p.BorderBottom, p.BorderLeft, p.BorderRight = true, true, true, true
		p.CenterSeparator, p.ColumnSeparator, p.RowSeparator = "│", "│", "─"
		p.HeaderBgColor, p.HeaderFgColor = 40, 32
		p.Print([]row{
			{"Name", "xiusin"},
			{"Version", "Development"},
			{"Author", "xiusin"},
			{"WebSite", "http://www.xiusin.com/"},
			{"IrisVersion", iris.Version},
		})
		runtime.GOMAXPROCS(runtime.NumCPU())

		go func() {
			time.Sleep(2 * time.Second)
			if runtime.GOOS == "darwin" {
				exec.Command(`open`, `http://localhost:2019/b/index/index`).Start()
			}
		}()
		config.Server()
	},
}

func init() {
	serveCmd.AddCommand(runCmd)
	rootCmd.AddCommand(serveCmd)
}
