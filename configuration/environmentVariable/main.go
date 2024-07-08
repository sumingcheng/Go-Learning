package main

import (
	"fmt"
	"os"
)

func main() {
	// 设置环境变量
	os.Setenv("MY_ENV", "Gopher")
	// 获取并打印环境变量
	value := os.Getenv("MY_ENV")
	fmt.Println("MY_ENV:", value)

	// 查找环境变量
	value, exists := os.LookupEnv("MY_ENV")
	if exists {
		fmt.Println("MY_ENV exists with value:", value)
	} else {
		fmt.Println("MY_ENV does not exist.")
	}

	// 先设置环境变量
	os.Setenv("APP_ENV", "Gopher")
	// 然后删除环境变量
	os.Unsetenv("APP_ENV")

	// 验证环境变量是否还存在
	value2, exists := os.LookupEnv("APP_ENV")
	if exists {
		fmt.Println("APP_ENV still exists with value2:", value2)
	} else {
		fmt.Println("APP_ENV has been removed.")
	}
}
