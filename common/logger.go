package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)
var (
	Logger = InitLogger()
)

func InitLogger() *zap.SugaredLogger {
	//tempLogger, err := zap.NewProduction()
	//cfg := zap.Config {
	//	//Encoding:         "json",
	//	Encoding:         "console",
	//	Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
	//	OutputPaths:      []string{"stderr"},
	//	ErrorOutputPaths: []string{"stderr"},
	//	EncoderConfig: zapcore.EncoderConfig {
	//		MessageKey: "msg",
	//
	//		LevelKey:    "severity",
	//		EncodeLevel: zapcore.CapitalLevelEncoder,
	//
	//		TimeKey:    "time",
	//		//EncodeTime: zapcore.ISO8601TimeEncoder,
	//		EncodeTime: SyslogTimeEncoder,
	//
	//		CallerKey:    "ecode",
	//		EncodeCaller: zapcore.ShortCallerEncoder,
	//		//EncodeCaller: common.CustomCallerEncoder,
	//	},
	//}

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.LevelKey = "severity"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.CallerKey = "ecode"




	tempLogger, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer tempLogger.Sync()

	return tempLogger.Sugar()
	//return wrap(tempLogger)
	//return tempLogger.WithOptions(
	//	zap.WrapCore(
	//		func(zapcore.Core) zapcore.Core {
	//			return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
	//		}))
	//return tempLogger.Sugar()
}

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func wrap(log *zap.Logger) *Mlogger {
	//core := log.clone()
	//core.callerSkip += 2
	return &Mlogger{log}
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}