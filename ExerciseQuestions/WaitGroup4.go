package main

import (
	"fmt"
	"sync"
)

func main() {
	// 初始化一个 WaitGroup 对象，用于等待一组 goroutine 完成。
	wg := &sync.WaitGroup{}

	println("wg", wg)
	// 使用 for 循环启动 5 个 goroutine。
	for i := 0; i < 5; i++ {
		// 在 goroutine 中执行匿名函数。传递 wg 和当前的循环变量 i。
		go func(
			wg *sync.WaitGroup,
			i int,
		) {
			// 在 goroutine 开始时，增加 WaitGroup 的计数。
			wg.Add(1)

			// 打印当前的循环变量值。
			fmt.Printf("i: %d\n", i)

			// 当前 goroutine 完成后，调用 Done 来减少 WaitGroup 的计数。
			wg.Done()
		}(wg, i) // 传递 wg 和 i 的副本给 goroutine。
	}

	// 主 goroutine 等待所有通过 wg.Add 注册的 goroutine 完成。
	wg.Wait()

	// 当所有 goroutine 完成后，打印 "exit"。
	fmt.Println("exit")
}
