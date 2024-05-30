package main

import (
	"fmt"
	"sync"
)

var (
	sMap sync.Map
	wg8  sync.WaitGroup
)

// 使用 sync.Map 的 Store 方法来安全地存储键值对，使用 Load 方法来安全地检索值。
func main() {
	// 设置启动的 goroutine 数量
	wg8.Add(20) // 10个写操作，10个读操作

	for i := 0; i < 10; i++ {
		go writeMap(i, i*2)
		go readMap(i)
	}

	// 等待所有 goroutine 完成
	wg8.Wait()
}

func writeMap(key, value int) {
	sMap.Store(key, value) // 使用 Store 方法写入数据
	fmt.Printf("已写入: Key: %d, Value: %d\n", key, value)
	wg8.Done() // 通知 WaitGroup 任务完成
}

func readMap(key int) {
	value, ok := sMap.Load(key) // 使用 Load 方法读取数据
	if ok {
		fmt.Printf("已读取: Key: %d, Value: %d\n", key, value)
	} else {
		fmt.Printf("键 %d 不存在\n", key)
	}
	wg8.Done() // 通知 WaitGroup 任务完成
}
