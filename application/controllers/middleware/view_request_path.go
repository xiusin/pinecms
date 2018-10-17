package middleware

import (
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const logPath = "runtime/logs/"

func init() {
	f, err := os.Stat(logPath)
	if (err != nil && !os.IsExist(err)) || f.IsDir() {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func ViewRequestPath(app *iris.Application) func(ctx context.Context) {
	return func(ctx context.Context) {
		app.Logger().SetOutput(newLogFile())
		ctx.Next()
	}
}

func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return logPath + today + ".log"
}

func newLogFile() *os.File {
	filename := todayFilename()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
