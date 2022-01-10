package logger

import (
	"bytes"
	"time"

	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type pineCmsLoggerWriter struct {
	orm       *xorm.Engine
	logCh     chan []byte
	closed    bool
	errorFlag []byte
}

func NewPineCmsLogger(orm *xorm.Engine, len uint) *pineCmsLoggerWriter {
	l := &pineCmsLoggerWriter{orm: orm, logCh: make(chan []byte, len), errorFlag: []byte("[ERRO]")}
	go l.BeginConsume()
	return l
}

func (p *pineCmsLoggerWriter) BeginConsume() {
	for {
		log, isCloser := <-p.logCh
		if !isCloser {
			return
		}
		_, err := p.orm.InsertOne(p.parseLog(log))
		if err != nil {
			pine.Logger().Warning("日志入库失败", err)
		}
	}
}

func (p *pineCmsLoggerWriter) Write(data []byte) (int, error) { // 只记录错误日志
	defer func() {
		if err := recover(); err != nil {
			p.closed = true
		}
	}()
	if !p.closed && bytes.Contains(data, p.errorFlag) {
		p.logCh <- data
	}
	return 0, nil
}

func (p *pineCmsLoggerWriter) parseLog(logData []byte) *tables.Log {
	return &tables.Log{
		Level:   uint8(logger.ErrorLevel),
		Message: *helper.Bytes2String(logData),
		Time:    tables.LocalTime(time.Now()),
	}
}

func (p *pineCmsLoggerWriter) Close() {
	close(p.logCh)
}
