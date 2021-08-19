package server

import (
	"github.com/xiusin/pinecms/cmd/version"
	"os"
	"runtime"

	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"github.com/xiusin/pine"
	config "github.com/xiusin/pinecms/src/server"
)

type row struct {
	Name  string `header:"Key"`
	Value string `header:"Val"`
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动pinecms服务",
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := cmd.Flags().GetBool("banner")
		if banner {
			p := tableprinter.New(os.Stdout)
			p.BorderTop, p.BorderBottom, p.BorderLeft, p.BorderRight = true, true, true, true
			p.CenterSeparator, p.ColumnSeparator, p.RowSeparator = "│", "│", "─"
			p.HeaderBgColor, p.HeaderFgColor = 40, 32
			p.Print([]row{
				{"Name", "PineCMS内容管理系统"},
				{"Status", "Development"},
				{"Author", "xiusin"},
				{"PineVersion", pine.Version},
				{"Version", version.Version},
				{"GoVersion", runtime.Version()},
			})
		}
		config.InitDB()
		config.Server()
	},
}

func init() {
	ServeCmd.AddCommand(startCmd)
	startCmd.Flags().Bool("banner", true, "显示或隐藏banner信息, false为隐藏")
	startCmd.Flags().Bool("daemon", true, "后台进程运行")
}
