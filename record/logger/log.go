package logger

import (
	"io"
	"log"
	"os"
)

const flag = log.Ldate | log.Ltime | log.Lshortfile

var (
	logFile     io.Writer
	logFilePath = "./logs/global.log"
)

const (
	infoPrefix  = "[INFO] "
	warnPrefix  = "[WARN] "
	errPrefix   = "[ERROR] "
	debugPrefix = "[DEBUG] "
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
	debugLogger *log.Logger
)

func New() error {
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	infoLogger = log.New(logFile, infoPrefix, flag)
	warnLogger = log.New(logFile, warnPrefix, flag)
	errLogger = log.New(logFile, errPrefix, flag)
	debugLogger = log.New(logFile, debugPrefix, flag)

	return nil
}

func Info(
	format string,
	vrs ...any,
) {
	infoLogger.Printf(format, vrs...)
}

func Warn(
	format string,
	vrs ...any,
) {
	warnLogger.Printf(format, vrs...)
}

func Error(
	format string,
	vrs ...any,
) {
	errLogger.Printf(format, vrs...)
}

func Debug(
	format string,
	vrs ...any,
) {
	debugLogger.Printf(format, vrs...)
}
