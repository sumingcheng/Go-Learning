package main

import "fmt"

func main() {
	const a = 10
	const b = 20
	// b 是一个常量，即使定义后没有在代码中使用，也不会导致编译错误。
	// 这与未使用的变量不同，未使用的变量会导致编译错误。
	fmt.Println(a)
}
