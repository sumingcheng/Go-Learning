package main

import (
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var ready bool
var wg4 sync.WaitGroup

// process 处理特定的进程，等待条件变量通知
func process(i int) {
	defer wg4.Done()
	defer cond.L.Unlock()
	cond.L.Lock()
	for !ready {
		cond.Wait()
	}
	println("进程", i, "已准备就绪")
}

func main() {
	for i := 0; i < 5; i++ {
		wg4.Add(1)
		go process(i)
	}
	println("所有协程已创建")

	time.Sleep(2 * time.Second) // 确保所有子协程都进入等待状态

	cond.L.Lock()
	ready = true
	cond.Broadcast()
	cond.L.Unlock()

	wg4.Wait() // 等待所有子协程完成
}
