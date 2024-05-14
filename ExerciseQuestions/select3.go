package main

import (
	"fmt"
	"time"
)

func A3() int {
	time.Sleep(100 * time.Millisecond) // A3函数等待100毫秒
	return 1                           // 返回1
}

func B3() int {
	time.Sleep(1000 * time.Millisecond) // B3函数等待1000毫秒
	return 2                            // 返回2
}

func main() {
	ch := make(chan int, 1) // 创建一个容量为1的通道

	go func() { // 启动一个goroutine
		select {
		case ch <- A3(): // 尝试发送A3()的结果到通道
		case ch <- B3(): // 尝试发送B3()的结果到通道
		default: // 如果以上都不可执行，则执行默认操作
			ch <- 3 // 发送3到通道
		}
	}()

	fmt.Println(<-ch) // 从通道接收值并打印
}
