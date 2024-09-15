package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}

	// 核心：切片作为函数参数时，传递的是切片的引用，而不会创建一个新的切片。
	b := variadic(a[:]...) // 注意这里从数组转为切片，并传递所有元素

	b[0], b[1] = b[1], b[0]

	fmt.Println(a)
}

func variadic(ints ...int) []int {
	return ints // 直接返回接收到的整型切片
}
