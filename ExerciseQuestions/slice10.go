package main

import "fmt"

func main() {
	var src = [3]int{1, 2, 3}
	var dst [3]int

	// 将数组转换为切片，并复制内容
	copy(dst[:], src[:])

	fmt.Println(dst) // 输出: [1 2 3]
}
