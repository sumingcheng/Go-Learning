package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// 首先，append 函数的第一个参数是目标切片，第二个参数是要追加的切片的元素。
	// 比如 append(s1, 4, 5, 6)。
	// 使用 s2... 语法。这种语法告诉 Go 将 s2 切片"解压"，将其元素展开为 append 函数的多个参数。
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
