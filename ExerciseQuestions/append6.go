package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var mu sync.Mutex
	var ints = make([]int, 0, 1000)

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			ints = append(ints, i)
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			ints = append(ints, i)
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(len(ints))
}

//  如果两个协程几乎同时调用 append()，可能一个协程的更改还没来得及完全写入
//  另一个协程就开始了写操作。这可能导致某些元素被覆盖
//  或者更糟糕的是，内部的数据结构（如切片的长度和容量）变得不一致。
