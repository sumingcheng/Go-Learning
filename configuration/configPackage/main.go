package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("./configuration/configPackage/app.yaml")
	if err != nil {
		log.Fatal("加载 .env 文件出错")
	}

	// 读取配置
	host := os.Getenv("host")
	port := os.Getenv("port")
	database := os.Getenv("database")
	username := os.Getenv("username")
	password := os.Getenv("password")

	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("database:", database)
	fmt.Println("username:", username)
	fmt.Println("password:", password)
}
