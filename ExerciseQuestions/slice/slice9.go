package main

import "fmt"

func main() {
	var src, dst []int   // 声明源切片和目标切片
	src = []int{1, 2, 3} // 初始化源切片

	// copy(dst, src) 在这里不会成功，因为dst没有初始化，其长度为0
	dst = make([]int, len(src)) // 正确初始化dst为与src相同长度
	copy(dst, src)              // 将src中的内容复制到dst

	fmt.Println(dst) // 打印复制后的dst
}
