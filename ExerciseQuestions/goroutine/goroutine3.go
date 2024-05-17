package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 在 main 函数中初始化通道，确保通道在使用前已初始化
	ch := make(chan int, 1)

	// 启动第一个 goroutine 来发送数据
	go func() {
		// 向通道发送数据
		ch <- 1
	}()

	// 启动第二个 goroutine 来接收数据
	go func() {
		// 等待 1 秒确保第一个 goroutine 已发送数据
		time.Sleep(time.Second)
		// 接收并打印从通道接收到的数据
		fmt.Println("Received:", <-ch)
	}()

	// 创建一个定时器，每秒触发一次
	ticker := time.Tick(1 * time.Second)
	// 无限循环，监控活跃的 goroutine 数量
	for range ticker {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
