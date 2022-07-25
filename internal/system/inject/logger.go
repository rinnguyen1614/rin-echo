package inject

import "go.uber.org/zap"

func GetLogger() *zap.Logger {
	if service.logger == nil {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		service.logger = logger
	}
	return service.logger
}
