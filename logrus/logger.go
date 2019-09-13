package logrus

import (
	"context"
	"io"
	"os"
	"time"
	lrs "github.com/Sirupsen/logrus"
)

type Logger lrs.Logger

type MutexWrap lrs.MutexWrap

func (mw *MutexWrap) Lock() {
	(*lrs.MutexWrap)(mw).Lock()
}

func (mw *MutexWrap) Unlock() {
	(*lrs.MutexWrap)(mw).Unlock()
}

func (mw *MutexWrap) Disable() {
	(*lrs.MutexWrap)(mw).Disable()
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` and `Hooks` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &Logger{
//      Out: os.Stderr,
//      Formatter: new(JSONFormatter),
//      Hooks: make(LevelHooks),
//      Level: logrus.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	return (*Logger)(&lrs.Logger{
		Out:          os.Stderr,
		Formatter:    (*lrs.TextFormatter)(new(TextFormatter)),
		Hooks:        (lrs.LevelHooks)(make(LevelHooks)),
		Level:        (lrs.Level)(InfoLevel),
		ExitFunc:     os.Exit,
		ReportCaller: false,
	})
}

// Adds a field to the log entry, note that it doesn't log until you call
// Debug, Print, Info, Warn, Error, Fatal or Panic. It only creates a log entry.
// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return (*Entry)((*lrs.Logger)(logger).WithField(key, value))
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	return (*Entry)((*lrs.Logger)(logger).WithFields((lrs.Fields)(fields)))
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *Logger) WithError(err error) *Entry {
	return (*Entry)((*lrs.Logger)(logger).WithError(err))
}

// Add a context to the log entry.
func (logger *Logger) WithContext(ctx context.Context) *Entry {
	return (*Entry)((*lrs.Logger)(logger).WithContext(ctx))
}

// Overrides the time of the log entry.
func (logger *Logger) WithTime(t time.Time) *Entry {
	return (*Entry)((*lrs.Logger)(logger).WithTime(t))
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	(*lrs.Logger)(logger).Logf((lrs.Level)(level), format, args)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Tracef(format, args)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Debugf(format, args)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Infof(format, args)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Printf(format, args)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Warnf(format, args)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Warningf(format, args)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Errorf(format, args)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Fatalf(format, args)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	(*lrs.Logger)(logger).Panicf(format, args)
}

func (logger *Logger) Log(level Level, args ...interface{}) {
	(*lrs.Logger)(logger).Log((lrs.Level)(level), args)
}

func (logger *Logger) Trace(args ...interface{}) {
	(*lrs.Logger)(logger).Trace(args)
}

func (logger *Logger) Debug(args ...interface{}) {
	(*lrs.Logger)(logger).Debug(args)
}

func (logger *Logger) Info(args ...interface{}) {
	(*lrs.Logger)(logger).Info(args)
}

func (logger *Logger) Print(args ...interface{}) {
	(*lrs.Logger)(logger).Print(args)
}

func (logger *Logger) Warn(args ...interface{}) {
	(*lrs.Logger)(logger).Warn(args)
}

func (logger *Logger) Warning(args ...interface{}) {
	(*lrs.Logger)(logger).Warning(args)
}

func (logger *Logger) Error(args ...interface{}) {
	(*lrs.Logger)(logger).Error(args)
}

func (logger *Logger) Fatal(args ...interface{}) {
	(*lrs.Logger)(logger).Fatal(args)
}

func (logger *Logger) Panic(args ...interface{}) {
	(*lrs.Logger)(logger).Panic(args)
}

func (logger *Logger) Logln(level Level, args ...interface{}) {
	(*lrs.Logger)(logger).Logln((lrs.Level)(level), args)
}

func (logger *Logger) Traceln(args ...interface{}) {
	(*lrs.Logger)(logger).Traceln(args)
}

func (logger *Logger) Debugln(args ...interface{}) {
	(*lrs.Logger)(logger).Debugln(args)
}

func (logger *Logger) Infoln(args ...interface{}) {
	(*lrs.Logger)(logger).Infoln(args)
}

func (logger *Logger) Println(args ...interface{}) {
	(*lrs.Logger)(logger).Println(args)
}

func (logger *Logger) Warnln(args ...interface{}) {
	(*lrs.Logger)(logger).Warnln(args)
}

func (logger *Logger) Warningln(args ...interface{}) {
	(*lrs.Logger)(logger).Warningln(args)
}

func (logger *Logger) Errorln(args ...interface{}) {
	(*lrs.Logger)(logger).Errorln(args)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	(*lrs.Logger)(logger).Fatalln(args)
}

func (logger *Logger) Panicln(args ...interface{}) {
	(*lrs.Logger)(logger).Panicln(args)
}

func (logger *Logger) Exit(code int) {
	(*lrs.Logger)(logger).Exit(code)
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func (logger *Logger) SetNoLock() {
	(*lrs.Logger)(logger).SetNoLock()
}

// SetLevel sets the logger level.
func (logger *Logger) SetLevel(level Level) {
	(*lrs.Logger)(logger).SetLevel((lrs.Level)(level))
}

// GetLevel returns the logger level.
func (logger *Logger) GetLevel() Level {
	return (Level)((*lrs.Logger)(logger).GetLevel())
}

// AddHook adds a hook to the logger hooks.
func (logger *Logger) AddHook(hook Hook) {
	(*lrs.Logger)(logger).AddHook((lrs.Hook)(hook))
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param
func (logger *Logger) IsLevelEnabled(level Level) bool {
	return (*lrs.Logger)(logger).IsLevelEnabled((lrs.Level)(level))
}

// SetFormatter sets the logger formatter.
func (logger *Logger) SetFormatter(formatter Formatter) {
	(*lrs.Logger)(logger).SetFormatter((lrs.Formatter)(formatter))
}

// SetOutput sets the logger output.
func (logger *Logger) SetOutput(output io.Writer) {
	(*lrs.Logger)(logger).SetOutput(output)
}

func (logger *Logger) SetReportCaller(reportCaller bool) {
	(*lrs.Logger)(logger).SetReportCaller(reportCaller)
}

// ReplaceHooks replaces the logger hooks and returns the old ones
func (logger *Logger) ReplaceHooks(hooks LevelHooks) LevelHooks {
	return (LevelHooks)((*lrs.Logger)(logger).ReplaceHooks((lrs.LevelHooks)(hooks)))
}
