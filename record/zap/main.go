package main

import (
	"go.uber.org/zap"
)

func main() {

	// zap 的基本使用
	//logger, err := zap.NewDevelopment()
	//
	//if err != nil {
	//	log.Fatal("logger creation failed")
	//}
	//
	//defer logger.Sync() // 刷新任何缓冲的日志条目
	//logger.Debug("这是一个调试日志")
	//logger.Info("这是一个信息日志")
	//logger.Warn("这是一个警告日志") // 会输出到标准错误
	//logger.Error("这是一个错误日志")
	//logger.DPanic("这是一个DPanic日志")
	//logger.Panic("这是一个Panic日志")

	// zap 的 Hooks
	//logger, err := zap.NewProduction(zap.Hooks(func(entry zapcore.Entry) error {
	//	log.Printf("zap Hooks log: %v", entry)
	//	return nil
	//}))
	//
	//if err != nil {
	//	log.Fatal("logger creation failed")
	//}
	//
	//logger.Sync()
	//
	//su := logger.Sugar()
	//su.Infof("这是一个调试日志 @ %s", "debug")
	//su.Infow("这是一个信息日志", "IP", "0.0.0.0", "port", 8080)

	// 自定义 Logger 的配置
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),         // 日志级别设置为信息级别
		Development:       false,                                       // 是否开启开发模式，开发模式下，日志将包括调用栈信息
		Encoding:          "json",                                      // 设置日志格式为json
		EncoderConfig:     zap.NewProductionEncoderConfig(),            // 生产环境下的日志编码配置
		OutputPaths:       []string{"stdout", "./zap/logs/global.log"}, // 日志输出路径，可以是标准输出或文件
		ErrorOutputPaths:  []string{"stderr"},                          // 错误输出路径，错误日志将写入标准错误
		InitialFields:     map[string]interface{}{"appName": "myApp"},  // 初始字段会添加到所有日志条目
		DisableCaller:     false,                                       // 是否禁止记录日志调用者信息
		DisableStacktrace: false,                                       // 是否禁止自动堆栈跟踪记录
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	//logger.Info("这是一个信息日志")
	//logger.Warn("这是一个警告日志")
	//logger.Error("这是一个错误日志")

	su := logger.Sugar()
	su.Warnf("这是一个警告日志 @ %s", "warn")
}
