package logrus

import (
	lrs "github.com/Sirupsen/logrus"
)

// FieldMap allows customization of the key names for default fields.
type FieldMap lrs.FieldMap

// JSONFormatter formats logs into parsable json
type JSONFormatter lrs.JSONFormatter

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	return (*lrs.JSONFormatter)(f).Format((*lrs.Entry)(entry))
}
