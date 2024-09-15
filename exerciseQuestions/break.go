package main

import "fmt"

func main() {
OuterLoop: // 定义一个标签
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j > 6 {
				fmt.Println("跳出循环时 i:", i, " j:", j)
				break OuterLoop // 跳出带有 OuterLoop 标签的循环
			}
			fmt.Println(i, "*", j, "=", i*j)
		}
	}
	fmt.Println("循环结束")
}
