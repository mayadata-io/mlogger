package logrus

import (
	lrs "github.com/Sirupsen/logrus"
)

// TextFormatter formats logs into text
type TextFormatter lrs.TextFormatter

// Format renders a single log entry
func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	return (*lrs.TextFormatter)(f).Format((*lrs.Entry)(entry))
}

