package test

import "go.uber.org/zap"

func Test() {
	zap.L().Info("这是一个【Test】普通日志")
	zap.S().Warn("这是一个【Test】警告日志")
}
