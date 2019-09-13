package logrus

import (
	"context"
	lrs "github.com/Sirupsen/logrus"
	"time"
)

// Defines the key when adding errors using WithError.
var ErrorKey = lrs.ErrorKey

// An entry is the final or intermediate Logrus logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry lrs.Entry

func NewEntry(logger *Logger) *Entry {
	return (*Entry)(&lrs.Entry{
		Logger: (*lrs.Logger)(logger),
		// Default is three fields, plus one optional.  Give a little extra room.
		Data: make(Fields, 6),
	})
}

// Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	return (*lrs.Entry)(entry).String()
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	return (*Entry)((*lrs.Entry)(entry).WithError(err))
}

// Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	return (*Entry)((*lrs.Entry)(entry).WithContext(ctx))
}

// Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return (*Entry)((*lrs.Entry)(entry).WithField(key, value))
}

// Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	return (*Entry)((*lrs.Entry)(entry).WithFields((lrs.Fields)(fields)))
}

// Overrides the time of the Entry.
func (entry *Entry) WithTime(t time.Time) *Entry {
	return (*Entry)((*lrs.Entry)(entry).WithTime(t))
}

func (entry Entry) HasCaller() (has bool) {
	return (lrs.Entry)(entry).HasCaller()
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	(*lrs.Entry)(entry).Log((lrs.Level)(level), args)
}

func (entry *Entry) Trace(args ...interface{}) {
	(*lrs.Entry)(entry).Trace(args)
}

func (entry *Entry) Debug(args ...interface{}) {
	(*lrs.Entry)(entry).Debug(args)
}

func (entry *Entry) Print(args ...interface{}) {
	(*lrs.Entry)(entry).Print(args)
}

func (entry *Entry) Info(args ...interface{}) {
	(*lrs.Entry)(entry).Info(args)
}

func (entry *Entry) Warn(args ...interface{}) {
	(*lrs.Entry)(entry).Warn(args)
}

func (entry *Entry) Warning(args ...interface{}) {
	(*lrs.Entry)(entry).Warning(args)
}

func (entry *Entry) Error(args ...interface{}) {
	(*lrs.Entry)(entry).Error(args)
}

func (entry *Entry) Fatal(args ...interface{}) {
	(*lrs.Entry)(entry).Fatal(args)
}

func (entry *Entry) Panic(args ...interface{}) {
	(*lrs.Entry)(entry).Panic(args)
}

// Entry Printf family functions

func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	(*lrs.Entry)(entry).Logf((lrs.Level)(level), format, args)
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Tracef(format, args)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Debugf(format, args)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Infof(format, args)
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Printf(format, args)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Warnf(format, args)
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Warningf(format, args)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Errorf(format, args)
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Fatalf(format, args)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	(*lrs.Entry)(entry).Panicf(format, args)
}

// Entry Println family functions

func (entry *Entry) Logln(level Level, args ...interface{}) {
	(*lrs.Entry)(entry).Logln((lrs.Level)(level), args)
}

func (entry *Entry) Traceln(args ...interface{}) {
	(*lrs.Entry)(entry).Traceln(args)
}

func (entry *Entry) Debugln(args ...interface{}) {
	(*lrs.Entry)(entry).Debugln(args)
}

func (entry *Entry) Infoln(args ...interface{}) {
	(*lrs.Entry)(entry).Infoln(args)
}

func (entry *Entry) Println(args ...interface{}) {
	(*lrs.Entry)(entry).Println(args)
}

func (entry *Entry) Warnln(args ...interface{}) {
	(*lrs.Entry)(entry).Warnln(args)
}

func (entry *Entry) Warningln(args ...interface{}) {
	(*lrs.Entry)(entry).Warningln(args)
}

func (entry *Entry) Errorln(args ...interface{}) {
	(*lrs.Entry)(entry).Errorln(args)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	(*lrs.Entry)(entry).Fatalln(args)
}

func (entry *Entry) Panicln(args ...interface{}) {
	(*lrs.Entry)(entry).Panicln(args)
}
