package logrus

import (
	"time"
	lrs "github.com/Sirupsen/logrus"
)


// Default key names for the default fields
const (
	defaultTimestampFormat = time.RFC3339
	FieldKeyMsg            = lrs.FieldKeyMsg
	FieldKeyLevel          = lrs.FieldKeyLevel
	FieldKeyTime           = lrs.FieldKeyTime
	FieldKeyLogrusError    = lrs.FieldKeyLogrusError
	FieldKeyFunc           = lrs.FieldKeyFunc
	FieldKeyFile           = lrs.FieldKeyFile
)

// The Formatter interface is used to implement a custom Formatter. It takes an
// `Entry`. It exposes all the fields, including the default ones:
//
// * `entry.Data["msg"]`. The message passed from Info, Warn, Error ..
// * `entry.Data["time"]`. The timestamp.
// * `entry.Data["level"]. The level the entry was logged at.
//
// Any additional fields added with `WithField` or `WithFields` are also in
// `entry.Data`. Format is expected to return an array of bytes which are then
// logged to `logger.Out`.
//type Formatter lrs.Formatter
type Formatter interface {
	Format(*Entry) ([]byte, error)
}