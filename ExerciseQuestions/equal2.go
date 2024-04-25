package main

import "fmt"

func main() {
	// 在 Go 语言中，interface{} 是一种特殊的类型，可以存储任意类型的值。
	// 然而，当你声明一个 interface{} 类型的变量而没有给它赋值时
	// 该变量的初始状态具有两个部分的零值：类型部分为 nil，值部分也为 nil。这就是所谓的 "空接口"。
	var i interface{}

	if i == nil {
		fmt.Println("nil")
		return
	} else {
		fmt.Println("not nil")
	}
}
