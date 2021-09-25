package casbin

import (
	zapx "rin-echo/common/zap"
	"sync/atomic"

	casbinlogger "github.com/casbin/casbin/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger  *zap.Logger
	enabled int32
}

func NewLogger(logger *zap.Logger, enabled int32) casbinlogger.Logger {
	return &Logger{
		logger:  logger,
		enabled: enabled,
	}
}

func (l *Logger) EnableLog(enable bool) {
	var enab int32
	if enable {
		enab = 1
	}
	atomic.StoreInt32(&l.enabled, enab)
}

// IsEnabled returns if logger is enabled.
func (l *Logger) IsEnabled() bool {
	return atomic.LoadInt32(&l.enabled) == 1
}

// LogModel log info related to model.
func (l *Logger) LogModel(model [][]string) {
	if !l.IsEnabled() {
		return
	}

	l.logger.Info("LogModel", zapx.StringMatrix("model", model))
}

// LogEnforce log info related to enforce.
func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	if !l.IsEnabled() {
		return
	}

	l.logger.Info(
		"LogEnforce",
		zap.String("matcher", matcher),
		zap.Array("request", zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
			for _, v := range request {
				if err := ae.AppendReflected(v); err != nil {
					return err
				}
			}
			return nil
		})),
		zap.Bool("result", result),
		zapx.StringMatrix("explains", explains),
	)
}

// LogRole log info related to role.
func (l *Logger) LogRole(roles []string) {
	if !l.IsEnabled() {
		return
	}

	l.logger.Info("LogRole", zap.Strings("roles", roles))
}

// LogPolicy log info related to policy.
func (l *Logger) LogPolicy(policy map[string][][]string) {
	if !l.IsEnabled() {
		return
	}

	l.logger.Info("LogPolicy", zapx.MapStringMatrix("policy", policy))
}
