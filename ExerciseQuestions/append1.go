package main

import "fmt"

func main() {
	s := make([]int, 5)
	// s 初始化后长度为5，容量为 5 再添加元素的时候会追加到后面
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
