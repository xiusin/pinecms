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
	p := tableprinter.New(os.Stdout)
	p.BorderTop, p.BorderBottom, p.BorderLeft, p.BorderRight = true, true, true, true
	p.CenterSeparator, p.ColumnSeparator, p.RowSeparator = "│", "│", "─"
	p.HeaderBgColor,p.HeaderFgColor = tablewriter.BgBlackColor,tablewriter.FgGreenColor
	p.Print([]author{
		{Name: "ProjectName", Value: "IrisCms"},
		{Name: "ProjectVersion", Value: "Development"},
		{Name: "Author", Value: "Lazy007"},
		{Name: "Github", Value: "https://github.com/lazy007/iriscms"},
		{Name: "Framework", Value: "IrisGo"},
		{Name: "IrisVersion", Value: iris.Version},
	})
	runtime.GOMAXPROCS(runtime.NumCPU())
	StartApplication()
}
