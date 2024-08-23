package main

import "fmt"

func main() {
	client := NewRedisClient()
	// 示例操作
	err := client.SetString("mykey", "Hello, World!")
	if err != nil {
		fmt.Println("Error setting string:", err)
	}
	// 添加其他类型的操作测试
}
