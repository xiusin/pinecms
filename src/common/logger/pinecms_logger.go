package logger

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"runtime/debug"
	"sync"
	"time"
)

type PineCmsLogger struct {
	sync.Mutex
	*logger.Logger
	Orm   *xorm.Engine
	Table *tables.Log
}

func (l *PineCmsLogger) ToDB(level logger.Level, ctx *pine.Context, args ...interface{}) {
	if l.Level < level {
		return
	}
	l.Lock()
	defer l.Unlock()
	log := tables.Log{
		Level:   uint8(level),
		Uri:     ctx.Request.URI().String(),
		Method:  string(ctx.Method()),
		Params:  string(ctx.Request.Body()),
		Message: fmt.Sprint(args...),
		Ip:      ctx.ClientIP(),
		Time:    tables.LocalTime(time.Now()),
	}
	if level == logger.ErrorLevel {
		log.Stack = *helper.Bytes2String(debug.Stack())
	}
	l.Orm.InsertOne(log)
}
