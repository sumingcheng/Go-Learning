package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 安全地关闭通道
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("a:", a)
		}
	}()

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("ok")
}
