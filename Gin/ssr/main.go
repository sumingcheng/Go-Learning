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
	r.LoadHTMLGlob("templates/*")

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Print(dir)

	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{})
	})

	r.Run(":8888")
}
