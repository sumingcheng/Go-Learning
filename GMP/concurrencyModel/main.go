package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	// 循环，从 channel 接收数据并处理
	for {
		data := <-ch
		fmt.Printf("Worker %d received data: %d\n", id, data)
		time.Sleep(time.Millisecond * 100) // 模拟工作负载
	}
}

func main() {
	ch := make(chan int) // 创建一个 channel

	// 创建 5 个 worker goroutines
	for i := 0; i < 5; i++ {
		go worker(i, ch) // 每个 goroutine 代表 GMP 模型中的一个 G
	}

	// 发送数据到 channel
	for data := 0; data < 10; data++ {
		ch <- data
	}

	// 等待一段时间以便所有 goroutines 完成处理
	time.Sleep(time.Second)
}
