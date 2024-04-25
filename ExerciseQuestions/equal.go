package main

import "fmt"

func main() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}

	// 数组的比较规则
	/*
		数组的类型相同：这意味着两个数组必须具有相同的数据类型和长度。
		数组的每个元素相等：数组中相应位置的元素必须一一对应相等。
	*/
	// 重点：不同长度的数组是不可以比较的！
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
