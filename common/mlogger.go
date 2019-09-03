package common

import (
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
)

// A Mlogger wraps the base Logger functionality in a slower, but less
// verbose, API. Any Logger can be converted to a SugaredLogger with its Sugar
// method.
//
// Unlike the Logger, the SugaredLogger doesn't insist on structured logging.
// For each log level, it exposes three methods: one for loosely-typed
// structured logging, one for println-style formatting, and one for
// printf-style formatting. For example, SugaredLoggers can produce InfoLevel
// output with Infow ("info with" structured context), Info, or Infof.
type Mlogger struct {
	base *zap.Logger
}

// Unwrap unwraps a Mlogger, exposing the original Logger. Desugaring
// is quite inexpensive, so it's reasonable for a single application to use
// both Loggers and SugaredLoggers, converting between them on the boundaries
// of performance-sensitive code.
func (s *Mlogger) Unwrap() *zap.Logger {
	base := s.base
	//base.callerSkip -= 2
	return base
}

// Named adds a sub-scope to the logger's name. See Logger.Named for details.
func (s *Mlogger) Named(name string) *Mlogger {
	return &Mlogger{base: s.base.Named(name)}
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//   sugaredLogger.With(
//     "hello", "world",
//     "failure", errors.New("oh no"),
//     Stack(),
//     "count", 42,
//     "user", User{Name: "alice"},
//  )
// is the equivalent of
//   unsugared.With(
//     String("hello", "world"),
//     String("failure", "oh no"),
//     Stack(),
//     Int("count", 42),
//     Object("user", User{Name: "alice"}),
//   )
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func (s *Mlogger) With(args ...interface{}) *Mlogger {
	return &Mlogger{base: s.base.With(s.sweetenFields(args)...)}
}

// Debug uses fmt.Sprint to construct and log a message.
func (s *Mlogger) Debug(args ...interface{}) {
	s.log(zap.DebugLevel, "", args, nil)
}

// Info uses fmt.Sprint to construct and log a message.
func (s *Mlogger) Info(args ...interface{}) {
	s.log(zap.InfoLevel, "", args, nil)
}

// Warn uses fmt.Sprint to construct and log a message.
func (s *Mlogger) Warn(args ...interface{}) {
	s.log(zap.WarnLevel, "", args, nil)
}

// Error uses fmt.Sprint to construct and log a message.
func (s *Mlogger) Error(args ...interface{}) {
	s.log(zap.ErrorLevel, "", args, nil)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *Mlogger) DPanic(args ...interface{}) {
	s.log(zap.DPanicLevel, "", args, nil)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (s *Mlogger) Panic(args ...interface{}) {
	s.log(zap.PanicLevel, "", args, nil)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (s *Mlogger) Fatal(args ...interface{}) {
	s.log(zap.FatalLevel, "", args, nil)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *Mlogger) Debugf(template string, args ...interface{}) {
	s.log(zap.DebugLevel, template, args, nil)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *Mlogger) Infof(template string, args ...interface{}) {
	s.log(zap.InfoLevel, template, args, nil)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *Mlogger) Warnf(template string, args ...interface{}) {
	s.log(zap.WarnLevel, template, args, nil)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *Mlogger) Errorf(template string, args ...interface{}) {
	s.log(zap.ErrorLevel, template, args, nil)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *Mlogger) DPanicf(template string, args ...interface{}) {
	s.log(zap.DPanicLevel, template, args, nil)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *Mlogger) Panicf(template string, args ...interface{}) {
	s.log(zap.PanicLevel, template, args, nil)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *Mlogger) Fatalf(template string, args ...interface{}) {
	s.log(zap.FatalLevel, template, args, nil)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (s *Mlogger) Debugw(msg string, keysAndValues ...interface{}) {
	s.log(zap.DebugLevel, msg, nil, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *Mlogger) Infow(msg string, keysAndValues ...interface{}) {
	s.log(zap.InfoLevel, msg, nil, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *Mlogger) Warnw(msg string, keysAndValues ...interface{}) {
	s.log(zap.WarnLevel, msg, nil, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *Mlogger) Errorw(msg string, keysAndValues ...interface{}) {
	s.log(zap.ErrorLevel, msg, nil, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s *Mlogger) DPanicw(msg string, keysAndValues ...interface{}) {
	s.log(zap.DPanicLevel, msg, nil, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s *Mlogger) Panicw(msg string, keysAndValues ...interface{}) {
	s.log(zap.PanicLevel, msg, nil, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s *Mlogger) Fatalw(msg string, keysAndValues ...interface{}) {
	s.log(zap.FatalLevel, msg, nil, keysAndValues)
}

// Sync flushes any buffered log entries.
func (s *Mlogger) Sync() error {
	return s.base.Sync()
}

func (s *Mlogger) log(lvl zapcore.Level, template string, fmtArgs []interface{}, context []interface{}) {
	// If logging at this level is completely disabled, skip the overhead of
	// string formatting.
	if lvl < zap.DPanicLevel && !s.base.Core().Enabled(lvl) {
		return
	}

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	if ce := s.base.Check(lvl, msg); ce != nil {
		ce.Write(s.sweetenFields(context)...)
	}
}

func (s *Mlogger) sweetenFields(args []interface{}) []zap.Field {
	if len(args) == 0 {
		return nil
	}

	// Allocate enough space for the worst case; if users pass only structured
	// fields, we shouldn't penalize them with extra allocations.
	fields := make([]zap.Field, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			s.base.DPanic(_oddNumberErrMsg, zap.Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			print(keyStr)
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		s.base.DPanic(_nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)
	zap.Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}


