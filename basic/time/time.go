package main

import (
	"fmt"
	"time"
)

func main() {
	// 加载时区
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// 使用特定时区构造时间
	shanghaiTime := time.Date(2024, 3, 18, 12, 0, 0, 0, loc)
	fmt.Println("上海时间：", shanghaiTime)

	// 当前时间转换为上海时间
	now := time.Now()
	nowInShanghai := now.In(loc)
	fmt.Println("现在的上海时间：", nowInShanghai)

	// 定义时间格式和时间字符串
	layout := "2006-01-02 15:04:05"
	timeStr := "2024-03-18 08:00:00"

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("加载时区失败:", err)
		return
	}

	// 使用 ParseInLocation 解析时间字符串
	t, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		fmt.Println("解析时间错误:", err)
		return
	}

	fmt.Println("解析得到的时间:", t)
	fmt.Println("解析得到的时间（UTC）:", t.UTC())
}
