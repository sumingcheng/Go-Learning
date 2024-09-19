package main

import "fmt"

func main() {
	// 初始化一个整数切片v，包含三个元素1, 2, 3
	v := []int{1, 2, 3}

	// 初始化i为0，n为切片v的长度，即3
	// 注意：这个循环条件计算的是初始的切片长度
	println(len(v))
	for i, n := 0, len(v); i < n; i++ {
		// 向切片v追加元素i，初始时i的值从0开始递增
		v = append(v, i)
	}

	// 打印最终的切片v
	fmt.Println(v) // 输出结果是：[1 2 3 0 1 2]
}
