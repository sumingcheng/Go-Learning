package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	// 快速检查切片的地址或确认切片非空。
	println("Numbers:", numbers) // Numbers: [3/3]0xc000010108
	// 能够直接输出切片内容及其结构，适合在处理数据和调试复杂结构时使用。
	fmt.Println("Numbers:", numbers) // Numbers: [1 2 3]
}
