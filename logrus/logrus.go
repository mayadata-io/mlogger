package logrus

import (
	lrs "github.com/Sirupsen/logrus"
	"github.com/mayadata-io/mlogger/common"
	"io"
)

var (
	logger = common.Logger
)

type Logger lrs.Logger

func StandardLogger() *Logger {
	return (*Logger)(lrs.StandardLogger())
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	lrs.SetOutput(out)
}

type Formatter lrs.Formatter
// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter Formatter) {
	lrs.SetFormatter(formatter)
}

type Level lrs.Level
// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	lrs.SetLevel((lrs.Level)(level))
}

// GetLevel returns the standard logger level.
func GetLevel() Level {
	return (Level)(lrs.GetLevel())
}

type Hook lrs.Hook
// AddHook adds a hook to the standard logger hooks.
func AddHook(hook Hook) {
	lrs.AddHook((lrs.Hook)(hook))
}

type Entry lrs.Entry
// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *Entry {
	return (*Entry)(lrs.WithError(err))
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *Entry {
	return (*Entry)(lrs.WithField(key, value))
}

type Fields lrs.Fields
// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *Entry {
	return (*Entry)(lrs.WithFields((lrs.Fields)(fields)))
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	lrs.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	lrs.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logger.Info(args)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	logger.Warn(args)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	logger.Warn(args)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	logger.Error(args)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	lrs.Debugf(format, args)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	lrs.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	lrs.Debugln(args)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	lrs.Println(args)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	logger.Info(args)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	logger.Warn(args)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	logger.Warn(args)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	logger.Error(args)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	logger.Panic(args)
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	logger.Fatal(args)
}

// A constant exposing all logging levels
//var AllLevels = lrs.AllLevels

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = lrs.PanicLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = lrs.FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = lrs.ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = lrs.WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = lrs.InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = lrs.DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	//TraceLevel = lrs.TraceLevel
)