package glog

import (

	gglog "github.com/golang/glog"
	"github.com/mayadata-io/mlogger/common"
	//"go.uber.org/zap"
)

var (
	logger = common.Logger
)

// Flush flushes all pending log I/O.
func Flush() {
	//gglog.Flush()
	logger.Sync()
}

// CopyStandardLogTo arranges for messages written to the Go "log" package's
// default logs to also appear in the Google logs for the named and lower
// severities.  Subsequent changes to the standard log's default output location
// or format may break this behavior.
//
// Valid names are "INFO", "WARNING", "ERROR", and "FATAL".  If the name is not
// recognized, CopyStandardLogTo panics.
func CopyStandardLogTo(name string) {
	//gglog.CopyStandardLogTo(name)
	//logger.CopyStandardLogTo(name)
	//TODO: not supported in this version of mlogger
	return
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	//gglog.Info(args)
	logger.Info(args)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, args ...interface{}) {
	gglog.InfoDepth(depth, args)
	//logger.InfoDepth(depth, args)
	//TODO: depth not supported in this version of mlogger
	//logger.Info(args)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	//gglog.Infoln(args)
	logger.Info(args)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...interface{}) {
	//gglog.Warning(args)
	logger.Warn(args)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, args ...interface{}) {
	gglog.WarningDepth(depth, args)
	//logger.WarningDepth(depth, args)
	//TODO: depth not supported in this version of mlogger
	//logger.Warn(args)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	//gglog.Warningln(args)
	logger.Warn(args)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...interface{}) {
	//gglog.Warningf(format, args)
	logger.Warnf(format, args)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	//gglog.Error(args)
	logger.Error(args)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, args ...interface{}) {
	gglog.ErrorDepth(depth, args)
	//logger.ErrorDepth(depth, args)
	//TODO: depth not supported in this version of mlogger
	//logger.Error(args)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	//gglog.Errorln(args)
	logger.Error(args)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	//gglog.Errorf(format, args)
	logger.Errorf(format, args)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	//gglog.Fatal(args)
	logger.Fatal(args)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func FatalDepth(depth int, args ...interface{}) {
	gglog.FatalDepth(depth, args)
	//logger.FatalDepth(depth, args)
	//TODO: depth not supported in this version of mlogger
	//logger.Fatal(args)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	//gglog.Fatalln(args)
	logger.Fatal(args)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	//gglog.Fatalf(format, args)
	logger.Fatalf(format, args)
}

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Exit(args ...interface{}) {
	//gglog.Exit(args)
	logger.Fatal(args)
}

// ExitDepth acts as Exit but uses depth to determine which call frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func ExitDepth(depth int, args ...interface{}) {
	gglog.ExitDepth(depth, args)
	//logger.ExitDepth(depth, args)
	//TODO: depth not supported in this version of mlogger
	//logger.Fatal(args)
}

// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func Exitln(args ...interface{}) {
	//gglog.Exitln(args)
	logger.Fatal(args)
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Exitf(format string, args ...interface{}) {
	//gglog.Exitf(format, args)
	logger.Fatalf(format, args)
}

type Level gglog.Level
type Verbose gglog.Verbose
// V reports whether verbosity at the call site is at least the requested level.
// The returned value is a boolean of type Verbose, which implements Info, Infoln
// and Infof. These methods will write to the Info log if called.
// Thus, one may write either
//	if glog.V(2) { glog.Info("log this") }
// or
//	glog.V(2).Info("log this")
// The second form is shorter but the first is cheaper if logging is off because it does
// not evaluate its arguments.
//
// Whether an individual call to V generates a log record depends on the setting of
// the -v and --vmodule flags; both are off by default. If the level in the call to
// V is at least the value of -v, or of -vmodule for the source file containing the
// call, the V call will log.
func V(level Level) Verbose {
	return Verbose(gglog.V(gglog.Level(level)))
}

// Info is equivalent to the global Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Info(args ...interface{}) {
	gglog.Verbose(v).Info(args)
	//logger.Verbose(v).Info(args)
	//TODO: verbose level not supported in this version of mlogger
	//logger.Info(args)
}

// Infoln is equivalent to the global Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infoln(args ...interface{}) {
	gglog.Verbose(v).Infoln(args)
	//logger.Verbose(v).Infoln(args)
	//TODO: verbose level not supported in this version of mlogger
	//logger.Info(args)
}

// Infof is equivalent to the global Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infof(format string, args ...interface{}) {
	gglog.Verbose(v).Infof(format, args)
	//logger.Verbose(v).Infof(format, args)
	//TODO: verbose level not supported in this version of mlogger
	//logger.Infof(format, args)
}