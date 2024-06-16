package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
)

func main() {
	// 配置 lumberjack 作为 zap 的日志输出文件
	logWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./zap2/logs/log-info.log", // 日志文件的位置
		MaxSize:    1,                          // 文件大小限制，单位：MB
		MaxBackups: 3,                          // 最多保留3个备份
		MaxAge:     28,                         // 文件最多保存28天
		Compress:   true,                       // 是否压缩/归档旧文件
	})

	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(config)

	const (
		fileFlag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		prem     = 0666
	)

	fileMap := map[string]string{
		"DEBUG":  "./zap2/logs/log-debug.log",
		"INFO":   "./zap2/logs/log-info.log",
		"WARN":   "./zap2/logs/log-warn.log",
		"ERROR":  "./zap2/logs/log-error.log",
		"PANIC":  "./zap2/logs/log-panic.log",
		"DPANIC": "./zap2/logs/log-dpanic.log",
		"FATAL":  "./zap2/logs/log-fatal.log",
	}

	debugFile, _ := os.OpenFile(fileMap["DEBUG"], fileFlag, prem)
	//infoFile, _ := os.OpenFile(fileMap["INFO"], fileFlag, prem)
	warnFile, _ := os.OpenFile(fileMap["WARN"], fileFlag, prem)
	errorFile, _ := os.OpenFile(fileMap["ERROR"], fileFlag, prem)
	panicFile, _ := os.OpenFile(fileMap["PANIC"], fileFlag, prem)
	dpanicFile, _ := os.OpenFile(fileMap["DPANIC"], fileFlag, prem)
	fatalFile, _ := os.OpenFile(fileMap["FATAL"], fileFlag, prem)

	teeCore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(debugFile), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logWriter), zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(warnFile), zapcore.WarnLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(errorFile), zapcore.ErrorLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(panicFile), zapcore.PanicLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(dpanicFile), zapcore.DPanicLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fatalFile), zapcore.FatalLevel),
	)

	logger := zap.New(teeCore, zap.AddCaller())

	defer logger.Sync()

	//logger.Error("这是一个错误日志")
	//logger.Debug("这是一个调试日志")
	//logger.Warn("这是一个警告日志")
	//logger.Panic("这是一个panic日志")
	//logger.DPanic("这是一个dpanic日志")
	//logger.Fatal("这是一个fatal日志")

	for i := 0; i < 100000; i++ {
		logger.Info("这是一个普通日志 - "+strconv.Itoa(i),
			zap.Namespace("name"),
			zap.String("name", "zhangsan"),
			zap.Int("age", 18),
		)
	}
}
