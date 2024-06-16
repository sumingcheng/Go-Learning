package main

import (
	"go.uber.org/zap"
	"logger/zap3/test"
	"logger/zap3/util"
)

func main() {
	prodLogger := util.NewLogger(util.LoggerConfig{
		Mode:       "prod",
		FilePath:   "./zap3/logs/log.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     1,
		Level:      zap.InfoLevel,
	})

	prodLogger.Info("这是一个普通日志", zap.String("module", "main"))
	prodLogger.Warn("这是一个错误日志", zap.String("module", "main"))

	// zap.L() 和 zap.S() 是 zap 日志库提供的两个全局函数
	// 用于获取全局的 zap.Logger 和 zap.SugaredLogger 实例。
	zap.L().Info("这是一个【全局】普通日志")
	zap.S().Info("这是一个【全局】普通日志")

	test.Test()

	for i := 0; i < 10000; i++ {
		prodLogger.Info("这是一个普通日志", zap.String("module", "main"))
		prodLogger.Warn("这是一个错误日志", zap.String("module", "main"))
		zap.L().Info("这是一个【全局】普通日志")
		zap.S().Info("这是一个【全局】普通日志")
	}
}
