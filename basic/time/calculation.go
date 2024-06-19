package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间：", now)

	// 加上10分钟
	later := now.Add(10 * time.Minute)
	fmt.Println("十分钟后：", later)

	// 减去一小时
	before := now.Add(-1 * time.Hour)
	fmt.Println("一小时前：", before)
}
