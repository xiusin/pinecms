package cmd

import (
	"github.com/kataras/iris/v12"
	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"github.com/xiusin/iriscms/src/config"
	"os"
	"runtime"
)

type row struct {
	Name  string `header:"Key"`
	Value string `header:"Val"`
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动web服务器",
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := cmd.Flags().GetBool("banner")
		if banner {
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
				{"Version", Version},
				{"GoVersion", runtime.Version()},
			})
		}

		runtime.GOMAXPROCS(runtime.NumCPU())
		config.Server()
	},
}

func init() {
	serveCmd.AddCommand(startCmd)
	startCmd.Flags().Bool("banner", true, "显示或隐藏banner信息, false为隐藏")
}
