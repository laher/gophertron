package gophers

import (
	"runtime/debug"

	"github.com/Sirupsen/logrus"
)

func Recover(g func(), message string) {
	defer func() {
		logrus.Debugf("[%s] DONE Recoverable processing", message)
		if x := recover(); x != nil {
			logrus.Errorf("[%s] PANIC - recovered run time panic: %v. Stacktrace: %s", message, x, string(debug.Stack()))
		}
	}()
	logrus.Debugf("[%s] START Recoverable processing", message)
	g()
}
