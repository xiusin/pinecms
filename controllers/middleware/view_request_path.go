package middleware

import (
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func ViewRequestPath(app *iris.Application) func(ctx context.Context) {
	return func(ctx context.Context) {
		app.Logger().SetOutput(newLogFile())
		ctx.Next()
	}
}

func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return "logs/" + today + ".log"
}

func newLogFile() *os.File {
	filename := todayFilename()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
