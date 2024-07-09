package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type GinConfig struct {
	Ip   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}

type Config struct {
	*GinConfig `mapstructure:"gin"`
	*DBConfig  `mapstructure:"db"`
}

var cfg = new(Config)

// Initialize 函数用于初始化配置，返回配置结构体指针和错误。
func Initialize() (*Config, error) {
	// 设置 Viper 来读取指定路径下的 YAML 配置文件。
	viper.SetConfigFile("configuration/viper-gin/config/config.yaml")

	// 设置 Viper 监听配置文件的更改。
	viper.WatchConfig()

	// 定义配置文件发生变化时的回调函数。
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 当配置文件发生变化时，重新从配置文件中解析数据到 cfg 结构体。
		if err := viper.Unmarshal(&cfg); err != nil {
			// 如果解析过程出错，记录致命错误并退出程序。
			log.Fatalf("解析配置文件失败：%s", err)
		}
	})

	// 尝试读取配置文件。
	if err := viper.ReadInConfig(); err != nil {
		// 如果读取配置文件失败，返回错误。
		return nil, err
	}

	// 解析配置文件内容到 cfg 结构体。
	if err := viper.Unmarshal(&cfg); err != nil {
		// 如果解析配置失败，返回错误。
		return nil, err
	}

	// 如果配置文件成功被读取和解析，返回 cfg 结构体指针和 nil 错误。
	return cfg, nil
}

func main() {
	cfg, err := Initialize()
	if err != nil {
		log.Fatalf("初始化配置失败：%s", err)
	}

	fmt.Printf("%#v", cfg.DBConfig)
}
