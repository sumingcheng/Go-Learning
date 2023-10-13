package main

import "fmt"

func main() {
	const PI float32 = 3.14
	fmt.Println(PI)

	// 单个声明
	const name = "admin"

	// 多个声明
	const (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)

	// iota 批量声明常量，第一个必须要初始化，后续的可以不初始化
	// 常量赋值，如果不写，会跟从之前的一个常量
	// iota 初始值为 0
	const (
		num1 = iota
		num2
		num3
		num4
	)

	fmt.Println(num1, num2, num3, num4)

	// 匿名常量
	const (
		a1 = iota
		_
		b2
		_
		c3
	)

	fmt.Println(a1, b2, c3)
}
