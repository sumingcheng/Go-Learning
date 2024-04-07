package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GlobalMiddleware 全局中间件
func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是全局中间件，我在处理每一个请求")
		c.Next()
	}
}

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// LocalMiddleware 局部中间件
func LocalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是局部中间件，只有特定路由会触发我")
		c.Next()
	}
}

func main() {
	router := gin.Default()

	// 应用全局中间件
	router.Use(GlobalMiddleware())
	// 应用跨域中间件
	router.Use(CORSMiddleware())

	// 一个没有中间件的路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "没有中间件"})
	})

	// 应用局部中间件的路由
	router.GET("/local", LocalMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "有局部中间件"})
	})

	// 启动服务
	router.Run(":8080")
}
