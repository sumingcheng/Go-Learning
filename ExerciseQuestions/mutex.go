package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用sync.RWMutex来支持更高效的读操作并发
type data struct {
	sync.RWMutex
}

// 改进的test方法，增加一个读写标记isWrite，判断是否为写操作
func (d *data) test(
	s string,
	isWrite bool,
) {
	if isWrite {
		d.Lock() // 写操作，使用互斥锁
	} else {
		d.RLock() // 读操作，使用读锁
	}

	// 确保相应的锁被释放
	defer func() {
		if isWrite {
			d.Unlock()
		} else {
			d.RUnlock()
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		// 为了示例保留睡眠调用，但在实际应用中应考虑移除或替换
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	var d data

	wg.Add(2)

	// 第一个goroutine执行读操作
	go func() {
		defer wg.Done()
		d.test("read", false) // isWrite设置为false，代表读操作
	}()

	// 第二个goroutine执行写操作
	go func() {
		defer wg.Done()
		d.test("write", true) // isWrite设置为true，代表写操作
	}()

	wg.Wait()
}
