package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigFile("./Consul/config.yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatalf("致命错误: 配置文件读取失败: %s \n", err)
	}
}

func main() {

	// 创建 Consul 客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = viper.GetString("consul.address") + ":" + viper.GetString("consul.port")
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("创建 Consul 客户端时发生错误: %s", err)
	}

	// 注册服务
	registration := &api.AgentServiceRegistration{
		ID:      "serviceA",
		Name:    "service-a",
		Address: "localhost",
		Port:    8080,
		Check: &api.AgentServiceCheck{
			HTTP:                           "http://172.16.50.251:8080/health",
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("向 Consul 注册服务 A 时发生错误: %s", err)
	}

	// 设置 Gin 路由
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/data", func(c *gin.Context) {
		c.String(http.StatusOK, "来自服务 A 的问候！")
	})

	// 启动 HTTP 服务器
	log.Println("服务 A 正在端口 8080 上运行")
	r.Run(":8080")
}
