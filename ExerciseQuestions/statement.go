package main

import "fmt"

func main() {
	// 短变量声明只能在函数内部使用 并且同一个变量只能声明一次
	one := 0
	one = 1
	fmt.Println(one)
}
