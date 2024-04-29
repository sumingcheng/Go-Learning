package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.GOMAXPROCS(1)设置Go运行时最多使用1个操作系统线程来执行Go代码。这确保了程序是单核执行。
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)

	intChan <- 1
	stringChan <- "hello"

	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		fmt.Println(value)
	}
}
