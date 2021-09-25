package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

var (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = Level(zap.DebugLevel)
	// InfoLevel is the default logging priority.
	InfoLevel = Level(zap.InfoLevel)
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = Level(zap.WarnLevel)
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = Level(zap.ErrorLevel)
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = Level(zap.DPanicLevel)
	// PanicLevel logs a message, then panics.
	PanicLevel = Level(zap.PanicLevel)
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = Level(zap.FatalLevel)
)

type HasLogLevel interface {
	Level() Level
	SetLevel(level Level)
}
