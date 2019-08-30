package log

import (
	"encoding"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// NewJSONLogger returns a Logger that encodes keyvals to the Writer as a
// single JSON object. Each log event produces no more than one call to
// w.Write. The passed Writer must be safe for concurrent use by multiple
// goroutines if the returned Logger will be used concurrently.
func NewJSONLogger(w io.Writer) Logger {
	return &jsonLogger{w}
}
