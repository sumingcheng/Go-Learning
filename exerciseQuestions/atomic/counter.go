package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	var wg sync.WaitGroup

	// 启动 10 个 goroutine，每个都增加 1000 次计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Count:", count)
}
