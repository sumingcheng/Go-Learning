package main

import "fmt"

func main() {
	originalSlice := []int{1, 2, 3}                                                     // 初始切片
	fmt.Println("originalSlice", originalSlice, len(originalSlice), cap(originalSlice)) // 打印初始切片

	newSlice := append(originalSlice, 9)                            // 扩容并添加新元素
	fmt.Println("newSlice", newSlice, len(newSlice), cap(newSlice)) // 打印新切片

	fmt.Println("originalSlice", originalSlice, len(originalSlice), cap(originalSlice)) // 再次打印原切片
}
