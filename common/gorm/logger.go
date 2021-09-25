package gorm

import (
	"context"
	"rin-echo/common/log"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
	gormutils "gorm.io/gorm/utils"
)

type LoggerConfig struct {
	SlowThreshold time.Duration
	LogLevel      log.Level
}

// adapter zap for gorm
type Logger struct {
	logger *zap.SugaredLogger
	LoggerConfig
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

var (
// Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), Config{
// 	SlowThreshold: 200 * time.Millisecond,
// 	LogLevel:      Warn,
// 	Colorful:      true,
// })
)

func NewLogger(logger *zap.Logger, config LoggerConfig) gormlogger.Interface {
	return NewSugaredLogger(logger.Sugar(), config)
}

func NewSugaredLogger(logger *zap.SugaredLogger, config LoggerConfig) gormlogger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	return &Logger{
		logger:       logger,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
		LoggerConfig: config,
	}
}

func (l Logger) SetAsDefault() {
	gormlogger.Default = &l
}

func (l *Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = toLogLevel(level)
	return &newlogger
}

func (l Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= log.InfoLevel {
		l.logger.Debugf(l.infoStr+msg, append([]interface{}{gormutils.FileWithLineNum()}, data...)...)
	}
}

func (l Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= log.WarnLevel {
		l.logger.Warnf(l.warnStr+msg, append([]interface{}{gormutils.FileWithLineNum()}, data...)...)
	}
}

func (l Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= log.ErrorLevel {
		l.logger.Errorf(l.errStr+msg, append([]interface{}{gormutils.FileWithLineNum()}, data...)...)
	}
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > log.DebugLevel {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= log.ErrorLevel:
			sql, rows := fc()
			l.logger.Error(l.traceErrStr, zap.String("src", gormutils.FileWithLineNum()), zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= log.WarnLevel:
			sql, rows := fc()
			l.logger.Warn(l.traceWarnStr, zap.String("src", gormutils.FileWithLineNum()), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		case l.LogLevel >= log.InfoLevel:
			sql, rows := fc()
			l.logger.Debug(l.traceStr, zap.String("src", gormutils.FileWithLineNum()), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		}
	}
}

func toLogLevel(l gormlogger.LogLevel) log.Level {
	var lv log.Level
	switch l {
	case gormlogger.Silent:
		lv = log.DebugLevel
	case gormlogger.Error:
		lv = log.ErrorLevel
	case gormlogger.Warn:
		lv = log.WarnLevel
	case gormlogger.Info:
		lv = log.InfoLevel
	case gormlogger.LogLevel(5):
		lv = log.DPanicLevel
	case gormlogger.LogLevel(6):
		lv = log.PanicLevel
	case gormlogger.LogLevel(7):
		lv = log.FatalLevel
	default:
		lv = log.DebugLevel
	}
	return lv
}
