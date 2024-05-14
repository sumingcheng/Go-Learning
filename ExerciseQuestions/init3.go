package main

import (
	"log"
	"os"
)

var Config map[string]string

func init() {
	Config = make(map[string]string)

	// 检查环境变量是否已设定
	appEnv, exists := os.LookupEnv("APP_ENV")
	if !exists {
		log.Fatal("必须设置 APP_ENV 环境变量")
	}

	// 根据环境变量加载配置
	if appEnv == "production" {
		Config["apiEndpoint"] = "https://api.example.com"
	} else {
		Config["apiEndpoint"] = "https://api.dev.example.com"
	}

	// 可以添加更多配置加载逻辑
}
