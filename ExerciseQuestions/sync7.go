package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	// 创建一个 MyMutex 实例
	var mu MyMutex

	// 锁定 mu
	mu.Lock()

	// 将 mu 赋值给 mu1
	var mu1 = mu

	// 对 mu 的 count 进行加一操作
	mu.count++

	// 解锁 mu
	mu.Unlock()

	// 锁定 mu1
	mu1.Lock()

	// 对 mu1 的 count 进行加一操作
	mu1.count++

	// 解锁 mu1
	mu1.Unlock()
	mu1.Unlock()

	// 打印 mu 和 mu1 的 count 值
	fmt.Println(mu.count, mu1.count)
}

// 当复制包含互斥锁的结构体时，应该谨慎处理以避免复制锁本身，因为这会导致两个独立的锁实例，可能引发并发控制上的问题
// 争取做法见 sync8.go
