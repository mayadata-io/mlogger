package main
import (
  "github.com/mayadata-io/mlogger/common"
  //"github.com/mayadata-io/mlogger/common"
  //"go.uber.org/zap"
  "time"
)

func main() {
  logger := common.Logger
  //logger, _ := zap.NewProduction()
  //defer logger.Sync() // flushes buffer, if any

  poolname := "testPool1"
  logger.Infow("failed to create pool",
    // Structured context as loosely typed key-value pairs.
    "rname", poolname,
    "attempt", 3,
    "backoff", time.Second,
    "ecode", "cstor.pool.create.failed",
  )
  logger.Infof("Failed to create pool: %s", poolname)
}


