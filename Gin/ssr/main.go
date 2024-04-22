package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	r := gin.Default()

	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("当前工作目录:", cwd)

	// 根据当前工作目录构建模板路径
	templatePath := filepath.Join(cwd, "./Gin/ssr/templates/**/*")
	r.LoadHTMLGlob(templatePath)

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{"title": "Main website"})
	})

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{"title": "Main website"})
	})

	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{"title": "Main website"})
	})

	r.Run(":8888")
}
