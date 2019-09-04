package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)
var (
	Logger = InitLogger()
)

func InitLogger() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.LevelKey = "severity"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.CallerKey = "ecode"
	cfg.EncoderConfig.EncodeCaller = MayaCallerEncoder


	tempLogger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer tempLogger.Sync()

	return tempLogger.Sugar()
}
