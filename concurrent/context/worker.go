package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, level int) {
	for {
		select {
		case <-ctx.Done():
			// 当上下文被取消时，Done channel 会被关闭
			fmt.Printf("Worker at level %d was told to stop\n", level)
			return
		default:
			// 模拟一些工作
			fmt.Printf("Worker at level %d is working...\n", level)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 创建一个根context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动第一级worker
	go worker(ctx, 1)

	// 启动第二级worker
	go func() {
		// 传递同一个ctx
		worker(ctx, 2)
	}()

	// 模拟在一段时间后取消操作
	time.Sleep(5 * time.Second)
	fmt.Println("主 goroutine 发送取消信号")
	cancel() // 取消所有使用ctx的goroutine

	// 给goroutine足够的时间来清理并退出
	time.Sleep(2 * time.Second)
}
