package main

import (
	"fmt"
	"sync"
)

var rwMutex sync.RWMutex
var data2 int
var wg sync.WaitGroup

func ReadData(id int) {
	defer rwMutex.RUnlock()
	rwMutex.RLock()
	fmt.Printf("Goroutine %d read data: %d\n", id, data2)

	wg.Done()
}

func WriteData(id, d int) {
	defer rwMutex.Unlock()
	rwMutex.Lock()
	data2 = d
	fmt.Printf("Goroutine %d write data: %d\n", id, d)
	wg.Done()
}

func main() {
	// 启动多个读写goroutines
	numGoroutines := 10
	wg.Add(numGoroutines * 2) // 因为有numGoroutines个读goroutine和numGoroutines个写goroutine

	for i := 0; i < numGoroutines; i++ {
		go WriteData(i, i*10) // 每个goroutine写入不同的数据
		go ReadData(i)        // 同时进行读操作
	}

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("Final Data:", data2)
}
