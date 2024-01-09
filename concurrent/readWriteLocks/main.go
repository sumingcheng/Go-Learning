package main

import (
	"fmt"
	"sync"
)

var (
	rwMutex sync.RWMutex
	count   int
	wg      sync.WaitGroup // 定义 WaitGroup
)

func main() {
	wg.Add(2) // 增加两个计数，对应两个 goroutine

	// 启动一个用于写操作的 goroutine
	go func() {
		defer wg.Done() // 完成时调用 Done
		Write(10)
	}()

	// 启动一个用于读操作的 goroutine
	go func() {
		defer wg.Done() // 完成时调用 Done
		Read()
	}()

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("所有goroutine执行完毕")
}

func Read() {
	rwMutex.RLock() // 申请读锁
	fmt.Println("读取的数值为：", count)
	rwMutex.RUnlock() // 释放读锁
}

func Write(x int) {
	rwMutex.Lock() // 申请写锁
	count = x
	fmt.Println("写入的数值为：", x)
	rwMutex.Unlock() // 释放写锁
}
