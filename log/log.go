package log

import (
	"log"
	"os"
	"runtime/debug"
)

var Verbose bool

var (
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.LUTC)
	infoLogger  = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.LUTC)
	errorLogger = log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.LUTC)
	fatalLogger = log.New(os.Stdout, "[FATAL] ", log.LstdFlags|log.LUTC)
)

var (
	Debugf = func(format string, v ...interface{}) {
		if Verbose {
			debugLogger.Printf(format, v...)
		}
	}
	Debug = func(v ...interface{}) {
		if Verbose {
			debugLogger.Println(v...)
		}
	}
	Infof = func(format string, v ...interface{}) {
		infoLogger.Printf(format, v...)
	}
	Info = func(v ...interface{}) {
		infoLogger.Println(v...)
	}
	Errorf = func(format string, v ...interface{}) {
		errorLogger.Printf(format, v...)
	}
	Error = func(v ...interface{}) {
		errorLogger.Println(v...)
	}
	Fatalf = func(format string, v ...interface{}) {
		fatalLogger.Printf(format, v...)
	}
	Fatal = func(v ...interface{}) {
		stack := debug.Stack()
		fatalLogger.Println(v...)
		fatalLogger.Println(stack)
		os.Exit(1)
	}
)
