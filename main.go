package main

import (
  //"time"
  //"github.com/mayadata-io/mlogger/common"
  "github.com/mayadata-io/mlogger/glog"
)

func main() {

  poolname := "testPool1"
  //logger.Infow("failed to create pool",
  //  "rname", poolname,
  //  "attempt", 3,
  //  "backoff", time.Second,
  //  //"ecode", "cstor.pool.create.failed",
  //)
  glog.Infof("Failed to create pool: %s", poolname)
}


