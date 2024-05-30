package main

import "fmt"

type Point struct{ x, y int }

func main() {
	s := []Point{
		{1, 2},
		{3, 4},
	}

	// 使用索引来直接访问并修改切片中的元素
	for i := range s {
		s[i].x, s[i].y = s[i].y, s[i].x
	}

	fmt.Println(s) // 输出修改后的切片
}
