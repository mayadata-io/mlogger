package logrus

import (
	lrs "github.com/Sirupsen/logrus"
)

// InternalFormatter formats logs into text
type InternalFormatter lrs.TextFormatter

// Format renders a single log entry
func (f *InternalFormatter) Format(entry *lrs.Entry) ([]byte, error) {
	return (*lrs.TextFormatter)(f).Format(entry)
}

type TextFormatter struct {
	tformatter 	lrs.TextFormatter
}

// Format renders a single log entry
func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	return f.tformatter.Format((*lrs.Entry)(entry))
}

// GetTFormatter return TextFormatter
func (f *TextFormatter) GetTFormatter() lrs.TextFormatter {
	return f.tformatter
}