package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

const (
	FATAL Level = iota
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

var levels = [...]string{
	"fatal",
	"error",
	"warn",
	"info",
	"debug",
	"trace"}

var OutLevel Level = ERROR

func (lvl Level) String() string { return levels[lvl] }

func setOutput(w io.Writer) {
	log.SetOutput(w)
}

func logMsg(level Level, v ...interface{}) string {
	return fmt.Sprintf("%s: %v", level, fmt.Sprint(v...))
}

func logMsgf(level Level, format string, v ...interface{}) string {
	return fmt.Sprintf("%s: %v", level, fmt.Sprintf(format, v...))
}

func output(level Level, v ...interface{}) {
	if level <= OutLevel {
		log.Print(logMsg(level, v...))
	}
}

func outputf(level Level, format string, v ...interface{}) {
	if level <= OutLevel {
		log.Print(logMsgf(level, format, v...))
	}
}

func Trace(v ...interface{})                 { output(TRACE, v...) }
func Tracef(format string, v ...interface{}) { outputf(TRACE, format, v...) }
func Debug(v ...interface{})                 { output(DEBUG, v...) }
func Debugf(format string, v ...interface{}) { outputf(DEBUG, format, v...) }
func Info(v ...interface{})                  { output(INFO, v...) }
func Infof(format string, v ...interface{})  { outputf(INFO, format, v...) }
func Warn(v ...interface{})                  { output(WARN, v...) }
func Warnf(format string, v ...interface{})  { outputf(WARN, format, v...) }
func Error(v ...interface{})                 { output(ERROR, v...) }
func Errorf(format string, v ...interface{}) { outputf(ERROR, format, v...) }
func Fatal(v ...interface{}) {
	output(FATAL, v...)
	os.Exit(1)
}
func Fatalf(format string, v ...interface{}) {
	outputf(FATAL, format, v...)
	os.Exit(1)
}
func Panic(v ...interface{}) {
	s := logMsg(FATAL, v...)
	output(FATAL, s)
	panic(s)
}
func Panicf(format string, v ...interface{}) {
	s := logMsgf(FATAL, format, v...)
	output(FATAL, s)
	panic(s)
}
