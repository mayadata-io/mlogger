package glog

import (
	"github.com/golang/glog"
)

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	glog.Info()
}

