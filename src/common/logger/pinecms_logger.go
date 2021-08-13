package logger

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"time"
)

type PineCmsLoggerWriter struct {
	Orm *xorm.Engine
}

func (p *PineCmsLoggerWriter) Write(data []byte) (int, error) {
	go p.Orm.InsertOne(&tables.Log{
		//Level:   uint8(level),
		//Uri:     ctx.Request.URI().String(),
		//Method:  string(ctx.Method()),
		//Params:  string(ctx.Request.Body()),
		Message: *helper.Bytes2String(data),
		//Ip:      ctx.ClientIP(),
		Time:    tables.LocalTime(time.Now()),
	})
	//	if level == logger.ErrorLevel {
	//		log.Stack = *helper.Bytes2String(debug.Stack())
	//	}
	return 0, nil
}
