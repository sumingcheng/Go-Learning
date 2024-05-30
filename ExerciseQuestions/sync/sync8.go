package main

import (
	"fmt"
	"sync"
)

// 定义一个结构体 MyMutex1，包含一个整型变量 count 和一个 sync.Mutex 互斥锁
type MyMutex1 struct {
	count int
	sync.Mutex
}

func main() {
	// 创建一个 MyMutex1 的指针实例
	mu := &MyMutex1{}

	// 锁定 mu
	mu.Lock()

	// 由于 mu 已经是一个指针，我们直接将它的引用赋值给 mu1
	mu1 := mu

	// 对 mu 的 count 进行加一操作
	mu.count++

	// 解锁 mu
	mu.Unlock()

	// 使用 mu1，这里不需要重新锁定，因为 mu1 和 mu 指向同一个对象
	mu1.Lock()

	// 对 mu1 的 count 进行加一操作
	mu1.count++

	// 解锁 mu1
	mu1.Unlock()

	// 打印 mu 和 mu1 的 count 值，由于它们指向同一个对象，结果应该是相同的
	fmt.Println(mu.count, mu1.count)
}
