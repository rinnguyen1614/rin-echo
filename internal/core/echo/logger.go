package echo

import (
	"fmt"
	"io"
	"os"

	"github.com/rinnguyen1614/rin-echo/internal/core/log"

	"github.com/labstack/echo/v4"
	echologger "github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.SugaredLogger
	level  zap.AtomicLevel
	writer zapcore.WriteSyncer
	prefix string
}

func NewLogger(logger *zap.Logger, prefix string) echo.Logger {
	return NewSugaredLogger(logger.Sugar(), "rin-echo:"+prefix)
}

func NewSugaredLogger(logger *zap.SugaredLogger, prefix string) echo.Logger {
	return &Logger{
		logger: logger,
		level:  zap.NewAtomicLevel(),
		writer: zapcore.Lock(os.Stderr), // check if the logger need a writer.
		prefix: prefix,
	}
}

func Log(logger echo.Logger, level log.Level, i interface{}) {
	if l, ok := logger.(*Logger); ok {
		suggaredLog(l, level, i)
		return
	}

	switch level {
	case log.DebugLevel:
		logger.Debug(i)
	case log.InfoLevel:
		logger.Info(i)
	case log.WarnLevel:
		logger.Warn(i)
	case log.ErrorLevel:
		logger.Error(i)
	case log.FatalLevel:
		logger.Fatal(i)
	case log.PanicLevel:
		logger.Panic(i)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}
}

func Logf(logger echo.Logger, level log.Level, format string, i ...interface{}) {
	if l, ok := logger.(*Logger); ok {
		suggaredLogf(l, level, format, i)
		return
	}

	switch level {
	case log.DebugLevel:
		logger.Debugf(format, i)
	case log.InfoLevel:
		logger.Infof(format, i)
	case log.WarnLevel:
		logger.Warnf(format, i)
	case log.ErrorLevel:
		logger.Errorf(format, i)
	case log.FatalLevel:
		logger.Fatalf(format, i)
	case log.PanicLevel:
		logger.Panicf(format, i)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}
}

func Logj(logger echo.Logger, level log.Level, j echologger.JSON) {
	if l, ok := logger.(*Logger); ok {
		suggaredLogj(l, level, j)
		return
	}

	switch level {
	case log.DebugLevel:
		logger.Debugj(j)
	case log.InfoLevel:
		logger.Infoj(j)
	case log.WarnLevel:
		logger.Warnj(j)
	case log.ErrorLevel:
		logger.Errorj(j)
	case log.FatalLevel:
		logger.Fatalj(j)
	case log.PanicLevel:
		logger.Panicj(j)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}
}

func (l *Logger) Output() io.Writer {
	return l.writer
}

func (l *Logger) SetOutput(w io.Writer) {
	zapcore.AddSync(w)
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

func (l *Logger) Level() echologger.Lvl {
	return echologger.Lvl(l.level.Level())
}

func (l *Logger) SetLevel(v echologger.Lvl) {
	l.level.SetLevel(zapcore.Level(toLogLevel(uint32(v))))
}

func (l *Logger) SetHeader(h string) {

}

func (l *Logger) Print(i ...interface{}) {
	suggaredLog(l, log.InfoLevel, i)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	suggaredLogf(l, log.InfoLevel, format, args...)
}

func (l *Logger) Printj(j echologger.JSON) {
	suggaredLogj(l, log.InfoLevel, j)
}

func (l *Logger) Debug(i ...interface{}) {
	suggaredLog(l, log.DebugLevel, i)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	suggaredLogf(l, log.DebugLevel, format, args...)
}

func (l *Logger) Debugj(j echologger.JSON) {
	suggaredLogj(l, log.DebugLevel, j)
}

func (l *Logger) Info(i ...interface{}) {
	suggaredLog(l, log.InfoLevel, i)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	suggaredLogf(l, log.InfoLevel, format, args...)
}

func (l *Logger) Infoj(j echologger.JSON) {
	suggaredLogj(l, log.InfoLevel, j)
}

func (l *Logger) Warn(i ...interface{}) {
	suggaredLog(l, log.WarnLevel, i)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	suggaredLogf(l, log.WarnLevel, format, args...)
}

func (l *Logger) Warnj(j echologger.JSON) {
	suggaredLogj(l, log.WarnLevel, j)
}

func (l *Logger) Error(i ...interface{}) {
	suggaredLog(l, log.ErrorLevel, i)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	suggaredLogf(l, log.ErrorLevel, format, args...)
}

func (l *Logger) Errorj(j echologger.JSON) {
	suggaredLogj(l, log.ErrorLevel, j)
}

func (l *Logger) Fatal(i ...interface{}) {
	suggaredLog(l, log.FatalLevel, i)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	suggaredLogf(l, log.FatalLevel, format, args...)
}

func (l *Logger) Fatalj(j echologger.JSON) {
	suggaredLogj(l, log.FatalLevel, j)
}

func (l *Logger) Panic(i ...interface{}) {
	suggaredLog(l, log.PanicLevel, i)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	suggaredLogf(l, log.PanicLevel, format, args...)
}

func (l *Logger) Panicj(j echologger.JSON) {
	suggaredLogj(l, log.PanicLevel, j)
}

func suggaredLog(logger *Logger, level log.Level, i interface{}) {
	log.SuggaredLog(logger.logger, level, i)
}

func suggaredLogf(logger *Logger, level log.Level, format string, args ...interface{}) {
	log.SuggaredLogf(logger.logger, level, format, args...)
}

func suggaredLogj(logger *Logger, level log.Level, j echologger.JSON) {
	log.SuggaredLogw(logger.logger, level, "", jsonToKeyAndValues(j)...)
}

func toLogLevel(l uint32) log.Level {
	var lv log.Level
	switch l {
	case 7:
		lv = log.FatalLevel
	case 6:
		lv = log.PanicLevel
	case 5:
		lv = log.DPanicLevel
	case uint32(echologger.ERROR):
		lv = log.ErrorLevel
	case uint32(echologger.WARN):
		lv = log.WarnLevel
	case uint32(echologger.INFO):
		lv = log.InfoLevel
	default:
		lv = log.DebugLevel
	}
	return lv
}

func jsonToKeyAndValues(j echologger.JSON) []interface{} {
	var kvs []interface{}
	for k, v := range j {
		kvs = append(kvs, k, v)
	}
	return kvs
}
