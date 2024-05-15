package main

import (
	"fmt"
	"sync"
)

var (
	mm     = make(map[int]int)
	mutex7 sync.Mutex
	wg7    sync.WaitGroup // 添加 WaitGroup 用于同步 goroutine
)

/*
因为在某些情况下 read 函数执行的时间早于对应 key 的 write 函数的执行。
由于 map 在 Go 中如果找不到键的值，默认返回该类型的零值
对于整数类型 int，零值就是 0
因此，当 read 函数先于 write 函数执行时，您会看到值为 0 的输出。
*/
func main() {
	// 为每个启动的 goroutine 计数
	wg7.Add(20) // 因为有10个写操作和10个读操作

	// 即使使用了互斥锁来控制访问同一个 map，也无法保证 write 操作和 read 操作之间的先后顺序
	for i := 0; i < 10; i++ {
		go write(i, i*2)
		go read(i)
	}

	wg7.Wait() // 等待所有 goroutine 完成
}

func write(key, value int) {
	defer wg7.Done()      // 完成后递减计数
	defer mutex7.Unlock() // 解锁
	mutex7.Lock()         // 锁定
	mm[key] = value
}

func read(key int) {
	defer wg7.Done()      // 完成后递减计数
	defer mutex7.Unlock() // 解锁
	mutex7.Lock()         // 锁定
	fmt.Printf("Key: %d, Value: %d\n", key, mm[key])
}
