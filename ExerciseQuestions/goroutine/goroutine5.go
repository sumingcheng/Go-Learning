package main

import (
	"fmt"
	"sync"
)

const M = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{} // 初始化WaitGroup，用于等待所有goroutine完成
	mu := &sync.Mutex{}     // 初始化互斥锁，用于保护map写入操作

	wg.Add(M) // 告诉WaitGroup需要等待M个goroutine

	for i := 0; i < M; i++ { // 循环启动M个goroutine
		go func(i int) { // 使用匿名函数并发执行
			defer wg.Done() // 在函数退出时调用Done，通知WaitGroup一个goroutine已完成
			mu.Lock()       // 在修改map前获取锁，防止并发写入冲突
			m[i] = i        // 将当前的索引i存入map
			mu.Unlock()     // 修改完成后释放锁
		}(i)
	}

	wg.Wait() // 阻塞，直到所有goroutine调用wg.Done

	println(len(m)) // 输出map的长度，应该为M，因为每个goroutine都添加了一个唯一的键
	fmt.Println(m)
}
