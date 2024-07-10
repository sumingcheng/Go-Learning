package Consul

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func init() {
	viper.SetConfigFile("./Consul/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
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
		ID:      "serviceB",
		Name:    "service-b",
		Address: "172.16.50.251",
		Port:    8081,
		Check: &api.AgentServiceCheck{
			HTTP:                           "http://172.16.50.251:8081/health",
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("向 Consul 注册服务 B 时发生错误: %s", err)
	}

	// 设置 Gin 路由
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/call-a", func(c *gin.Context) {
		// 发现服务 A 并调用其 API
		services, _, err := consulClient.Health().Service("service-a", "", true, nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "从 Consul 获取服务时发生错误")
			return
		}
		if len(services) > 0 {
			serviceA := services[0].Service
			url := "http://" + serviceA.Address + ":" + strconv.Itoa(serviceA.Port) + "/data"
			resp, err := http.Get(url)
			if err != nil {
				c.String(http.StatusInternalServerError, "调用服务 A 时发生错误")
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.String(http.StatusInternalServerError, "读取服务 A 响应时发生错误")
				return
			}
			c.String(http.StatusOK, string(body))
		} else {
			c.String(http.StatusNotFound, "未找到服务 A")
		}
	})

	// 启动 HTTP 服务器
	log.Println("服务 B 正在端口 8081 上运行")
	r.Run(":8081")
}
