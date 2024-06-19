package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 格式化时间
	fmt.Println("RFC3339格式：", now.Format(time.RFC3339))
	fmt.Println("自定义格式：", now.Format("2006-01-02 15:04:05"))

	// 解析字符串为时间
	timeStr := "2024-03-18 08:00:00"
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	fmt.Println("解析的时间：", parsedTime)
}
