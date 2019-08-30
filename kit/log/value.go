package log

import (
	"runtime"
	"strconv"
	"strings"
	"time"
)

// A Valuer generates a log value. When passed to With or WithPrefix in a
// value element (odd indexes), it represents a dynamic value which is re-
// evaluated with each log event.
type Valuer func() interface{}

// Timestamp returns a timestamp Valuer. It invokes the t function to get the
// time; unless you are doing something tricky, pass time.Now.
//
// Most users will want to use DefaultTimestamp or DefaultTimestampUTC, which
// are TimestampFormats that use the RFC3339Nano format.
func Timestamp(t func() time.Time) Valuer {
	return func() interface{} { return t() }
}

// TimestampFormat returns a timestamp Valuer with a custom time format. It
// invokes the t function to get the time to format; unless you are doing
// something tricky, pass time.Now. The layout string is passed to
// Time.Format.
//
// Most users will want to use DefaultTimestamp or DefaultTimestampUTC, which
// are TimestampFormats that use the RFC3339Nano format.
func TimestampFormat(t func() time.Time, layout string) Valuer {
	return func() interface{} {
		return timeFormat{
			time:   t(),
			layout: layout,
		}
	}
}

// Caller returns a Valuer that returns a file and line from a specified depth
// in the callstack. Users will probably want to use DefaultCaller.
func Caller(depth int) Valuer {
	return func() interface{} {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		// using idx+1 below handles both of following cases:
		// idx == -1 because no "/" was found, or
		// idx >= 0 and we want to start at the character after the found "/".
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

var (
	// DefaultTimestamp is a Valuer that returns the current wallclock time,
	// respecting time zones, when bound.
	DefaultTimestamp = TimestampFormat(time.Now, time.RFC3339Nano)

	// DefaultTimestampUTC is a Valuer that returns the current time in UTC
	// when bound.
	DefaultTimestampUTC = TimestampFormat(
		func() time.Time { return time.Now().UTC() },
		time.RFC3339Nano,
	)

	// DefaultCaller is a Valuer that returns the file and line where the Log
	// method was invoked. It can only be used with log.With.
	DefaultCaller = Caller(3)
)
