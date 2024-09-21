package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type FileLogger struct {
	fileName string
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile(f.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		panic(err)
	}
}

func logMessage(logger Logger, message string) {
	logger.Log(message)
}

func main() {
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{fileName: "log.txt"}

	// 使用控制台日志记录器
	logMessage(consoleLogger, "This is a log to console")

	// 使用文件日志记录器
	logMessage(fileLogger, "This is a log to file")
}
