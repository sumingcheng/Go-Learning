package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前工作目录:", dir)

	router := gin.Default()

	router.LoadHTMLGlob("./Gin/basic1/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"Title":   "主页",
			"Header":  "欢迎来到我的网站",
			"Content": "这是使用Gin进行服务端渲染的示例。",
		})
	})

	router.Run(":8080")
}
