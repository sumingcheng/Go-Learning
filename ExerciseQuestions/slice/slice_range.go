package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5} // 创建一个整型切片

	m := make(map[int]int) // 创建一个整型键和值的映射

	// 遍历切片，将其元素添加到映射中
	// key为切片的索引，val为对应的值
	for key, val := range slice {
		m[key] = val
	}

	// 遍历映射并打印每个键值对
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}
