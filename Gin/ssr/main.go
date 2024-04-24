package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

var students = []Student{
	{1, "张三", 18, 90},
	{2, "李四", 19, 80},
	{3, "王五", 20, 70},
	{4, "赵六", 21, 60},
}

func main() {
	r := gin.Default()

	// 处理跨域
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})

	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("当前工作目录:", cwd)

	// 根据当前工作目录构建模板路径
	templatePath := filepath.Join(cwd, "./Gin/ssr/templates/**/*")
	staticPath := filepath.Join(cwd, "./Gin/ssr/client")
	r.LoadHTMLGlob(templatePath)
	fmt.Println("模板路径:", staticPath)
	r.Static("/client", staticPath)

	r.GET("/", func(c *gin.Context) {
		//c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Redirect(http.StatusFound, "/list")
		//c.Redirect(http.StatusMovedPermanently, "/list")
		c.HTML(http.StatusOK, "home.html", gin.H{"title": "Main website"})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{"title": "Main website"})
	})

	r.GET("/react", func(c *gin.Context) {
		c.HTML(http.StatusOK, "react.html", gin.H{"title": "Main website"})
	})

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"title":    "Main website",
			"students": students,
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{"title": "Main website"})
	})

	r.GET("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": students,
		})
	})

	r.Run(":9999")
}
