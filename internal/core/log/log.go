package log

import (
	"fmt"

	"go.uber.org/zap"
)

func Log(logger *zap.Logger, level Level, message string, fields ...zap.Field) {
	switch level {
	case DebugLevel:
		logger.Debug(message, fields...)
	case InfoLevel:
		logger.Info(message, fields...)
	case WarnLevel:
		logger.Warn(message, fields...)
	case ErrorLevel:
		logger.Error(message, fields...)
	case DPanicLevel:
		logger.DPanic(message, fields...)
	case PanicLevel:
		logger.Panic(message, fields...)
	case FatalLevel:
		logger.Fatal(message, fields...)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}
}

func SuggaredLog(logger *zap.SugaredLogger, level Level, args ...interface{}) {
	switch level {
	case DebugLevel:
		logger.Debug(args...)
	case InfoLevel:
		logger.Info(args...)
	case WarnLevel:
		logger.Warn(args...)
	case ErrorLevel:
		logger.Error(args...)
	case DPanicLevel:
		logger.DPanic(args...)
	case PanicLevel:
		logger.Panic(args...)
	case FatalLevel:
		logger.Fatal(args...)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}

}

func SuggaredLogf(logger *zap.SugaredLogger, level Level, template string, args ...interface{}) {
	switch level {
	case DebugLevel:
		logger.Debugf(template, args...)
	case InfoLevel:
		logger.Infof(template, args...)
	case WarnLevel:
		logger.Warnf(template, args...)
	case ErrorLevel:
		logger.Errorf(template, args...)
	case DPanicLevel:
		logger.DPanicf(template, args...)
	case PanicLevel:
		logger.Panicf(template, args...)
	case FatalLevel:
		logger.Fatalf(template, args...)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}
}

func SuggaredLogw(logger *zap.SugaredLogger, level Level, msg string, keysAndValues ...interface{}) {
	switch level {
	case DebugLevel:
		logger.Debugw(msg, keysAndValues...)
	case InfoLevel:
		logger.Infow(msg, keysAndValues...)
	case WarnLevel:
		logger.Warnw(msg, keysAndValues...)
	case ErrorLevel:
		logger.Errorw(msg, keysAndValues...)
	case DPanicLevel:
		logger.DPanicw(msg, keysAndValues...)
	case PanicLevel:
		logger.Panicw(msg, keysAndValues...)
	case FatalLevel:
		logger.Fatalw(msg, keysAndValues...)
	default:
		panic(fmt.Sprint("Unknown LogLevel value: ", level))
	}
}
