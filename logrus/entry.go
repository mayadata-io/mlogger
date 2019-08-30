package logrus

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Defines the key when adding errors using WithError.
var ErrorKey = "error"

// An entry is the final or intermediate Logrus logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry struct {
	Logger *Logger

	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	Level Level

	// Calling method, with package name
	Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer

	// Contains the context set by the user. Useful for hook processing etc.
	Context context.Context

	// err may contain a field formatting error
	err string
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		// Default is three fields, plus one optional.  Give a little extra room.
		Data: make(Fields, 6),
	}
}

// Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	serialized, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return "", err
	}
	str := string(serialized)
	return str, nil
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	return entry.WithField(ErrorKey, err)
}

// Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	return &Entry{Logger: entry.Logger, Data: entry.Data, Time: entry.Time, err: entry.err, Context: ctx}
}

// Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return entry.WithFields(Fields{key: value})
}

// Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	data := make(Fields, len(entry.Data)+len(fields))
	for k, v := range entry.Data {
		data[k] = v
	}
	fieldErr := entry.err
	for k, v := range fields {
		isErrField := false
		if t := reflect.TypeOf(v); t != nil {
			switch t.Kind() {
			case reflect.Func:
				isErrField = true
			case reflect.Ptr:
				isErrField = t.Elem().Kind() == reflect.Func
			}
		}
		if isErrField {
			tmp := fmt.Sprintf("can not add field %q", k)
			if fieldErr != "" {
				fieldErr = entry.err + ", " + tmp
			} else {
				fieldErr = tmp
			}
		} else {
			data[k] = v
		}
	}
	return &Entry{Logger: entry.Logger, Data: data, Time: entry.Time, err: fieldErr, Context: entry.Context}
}

// Overrides the time of the Entry.
func (entry *Entry) WithTime(t time.Time) *Entry {
	return &Entry{Logger: entry.Logger, Data: entry.Data, Time: t, err: entry.err, Context: entry.Context}
}

// getCaller retrieves the name of the first non-logrus calling function
func getCaller() *runtime.Frame {

	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		logrusPackage = getPackageName(runtime.FuncForPC(pcs[1]).Name())

		// now that we have the cache, we can skip a minimum count of known-logrus functions
		// XXX this is dubious, the number of frames may vary
		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != logrusPackage {
			return &f
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

func (entry Entry) HasCaller() (has bool) {
	return entry.Logger != nil &&
		entry.Logger.ReportCaller &&
		entry.Caller != nil
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.log(level, fmt.Sprint(args...))
	}
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.Log(TraceLevel, args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.Log(DebugLevel, args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.Log(InfoLevel, args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.Log(WarnLevel, args...)
}

func (entry *Entry) Warning(args ...interface{}) {
	entry.Warn(args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.Log(ErrorLevel, args...)
}

func (entry *Entry) Fatal(args ...interface{}) {
	entry.Log(FatalLevel, args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panic(args ...interface{}) {
	entry.Log(PanicLevel, args...)
	panic(fmt.Sprint(args...))
}

// Entry Printf family functions

func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.Log(level, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.Logf(TraceLevel, format, args...)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.Logf(DebugLevel, format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.Logf(InfoLevel, format, args...)
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.Infof(format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.Logf(WarnLevel, format, args...)
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.Warnf(format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.Logf(ErrorLevel, format, args...)
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.Logf(FatalLevel, format, args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.Logf(PanicLevel, format, args...)
}

// Entry Println family functions

func (entry *Entry) Logln(level Level, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.Log(level, entry.sprintlnn(args...))
	}
}

func (entry *Entry) Traceln(args ...interface{}) {
	entry.Logln(TraceLevel, args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.Logln(DebugLevel, args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.Logln(InfoLevel, args...)
}

func (entry *Entry) Println(args ...interface{}) {
	entry.Infoln(args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.Logln(WarnLevel, args...)
}

func (entry *Entry) Warningln(args ...interface{}) {
	entry.Warnln(args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.Logln(ErrorLevel, args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.Logln(FatalLevel, args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.Logln(PanicLevel, args...)
}

