package logrus

import (
	lrs "github.com/Sirupsen/logrus"
	"github.com/mayadata-io/mlogger/common"
)

var (
	logger = common.Logger
)

// Fields type, used to pass to `WithFields`.
type Fields lrs.Fields

// Level type
type Level lrs.Level

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	return (lrs.Level)(level).String()
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
	lrslevel, err := lrs.ParseLevel(lvl)
	return (Level)(lrslevel), err
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (level *Level) UnmarshalText(text []byte) error {
	return (*lrs.Level)(level).UnmarshalText(text)
}

func (level Level) MarshalText() ([]byte, error) {
	return (lrs.Level)(level).MarshalText()
}

// A constant exposing all logging levels
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = Level(lrs.PanicLevel)
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = Level(lrs.FatalLevel)
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = Level(lrs.ErrorLevel)
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = Level(lrs.WarnLevel)
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = Level(lrs.InfoLevel)
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = Level(lrs.DebugLevel)
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel = Level(lrs.TraceLevel)
)

// StdLogger is what your logrus-enabled library should take, that way
// it'll accept a stdlib logger and a logrus logger. There's no standard
// interface, this is the closest we get, unfortunately.
type StdLogger lrs.StdLogger

// The FieldLogger interface generalizes the Entry and Logger types
type FieldLogger lrs.FieldLogger

// Ext1FieldLogger (the first extension to FieldLogger) is superfluous, it is
// here for consistancy. Do not use. Use Logger or Entry instead.
type Ext1FieldLogger lrs.Ext1FieldLogger
