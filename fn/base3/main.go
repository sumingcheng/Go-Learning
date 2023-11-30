package main

import "fmt"

func main() {
	test(1, 2)
}

// 相同的类型可以只写一个
func test(a, b int) {
	fmt.Println(a + b)
}
