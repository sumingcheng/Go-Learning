package main

import (
	"fmt"
)

var a int

func init() {
	a = 10
	fmt.Println("执行的初始化函数")
}

// init 函数会在 main 函数之前执行
func main() {
	fmt.Println("执行的 main 函数")
	fmt.Println("a的值：", a)
}
