package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())       // 查看 CPU 核心数
	fmt.Println(runtime.NumGoroutine()) // 查看当前 goroutine 数量
	fmt.Println(runtime.NumCgoCall())   // 查看当前 cgo 调用次数
}
