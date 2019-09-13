package logrus

import (
	lrs "github.com/Sirupsen/logrus"
)

type WrapFormatter struct {
	internal Formatter
}

// Format renders a single log entry
func (f *WrapFormatter) Format(entry *lrs.Entry) ([]byte, error) {
	return f.internal.Format((*Entry)(entry))
}

// TextFormatter formats logs into text
type TextFormatter lrs.TextFormatter

// Format renders a single log entry
func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	return (*lrs.TextFormatter)(f).Format((*lrs.Entry)(entry))
}
