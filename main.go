package main
import (
  "github.com/mayadata-io/mlogger/glog"
)

func main() {
  //logger, _ := zap.NewProduction()
  //defer logger.Sync() // flushes buffer, if any
  //sugar := logger.Sugar()
  //sugar.Infow("failed to fetch URL",
  //  // Structured context as loosely typed key-value pairs.
  //  "url", url,
  //  "attempt", 3,
  //  "backoff", time.Second,
  //)
  //sugar.Infof("Failed to fetch URL: %s", url)
  //glog.init()
  glog.Info("hi")
}


