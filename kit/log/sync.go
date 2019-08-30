package log

import (
	"io"
	"sync"
	"sync/atomic"
)

// SwapLogger wraps another logger that may be safely replaced while other
// goroutines use the SwapLogger concurrently. The zero value for a SwapLogger
// will discard all log events without error.
//
// SwapLogger serves well as a package global logger that can be changed by
// importers.
type SwapLogger struct {
	logger atomic.Value
}

type loggerStruct struct {
	Logger
}

// Log implements the Logger interface by forwarding keyvals to the currently
// wrapped logger. It does not log anything if the wrapped logger is nil.
func (l *SwapLogger) Log(keyvals ...interface{}) error {
	s, ok := l.logger.Load().(loggerStruct)
	if !ok || s.Logger == nil {
		return nil
	}
	return s.Log(keyvals...)
}

// Swap replaces the currently wrapped logger with logger. Swap may be called
// concurrently with calls to Log from other goroutines.
func (l *SwapLogger) Swap(logger Logger) {
	l.logger.Store(loggerStruct{logger})
}

// NewSyncWriter returns a new writer that is safe for concurrent use by
// multiple goroutines. Writes to the returned writer are passed on to w. If
// another write is already in progress, the calling goroutine blocks until
// the writer is available.
//
// If w implements the following interface, so does the returned writer.
//
//    interface {
//        Fd() uintptr
//    }
func NewSyncWriter(w io.Writer) io.Writer {
	switch w := w.(type) {
	case fdWriter:
		return &fdSyncWriter{fdWriter: w}
	default:
		return &syncWriter{Writer: w}
	}
}

// NewSyncLogger returns a logger that synchronizes concurrent use of the
// wrapped logger. When multiple goroutines use the SyncLogger concurrently
// only one goroutine will be allowed to log to the wrapped logger at a time.
// The other goroutines will block until the logger is available.
func NewSyncLogger(logger Logger) Logger {
	return &syncLogger{logger: logger}
}
