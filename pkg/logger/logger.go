package logger

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sLogger := logger.Sugar()
	zap.ReplaceGlobals(logger)
	return sLogger
}
