package tasks

import pineLogger "github.com/xiusin/logger"

type logger struct {
	pineLogger.AbstractLogger
}

func (l logger) Info(msg string, keysAndValues ...interface{}) {
	l.AbstractLogger.Print(msg, keysAndValues)
}

func (l logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.Errorf("%s: 错误: %s, 参数: %s", msg, err, keysAndValues)
}
