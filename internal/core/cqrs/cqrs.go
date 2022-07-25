package cqrs

import "go.uber.org/zap"

func LogCommandExecution(logger *zap.Logger, commandName string, cmd interface{}, err error) {
	var fields []zap.Field
	fields = append(fields, zap.Any("cmd", cmd))

	if err == nil {
		logger.Info(commandName+" command succeeded", fields...)
	} else {
		fields = append(fields, zap.Any("err", err))
		logger.Error(commandName+" command failed", fields...)
	}
}
