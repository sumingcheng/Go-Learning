package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")                 // 配置文件名称 (无扩展名)
	viper.SetConfigType("yaml")                   // 如果配置文件没有扩展名，则需要设置类型
	viper.AddConfigPath("./configuration/viper/") // 配置文件路径
	viper.AddConfigPath("$HOME/.appname")         // 可以添加多个搜索路径
	viper.AddConfigPath(".")                      // 当前工作目录
}

func main() {
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 使用 viper.Get 获取配置信息
	fmt.Println(viper.Get("username"))
	//readConfig()
}

func readConfig() {
	viper.SetConfigFile("./configuration/viper/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(viper.Get("username"))
}
