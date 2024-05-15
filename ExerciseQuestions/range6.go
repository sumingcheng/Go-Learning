package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 启动一个生产者 goroutine
	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second) // 每秒发送一次数据
			ch <- i
			i++
		}
	}()

	// 消费者逻辑
	for {
		select {
		case num := <-ch:
			fmt.Println("Received:", num)
		case <-time.After(5 * time.Second): // 5秒无数据接收则超时
			fmt.Println("No data within 5 seconds, loop continues...")
		}
	}
}
