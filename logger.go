package logger

import (
	"fmt"
	"time"
	"runtime"
	"path"
	"os"
)

const (
	//输出格式
	MSG_FORMAT  = "%s %s:%d ▶ %s "
	//info 前缀
	INFO_TAG    = "[INFO] "
	//error 前缀
	ERROR_TAG   = "[ERROR] "
	//warning 前缀
	WARNING_TAG = "[WARNING] "
	//warning 前缀
	FATAL_TAG   = "[FATAL] "
)

type loggerInfo struct {
	Time     string
	Line     int
	Filename string
	Message  string
	format   string
}

func Info(message string)  {
	messageFormat(INFO_TAG, message)
}

func Infof(format string, a ...interface{})  {
	messageFormat(INFO_TAG, fmt.Sprintf(format, a ...))
}

func Error(message string)  {
	messageFormat(ERROR_TAG, message)
}

func Errorf(format string, a ...interface{})  {
	messageFormat(ERROR_TAG, fmt.Sprintf(format, a ...))
}

func Warning(message string)  {
	messageFormat(WARNING_TAG, message)
}

func Warningf(format string, a ...interface{})  {
	messageFormat(WARNING_TAG, fmt.Sprintf(format, a ...))
}

func Fatal(message string)  {
	messageFormat(FATAL_TAG, message)
	os.Exit(0)
}

func Fatalf(format string, a ...interface{})  {
	messageFormat(FATAL_TAG, fmt.Sprintf(format, a ...))
	os.Exit(0)
}

func output(info *loggerInfo)  {
	msg := fmt.Sprintf(info.format, info.Time, info.Filename, info.Line, info.Message)
	fmt.Println(msg)
}

func messageFormat(tag, message string)  {
	_, filename, line, _ := runtime.Caller(2)
	filename = path.Base(filename)
	info := &loggerInfo{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Line: line,
		Filename: filename,
		Message: message,
		format: tag + MSG_FORMAT,
	}
	output(info)
}