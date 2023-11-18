package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func newLogger() *Logger {
	return &Logger{
		debugLogger:   log.New(os.Stdout, "\033[34m[DEBUG]\033[0m ", log.Ldate|log.Ltime),
		infoLogger:    log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stdout, "\033[33m[WARNING]\033[0m ", log.Ldate|log.Ltime),
		errorLogger:   log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ldate|log.Ltime),
	}
}

var l = newLogger()

func Debug(v ...interface{}) {
	l.debugLogger.Output(2, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprintln(v...))
}

func Warning(v ...interface{}) {
	l.warningLogger.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintln(v...))
}

func logf(level LogLevel, format string, v ...interface{}) {
	switch level {
	case DebugLevel:
		l.debugLogger.Output(2, fmt.Sprintf(format, v...))
	case InfoLevel:
		l.infoLogger.Output(2, fmt.Sprintf(format, v...))
	case WarningLevel:
		l.warningLogger.Output(2, fmt.Sprintf(format, v...))
	case ErrorLevel:
		l.errorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Debugf(format string, v ...interface{}) {
	logf(DebugLevel, format, v...)
}

func Infof(format string, v ...interface{}) {
	logf(InfoLevel, format, v...)
}

func Warningf(format string, v ...interface{}) {
	logf(WarningLevel, format, v...)
}

func Errorf(format string, v ...interface{}) {
	logf(ErrorLevel, format, v...)
}

func removeCwdPath(path string) string {
	cwd, err := os.Getwd()
	if err != nil {
		return path
	}

	return path[len(cwd)+1:]
}

func runLog(level LogLevel, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	logf(level, "[%s:%d] %s", removeCwdPath(file), line, fmt.Sprintln(v...))
}

func Debugln(v ...interface{}) {
	runLog(DebugLevel, v...)
}

func Infoln(v ...interface{}) {
	runLog(InfoLevel, v...)
}

func Warningln(v ...interface{}) {
	runLog(WarningLevel, v...)
}

func Errorln(v ...interface{}) {
	runLog(ErrorLevel, v...)
}
