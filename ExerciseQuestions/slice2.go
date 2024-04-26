package main

import "fmt"

func main() {
	var a []int // 不分配内存

	// 在 Go 语言中，使用 := 语法进行变量声明和初始化时，会先计算右侧的表达式。
	// 对于 b := []int{} 这种写法，右侧的 []int{} 创建了一个空的非 nil 切片，
	// 这个操作包括了初始化一个底层数组（虽然此时数组长度为0）并将切片 b 指向这个数组。
	b := []int{} // 分配内存

	fmt.Println(a, b)
}
