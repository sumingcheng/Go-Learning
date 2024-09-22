package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/ginTest/service"
)

func main() {
	r := gin.Default()
	userService := &service.UserServiceImpl{} // 实际服务的实例化

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := userService.GetUser(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
