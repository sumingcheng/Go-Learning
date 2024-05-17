package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 设置最大并发系统线程数为1
	runtime.GOMAXPROCS(1)

	// 启动一个匿名goroutine
	go func() {
		// 从0到9循环打印数字
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	// 主协程进入死循环，阻止程序退出
	// for{} 独占资源，导致其他goroutine无法执行
	for {
	}
}
