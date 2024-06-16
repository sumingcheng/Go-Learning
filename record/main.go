package main

import (
	"log"
	"logger/logger"
)

func main() {
	err := logger.New() // 使用 logger 包
	if err != nil {
		log.Fatalf("Failed to open logger file: %s", err.Error())
		return
	}

	logger.Info("This is an info message: %s", "[info]")
	logger.Warn("This is a warning message: %s", "[warn]")
	logger.Error("This is an error message: %s", "[error]")
	logger.Debug("This is a debug message: %s", "[debug]")
}
