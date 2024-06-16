package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type LoggerConfig struct {
	Mode       string
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Level      zapcore.LevelEnabler
}

func NewLogger(config LoggerConfig) *zap.Logger {
	lumberLogger := &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	defer func(lumberLogger *lumberjack.Logger) {
		err := lumberLogger.Close()
		if err != nil {
			log.Fatal("Failed to close lumberjack logger: ", err)
			return
		}
	}(lumberLogger)

	var cfg zapcore.EncoderConfig

	switch config.Mode {
	case "dev":
		cfg = zap.NewDevelopmentEncoderConfig()
	case "prod":
		cfg = zap.NewProductionEncoderConfig()
	default:
		cfg = zap.NewProductionEncoderConfig()
	}

	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberLogger), config.Level)

	logger := zap.New(core)

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("Failed to sync logger: ", err)
		}
	}(logger)

	zap.ReplaceGlobals(logger)

	return logger
}
