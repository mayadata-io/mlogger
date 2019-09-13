package logrus

import (
	lrs "github.com/Sirupsen/logrus"
)

// Exit runs all the Logrus atexit handlers and then terminates the program using os.Exit(code)
func Exit(code int) {
	lrs.Exit(code)
}

// RegisterExitHandler appends a Logrus Exit handler to the list of handlers,
// call logrus.Exit to invoke all handlers. The handlers will also be invoked when
// any Fatal log entry is made.
//
// This method is useful when a caller wishes to use logrus to log a fatal
// message but also needs to gracefully shutdown. An example usecase could be
// closing database connections, or sending a alert that the application is
// closing.
func RegisterExitHandler(handler func()) {
	lrs.RegisterExitHandler(handler)
}

// DeferExitHandler prepends a Logrus Exit handler to the list of handlers,
// call logrus.Exit to invoke all handlers. The handlers will also be invoked when
// any Fatal log entry is made.
//
// This method is useful when a caller wishes to use logrus to log a fatal
// message but also needs to gracefully shutdown. An example usecase could be
// closing database connections, or sending a alert that the application is
// closing.
func DeferExitHandler(handler func()) {
	lrs.DeferExitHandler(handler)
}
