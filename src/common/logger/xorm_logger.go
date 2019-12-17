package logger

import (
	"fmt"
	"io"
	"log"

	"xorm.io/core"
)

// SimpleLogger is the default implment of core.ILogger
type IrisCmsXormLogger struct {
	DEBUG   *log.Logger
	ERR     *log.Logger
	INFO    *log.Logger
	WARN    *log.Logger
	level   core.LogLevel
	showSQL bool
}

var _ core.ILogger = &IrisCmsXormLogger{}

func NewIrisCmsXormLogger(out io.Writer) *IrisCmsXormLogger {
	prefix, flag, l := "[ORM]", log.LstdFlags, core.LOG_DEBUG
	return &IrisCmsXormLogger{
		DEBUG: log.New(out, fmt.Sprintf("%s [DEBU] ", prefix), flag),
		ERR:   log.New(out, fmt.Sprintf("%s [ERRO] ", prefix), flag),
		INFO:  log.New(out, fmt.Sprintf("%s [INFO] ", prefix), flag),
		WARN:  log.New(out, fmt.Sprintf("%s [WARN] ", prefix), flag),
		level: l,
	}
}

func (s *IrisCmsXormLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.ERR.Output(2, fmt.Sprint(v...))
	}
	return
}

func (s *IrisCmsXormLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.ERR.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *IrisCmsXormLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprint(v...))
	}
	return
}

func (s *IrisCmsXormLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *IrisCmsXormLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.INFO.Output(2, fmt.Sprint(v...))
	}
	return
}

func (s *IrisCmsXormLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.INFO.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *IrisCmsXormLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.WARN.Output(2, fmt.Sprint(v...))
	}
	return
}

// Warnf implement core.ILogger
func (s *IrisCmsXormLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.WARN.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

func (s *IrisCmsXormLogger) Level() core.LogLevel {
	return s.level
}

func (s *IrisCmsXormLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

func (s *IrisCmsXormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

func (s *IrisCmsXormLogger) IsShowSQL() bool {
	return s.showSQL
}
