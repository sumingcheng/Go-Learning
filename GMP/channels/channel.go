package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 100)

	// 2个生产者，往channel里写入元素
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			ch <- i
		}
	}()

	// 1个消费者
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		sum := 0
		for {
			a, ok := <-ch
			if !ok { // channel被关闭，且channel为空，跳出循环
				break
			} else {
				sum += a
			}
		}
		fmt.Printf("sum=%d\n", sum)
	}()

	wg.Wait()
	close(ch)

	wg2.Wait()
}
