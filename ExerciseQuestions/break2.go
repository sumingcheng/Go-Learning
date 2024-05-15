package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break // 只会跳出内层的 j 循环
			}
			fmt.Println("内层循环：", i, j)
		}
		fmt.Println("外层循环：", i)
	}
}
