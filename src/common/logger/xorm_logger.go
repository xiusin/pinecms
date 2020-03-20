package logger

import (
	"fmt"
	"io"
	"log"

	"xorm.io/core"
)

type XormLogger struct {
	DEBUG   *log.Logger
	ERR     *log.Logger
	INFO    *log.Logger
	WARN    *log.Logger
	level   core.LogLevel
	showSQL bool
}

var _ core.ILogger = &XormLogger{}

func NewIrisCmsXormLogger(out io.Writer, level core.LogLevel) *XormLogger {
	flag, l :=  log.LstdFlags, level
	return &XormLogger{
		DEBUG: log.New(out, "[DEBU]", flag),
		ERR:   log.New(out, "[ERRO]", flag),
		INFO:  log.New(out, "[INFO]", flag),
		WARN:  log.New(out, "[WARN]", flag),
		level: l,
	}
}

func (s *XormLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.ERR.Output(2, fmt.Sprint(v...))
	}
	return
}

func (s *XormLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.ERR.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *XormLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprint(v...))
	}
	return
}

func (s *XormLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *XormLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.INFO.Output(2, fmt.Sprint(v...))
	}
	return
}

func (s *XormLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.INFO.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *XormLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.WARN.Output(2, fmt.Sprint(v...))
	}
	return
}

// Warnf implement core.ILogger
func (s *XormLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.WARN.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *XormLogger) Level() core.LogLevel {
	return s.level
}

func (s *XormLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

func (s *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

func (s *XormLogger) IsShowSQL() bool {
	return s.showSQL
}
