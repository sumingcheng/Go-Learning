package main

import (
	"fmt"
)

func main() {
	deferCall()
}

// defer2.go 的执⾏顺序是后进先出
func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
