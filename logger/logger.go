package logger

import (
	domain "github.com/Mulderink1/personal_message_mc-services-domain"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() *zap.Logger {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return log
}

func NewLogger() *domain.Logger {
	return &domain.Logger{
		Logger: initLogger(),
	}
}
