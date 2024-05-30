package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(
	id int,
	wg *sync.WaitGroup,
) {
	defer wg.Done() // 确保 Done() 在函数退出时调用
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 增加一个 goroutine 计数
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有的 goroutine 完成
	fmt.Println("All workers done")
}
