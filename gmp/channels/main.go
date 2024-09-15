package main

import (
	"fmt"
)

func main() {
	done := make(chan bool, 10) // 创建一个通道
	for i := 0; i < 10; i++ {
		go printNumber(i, done)
	}
	for i := 0; i < 10; i++ {
		<-done // 等待 goroutine 完成
	}
}

func printNumber(i int, done chan bool) {
	fmt.Println(i)
	done <- true // 发送完成信号
}
