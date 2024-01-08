package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化 Gin 路由器
	r := gin.Default()

	// 添加 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 启动服务器
	r.Run(":8080") // 默认监听 localhost:8080
}
