package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var loginReq UserLoginRequest
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := loginReq.ValidateAll(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	})

	r.POST("/register", func(c *gin.Context) {
		var registerReq UserRegisterRequest
		if err := c.ShouldBindJSON(&registerReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := registerReq.ValidateAll(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	})

	// 运行 gin 服务器
	r.Run(":8888")
}
