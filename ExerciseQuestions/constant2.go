package main

import "fmt"

func main() {
	const (
		a = 10
		// 不指定初始化值，表示使用上一个常量的值
		a1
	)

	fmt.Println(a)
	fmt.Println(a1)
}
