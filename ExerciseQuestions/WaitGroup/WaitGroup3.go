package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // 增加 WaitGroup 计数，表示需要等待一个 goroutine 完成

	go func() {
		fmt.Println("1") // 打印 "1"

		wg.Done() // 完成一个任务，减少 WaitGroup 计数

		wg.Add(1) // 增加 WaitGroup 计数，再增加一个任务
	}()

	wg.Wait() // 等待所有 WaitGroup 计数为 0，即所有任务完成。这里只有一个任务，所以等待这个任务完成
	// 打印 "1"
	// panic
}
