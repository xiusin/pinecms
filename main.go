package main

import (
	. "iriscms/config"
	"github.com/landoop/tableprinter"
	"os"
	"github.com/kataras/tablewriter"
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
		{Name: "Framework", Value: "Iriscms"},
		{Name: "Author", Value: "Lazy007"},
		{Name: "Github", Value: "https://github.com/lazy007/iriscms"},
		{Name: "QQ", Value: "826466266@qq.com"},
	})
	StartApplication()
}
