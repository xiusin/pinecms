package main //github.com/lazy007/iriscms

import (
	. "iriscms/config"
	"os"
	"runtime"

	"github.com/kataras/iris"
	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
)

type author struct {
	Name  string `header:"Name"`
	Value string `header:"Value"`
}

func main() {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor // set header background color for all headers.
	printer.HeaderFgColor = tablewriter.FgGreenColor // set header foreground color for all headers.
	printer.Print([]author{
		{Name: "ProjectName", Value: "Iriscms"},
		{Name: "ProjectVersion", Value: "Dev"},
		{Name: "Author", Value: "Lazy007"},
		{Name: "Github", Value: "https://github.com/lazy007/iriscms"},
		{Name: "QQ", Value: "826466266@qq.com"},
		{Name: "Framework", Value: "iris"},
		{Name: "IrisVersion", Value: iris.Version},
	})
	runtime.GOMAXPROCS(runtime.NumCPU())
	StartApplication()
}
