package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	ContentType string `header:"Content-Type"`
	UserAgent   string `header:"User-Agent"`
}

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		var h Header
		// 绑定 Header 到结构体
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Content-Type": h.ContentType, "User-Agent": h.UserAgent})
	})

	router.Run(":8080")
}
