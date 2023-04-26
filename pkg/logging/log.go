package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// NewLogger gives a new zap sugared logger, common function in case we want to modify how we log
func NewLogger(name string) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, _ := config.Build()
	return logger.With(zap.String("logger_name", name)).Sugar()
}
