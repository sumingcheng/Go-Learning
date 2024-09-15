package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 创建一个验证器实例
var validate = validator.New()

type Header struct {
	ContentType string `header:"Content-Type" validate:"required"` // 校验必须是 JSON 类型
	UserAgent   string `header:"User-Agent" validate:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		var h Header
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 进行 Header 验证
		if err := validate.Struct(h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Content-Type": h.ContentType, "User-Agent": h.UserAgent})
	})

	router.Run(":8888")
}
