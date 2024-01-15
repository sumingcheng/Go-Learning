package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个会自动取消的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 启动一个goroutine，传递上下文
	go doSomething(ctx, "level 1")

	// 等待足够的时间让 `doSomething` 执行
	time.Sleep(3 * time.Second)
}

func doSomething(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		// 如果收到Done信号，说明父上下文已被取消
		fmt.Println(name, "got the done signal and is canceling")
	case <-time.After(1 * time.Second):
		// 否则，模拟工作，并创建一个新的子goroutine
		fmt.Println(name, "is doing some work")
		// 假设我们要创建另一个级别的goroutine
		go doSomething(ctx, "level 2")
	}
}
