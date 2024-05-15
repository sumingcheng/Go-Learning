package main

import "fmt"

func main() {
	a := 1

	for i := 0; i < 5; i++ {
		// 声明了一个新的局部变量 a，而不是修改外部的 a
		a := a + 1
		a = a * 2
	}

	fmt.Println(a)
}
