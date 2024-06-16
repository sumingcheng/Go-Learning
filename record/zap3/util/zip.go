package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

// LoggerConfig 定义了日志配置的结构体
type LoggerConfig struct {
	Mode       string               // 运行模式：dev 或 prod
	FilePath   string               // 日志文件路径
	MaxSize    int                  // 日志文件最大大小（MB）
	MaxBackups int                  // 日志文件最多备份数量
	MaxAge     int                  // 日志文件最多保存天数
	Compress   bool                 // 是否压缩旧日志文件
	Level      zapcore.LevelEnabler // 日志级别
}

// NewLogger 根据 LoggerConfig 创建一个新的 zap.Logger 实例
func NewLogger(config LoggerConfig) *zap.Logger {
	// 配置 lumberjack 作为日志文件的输出
	lumberLogger := &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	// 在函数结束时关闭 lumberjack 日志器
	defer func(lumberLogger *lumberjack.Logger) {
		err := lumberLogger.Close()
		if err != nil {
			log.Fatal("Failed to close lumberjack logger: ", err)
			return
		}
	}(lumberLogger)

	var cfg zapcore.EncoderConfig

	// 根据运行模式选择日志编码配置
	switch config.Mode {
	case "dev":
		cfg = zap.NewDevelopmentEncoderConfig()
	case "prod":
		cfg = zap.NewProductionEncoderConfig()
	default:
		cfg = zap.NewProductionEncoderConfig()
	}

	// 设置时间编码器为 ISO8601
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(cfg)

	// 创建 zapcore.Core
	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberLogger), config.Level)

	// 创建 logger 实例
	logger := zap.New(core)

	// 在函数结束时同步 logger
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("Failed to sync logger: ", err)
		}
	}(logger)

	// 替换全局 logger
	zap.ReplaceGlobals(logger)

	return logger
}
