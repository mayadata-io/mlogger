package logrus

import (
	lrs "github.com/Sirupsen/logrus"
	"io"
)

func (logger *Logger) Writer() *io.PipeWriter {
	return (*lrs.Logger)(logger).Writer()
}

func (logger *Logger) WriterLevel(level Level) *io.PipeWriter {
	return (*lrs.Logger)(logger).WriterLevel((lrs.Level)(level))
}

func (entry *Entry) Writer() *io.PipeWriter {
	return (*lrs.Entry)(entry).Writer()
}

func (entry *Entry) WriterLevel(level Level) *io.PipeWriter {
	return (*lrs.Entry)(entry).WriterLevel((lrs.Level)(level))
}