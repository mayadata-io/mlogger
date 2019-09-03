package common

import (
	"go.uber.org/zap"
	"log"
)
var (
	Logger = InitLogger()
)

func InitLogger() *zap.SugaredLogger {
	tempLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer tempLogger.Sync()

	//set formatting
	return tempLogger.Sugar()
}
