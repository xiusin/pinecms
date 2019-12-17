package main

import (
	"os"
	"runtime"

	"github.com/xiusin/iriscms/src/config"

	"github.com/kataras/iris"
	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
)

type row struct {
	Name  string `header:"Key"`
	Value string `header:"Val"`
}

func main() {
	p := tableprinter.New(os.Stdout)
	p.BorderTop, p.BorderBottom, p.BorderLeft, p.BorderRight = true, true, true, true
	p.CenterSeparator, p.ColumnSeparator, p.RowSeparator = "│", "│", "─"
	p.HeaderBgColor, p.HeaderFgColor = tablewriter.BgBlackColor, tablewriter.FgGreenColor
	p.Print([]row{
		{"Name", "xiusin"},
		{"Version", "Development"},
		{"Author", "xiusin"},
		{"WebSite", "http://www.xiusin.com/"},
		{"IrisVersion", iris.Version},
	})
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.Server()
}
