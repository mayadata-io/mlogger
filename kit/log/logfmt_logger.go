package log

import (
	"bytes"
	"io"
	"sync"

	"github.com/go-logfmt/logfmt"
)

// NewLogfmtLogger returns a logger that encodes keyvals to the Writer in
// logfmt format. Each log event produces no more than one call to w.Write.
// The passed Writer must be safe for concurrent use by multiple goroutines if
// the returned Logger will be used concurrently.
func NewLogfmtLogger(w io.Writer) Logger {
	return &logfmtLogger{w}
}
