package main

import (
	"fmt"
	"runtime"
	"sync"
)

const N = 26 // 定义常量 N，代表字母表中的字母数量，即 26 个英文字母。

func main() {
	const GOMAXPROCS = 1           // 设置使用的最大 CPU 核数为 1
	runtime.GOMAXPROCS(GOMAXPROCS) // 应用 GOMAXPROCS 设置，让程序运行在单核 CPU 上

	var wg sync.WaitGroup // 声明一个 WaitGroup，用于处理并发任务的同步
	wg.Add(2 * N)         // 添加等待处理的任务总数，每个字母对应两个任务（小写和大写），共 2*N 个

	// 循环 N 次，每次循环对应字母表中的一个字母
	for i := 0; i < N; i++ {
		// 启动第一个 goroutine，负责输出小写字母
		go func(i int) {
			defer wg.Done()         // 在 goroutine 完成时，调用 Done() 减少 WaitGroup 的计数
			runtime.Gosched()       // 让出 CPU 时间片，允许其他 goroutine 运行
			fmt.Printf("%c", 'a'+i) // 输出当前索引对应的小写字母
		}(i)

		// 启动第二个 goroutine，负责输出大写字母
		go func(i int) {
			defer wg.Done()         // 在 goroutine 完成时，调用 Done() 减少 WaitGroup 的计数
			fmt.Printf("%c", 'A'+i) // 输出当前索引对应的大写字母
		}(i)
	}

	wg.Wait() // 等待所有 goroutine 完成
}
