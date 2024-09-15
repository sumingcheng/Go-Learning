package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(
	id int,
	wg *sync.WaitGroup,
	results chan<- int,
) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作
	results <- id * 2       // 返回结果
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker2(i, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results) // 所有工作完成后关闭通道
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
